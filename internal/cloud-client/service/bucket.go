package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/byteso/Xcloud/api/cloud-client/v1/types"
	"github.com/byteso/Xcloud/internal/config"
	"github.com/byteso/Xcloud/internal/database"
	"github.com/byteso/Xcloud/internal/entity"
	"github.com/byteso/Xcloud/internal/repository"
	"github.com/h2non/filetype"
	"github.com/minio/minio-go/v7"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"go.mongodb.org/mongo-driver/bson"
)

func convert(object *minio.Object) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(object)
	b := buf.Bytes()

	return b
}

func CreateBucket(bucketName string) error {
	var client = database.ObjectClient
	fmt.Println(bucketName)
	err := client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: "us-east-1", ObjectLocking: false})
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("hello")

	return nil
}

func CreateFolder(bucketName, folderName string) error {
	err := UploadSource(bucketName, folderName, nil)
	if err != nil {
		return err
	}

	return nil
}

func GetSource(bucketName string, id string) (types.ResponseGetSource, error) {
	var (
		client = database.ObjectClient
		r      types.ResponseGetSource
	)

	object, err := client.GetObject(context.Background(), bucketName, id, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return r, err
	}

	b := convert(object)
	reader := bytes.NewBuffer(b)

	kind, _ := filetype.Match(b)
	if kind == filetype.Unknown {
		fmt.Println("error")
		return r, errors.New("unknown")
	}

	r.ContentType = kind.MIME.Value
	r.ContentLength = int64(reader.Len())
	r.ExtraHeaders = map[string]string{"Content-Disposition": `attachment; filename="` + id + `"`}
	r.Reader = reader

	return r, nil
}

func UploadSource(bucketName string, folderName string, data []*multipart.FileHeader) error {
	var (
		client = database.ObjectClient

		ff     []types.UploadSource
		filter = bson.D{{"bucketName", bucketName}}
	)

	for _, d := range data {
		var f types.UploadSource
		data, err := d.Open()
		if err != nil {
			return err
		}

		defer data.Close()

		f.Name = d.Filename
		f.Size = d.Size
		f.Data = data

		ff = append(ff, f)

	}
	if ff == nil {
		var f types.UploadSource
		empty, err := os.Open("../../empty")
		if err != nil {
			return err
		}

		defer empty.Close()

		stat, err := empty.Stat()
		if err != nil {
			return err
		}
		f.Name = stat.Name()
		f.Size = stat.Size()
		f.Data = empty

		ff = append(ff, f)

		defer empty.Close()
	}

	for _, f := range ff {
		var (
			path          string
			fileBytesCopy []byte
			dateByte      []byte
		)

		if folderName == "" {
			fileBytes, err := ioutil.ReadAll(f.Data)
			if err != nil {
				return err
			}

			fileBytesCopy = fileBytes
			f.Data = bytes.NewReader(fileBytes)

			c := config.Config.CloudClient
			if filetype.IsImage(fileBytes) || filetype.IsVideo(fileBytes) {
				path = c.DefaultPhotoFolder
			} else {
				path = c.DefaultFileFolder
			}
		} else {
			path = folderName
		}

		// get data
		exif.RegisterParsers(mknote.All...)

		x, err := exif.Decode(f.Data)
		if err != nil {
			return err
		}

		fmt.Println(x)

		dateByte = fileBytesCopy
		f.Data = bytes.NewReader(fileBytesCopy)

		photoInfo := GetPhotoInfo(x)

		date := time.Unix(int64(photoInfo.TakenDateTime), 0)

		yearStr := strconv.Itoa(date.Year())
		monthStr := strconv.Itoa(int(date.Month()))

		pathYear := yearStr + "/"
		pathMonth := monthStr + "/"

		upload, err := client.PutObject(context.Background(), bucketName, path+pathYear+pathMonth+f.Name, f.Data, f.Size, minio.PutObjectOptions{})
		fmt.Println(upload.Location)
		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println(upload)

		// create thumbnails

		// create file info
		f.Path = path + pathYear + pathMonth
		f.Data = bytes.NewReader(dateByte)
		err = CreateFileInfo(bucketName, f)
		if err != nil {
			return err
		}

	}

	usedCapacity := GetBucketUsage(bucketName)

	update := bson.D{{"$set", bson.D{{"usedCapacity", usedCapacity}}}}

	err := repository.Update(entity.DatabaseNameStorage, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func DownloadSource(bucketName string, folderName string, key string) (types.ResponseDownloadSource, error) {
	var (
		client   = database.ObjectClient
		response types.ResponseDownloadSource
	)
	object, err := client.GetObject(context.Background(), bucketName, folderName+key, minio.GetObjectOptions{})
	if err != nil {
		return response, err
	}

	b := convert(object)

	kind, _ := filetype.Match(b)
	if kind == filetype.Unknown {
		return response, errors.New("unknown")
	}

	response.ContentType = kind.MIME.Value
	response.Data = b

	return response, nil
}

func DeleteSource(bucketName string, folderName string, key string) error {
	var client = database.ObjectClient
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}

	err := client.RemoveObject(context.Background(), bucketName, key, opts)
	if err != nil {
		return err
	}

	return nil
}

func GetBucketUsage(bucketName string) float64 {
	path := "minio/" + bucketName
	multiple := 1
	cmd := exec.Command("mc", "du", path)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	arr := strings.Split(string(out), "iB")
	fmt.Println(len(arr[0]))

	if ok := strings.Contains(arr[0], "M"); ok {
		multiple = 1024
	} else if ok := strings.Contains(arr[0], "G"); ok {
		multiple = 1024 * 1024
	}

	usedCapacity, err := strconv.ParseFloat(arr[0][:len(arr[0])-1], 64)
	if err != nil {
		log.Fatal(err)
	}

	return float64(usedCapacity * float64(multiple))
}
