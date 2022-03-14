package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/byteso/Xcloud/api/cloud-client/v1/types"
	"github.com/byteso/Xcloud/internal/auth"
	"github.com/byteso/Xcloud/internal/config"
	"github.com/byteso/Xcloud/internal/entity"
	"github.com/byteso/Xcloud/internal/repository"
	"github.com/h2non/filetype"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Response struct {
	ResourceSets []ResourceSets `json:"resourceSets"`
}

type ResourceSets struct {
	Resources []Resources `json:"resources"`
}

type Resources struct {
	Address Address `json:"address"`
}

type Address struct {
	AdminDistrict  string `json:"adminDistrict"`
	AdminDistrict2 string `json:"adminDistrict2"`
	CountryRegion  string `json:"countryRegion"`
	Locality       string `json:"locality"`
}

func getAddress(longitude float64, latitude float64, lang string) (types.Address, error) {
	var (
		a types.Address
		c = config.Config.Location
	)

	params := url.Values{}
	latitudeStr := strconv.FormatFloat(latitude, 'f', 6, 64)
	longitudeStr := strconv.FormatFloat(longitude, 'f', 6, 64)
	getUrl, err := url.Parse(c.BingMapUrl + latitudeStr + "," + longitudeStr)
	if err != nil {
		return a, err
	}

	params.Set("key", c.BingMapKey)
	params.Set("c", lang)
	getUrl.RawQuery = params.Encode()

	urlPath := getUrl.String()
	resp, err := http.Get(urlPath)
	if err != nil {
		return a, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var r Response

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return a, err
		}

		err = json.Unmarshal(body, &r)
		if err != nil {
			return a, err
		}

		address := r.ResourceSets[0].Resources[0].Address

		a.City = address.Locality
		a.CountryOrRegion = address.CountryRegion
		a.Locality = address.AdminDistrict2
		a.State = address.AdminDistrict
	}

	return a, nil
}

func GetPhotoInfo(exifInfo *exif.Exif) entity.Photo {
	var p entity.Photo

	cameraMake, _ := exifInfo.Get("Make")
	cameraModel, _ := exifInfo.Get("Model")
	ExposureDenominator, _ := exifInfo.Get("ExposureTime")
	FNumber, _ := exifInfo.Get("FNumber")
	FocalLength, _ := exifInfo.Get("FocalLength")
	Iso, _ := exifInfo.Get("ISOSpeedRatings")
	Orientation, _ := exifInfo.Get("Orientation")
	TakenDateTime, _ := exifInfo.Get("DateTimeOriginal")

	cameraMakeStr := strings.Split(cameraMake.String(), "\"")
	cameraModelStr := strings.Split(cameraModel.String(), "\"")

	p.CameraMake = cameraMakeStr[1]
	p.CameraModel = cameraModelStr[1]

	num, den, _ := ExposureDenominator.Rat2(0)
	p.ExposureDenominator = float64(den)
	p.ExposureNumerator = float64(num)

	num, den, _ = FNumber.Rat2(0)
	p.FNumber = float64(num) / float64(den)

	num, den, _ = FocalLength.Rat2(0)
	p.FocalLength = float64(num) / float64(den)

	p.Iso, _ = strconv.ParseFloat(Iso.String(), 64)

	if Orientation != nil {
		fmt.Println("hello")
		fmt.Println(Orientation.String())
		p.Orientation, _ = strconv.ParseFloat(Orientation.String(), 64)
	}

	dateTimeStr := strings.FieldsFunc(TakenDateTime.String(), func(r rune) bool {
		return r == '"' || r == ' '
	})

	TakenTime, _ := time.ParseInLocation("2006:01:02 15:04:05", dateTimeStr[0]+" "+dateTimeStr[1], time.UTC)

	p.TakenDateTime = uint64(TakenTime.Unix())

	return p
}

func CreateFileInfo(bucketName string, data types.UploadSource) error {
	var f entity.FileInfo

	buf, err := ioutil.ReadAll(data.Data)
	if err != nil {
		return err
	}

	if filetype.IsImage(buf) {
		var (
			l entity.Location
			p entity.Photo
		)

		kind, _ := filetype.Match(buf)
		if kind == filetype.Unknown {
			return errors.New("UnknownFileType")
		}

		f.BucketName = bucketName
		f.Id = auth.NewSHA256(buf)
		fmt.Printf(f.Id)
		f.File.MimeType = kind.MIME.Value
		f.FileSystemInfo = entity.FileSystemInfo{
			CreatedDateTime:      uint64(time.Now().Unix()),
			LastModifiedDateTime: uint64(time.Now().Unix()),
		}
		f.Name = data.Name
		f.ParentReference.Path = data.Path
		f.Size = float64(data.Size)

		exif.RegisterParsers(mknote.All...)

		data.Data = bytes.NewReader(buf)
		x, err := exif.Decode(data.Data)
		if err != nil {
			return err
		}

		cameraMake, _ := x.Get("Make")
		cameraMakeStr := strings.Split(cameraMake.String(), "\"")
		if cameraMakeStr[1] == "Apple" {
			alt, _ := x.Get("GPSAltitude")
			num, den, _ := alt.Rat2(0)
			l.Altitude = float64(num) / float64(den)
		}

		lat, _ := x.Get("GPSLatitude")
		for i := 0; i < int(lat.Count); i++ {
			num, den, _ := lat.Rat2(i)

			l.Latitude += (float64(num) / float64(den)) / math.Pow(60, float64(i))
		}

		lot, _ := x.Get("GPSLongitude")
		for i := 0; i < int(lat.Count); i++ {
			num, den, _ := lot.Rat2(i)
			l.Longitude += (float64(num) / float64(den)) / math.Pow(60, float64(i))
		}

		/*
			f.Location.Address, err = getAddress(f.Location.Longitude, f.Location.Altitude, "")
			if err != nil {
				return err
			}
		*/

		p = GetPhotoInfo(x)

		insert, err := bson.Marshal(struct {
			entity.FileInfo
			entity.Location
			entity.Photo
		}{
			f,
			l,
			p,
		})
		if err != nil {
			return err
		}

		if err := repository.Insert(entity.DatabaseNameFileInfoName, insert); err != nil {
			return err
		}

	}

	return nil
}

func DeleteFileInfo(bucketName string, id string) error {
	return nil
}

func Items(bucketName string) ([]types.Items, error) {
	var (
		mm = make(map[string]struct {
			types.Items
			unix int64
		})
		response []types.Items
	)
	filter := bson.D{{"fileinfo.bucketName", bucketName}}
	result, err := repository.Find(entity.DatabaseNameFileInfoName, filter)
	if err != nil {
		return response, err
	}

	for _, r := range result {
		d := r[3].Value.(primitive.D)
		unix := d[8].Value.(int64)

		dateTime := time.Unix(unix, 0)
		dateTimeStr := dateTime.Format("2006-01-02 15:04:05")
		tag := strings.Split(dateTimeStr, " ")

		tagTime, _ := time.ParseInLocation("2006-01-02", tag[0], time.UTC)
		tagUnix := tagTime.Unix()

		mm[tag[0]] = struct {
			types.Items
			unix int64
		}{
			Items: types.Items{
				TakenDateTime: tag[0],
				TotalCount:    mm[tag[0]].TotalCount + 1,
			},
			unix: tagUnix,
		}
	}

	for _, m := range mm {
		m.Items.Id = int64(len(response))
		response = append(response, m.Items)

		for i := len((response)) - 1; i > 0; i-- {
			date, _ := time.ParseInLocation("2006-01-02", response[i-1].TakenDateTime, time.UTC)
			unix := date.Unix()
			if unix < m.unix {
				temp := response[i-1]
				response[i-1] = m.Items
				response[i] = temp

				response[i-1].Id -= 1
				response[i].Id += 1
			}
		}
	}

	return response, nil
}

func GetFileInfo(bucketName string, id string) ([]interface{}, error) {
	var (
		response []interface{}

		f entity.FileInfo
		p entity.Photo
		l entity.Location

		filter = bson.D{{"fileinfo.bucketName", bucketName}}
	)

	if id != "" {
		filter = bson.D{{"fileinfo.bucketName", bucketName}, {"fileinfo.id", id}}
	}

	result, err := repository.Find(entity.DatabaseNameFileInfoName, filter)
	if err != nil {
		fmt.Println(err)
		return response, err
	}

	for _, r := range result {
		repository.Convert(r[1].Value, &f)
		if err != nil {
			fmt.Println(err)
			return response, err
		}

		mimeType := strings.Split(f.File.MimeType, "/")

		if mimeType[0] == "image" {

			if err := repository.Convert(r[2].Value, &l); err != nil {
				return response, err
			}

			if err := repository.Convert(r[3].Value, &p); err != nil {
				return response, err
			}
		}

		responseFileInfo := types.ResponseFileInfo{
			Id:              f.Id,
			File:            types.File(f.File),
			FileSystemInfo:  types.FileSystemInfo(f.FileSystemInfo),
			Name:            f.Name,
			ParentReference: types.ParentReference(f.ParentReference),
			Thumbnails: types.Thumbnails{
				Large:  types.Large(f.Thumbnails.Large),
				Medium: types.Medium(f.Thumbnails.Medium),
				Small:  types.Small(f.Thumbnails.Small),
			},
			Size: f.Size,
		}

		lang := config.Config.Location.English
		address, _ := getAddress(l.Longitude, l.Latitude, lang)

		responseLocation := types.Location{
			Address:     address,
			Altitude:    l.Altitude,
			DisplayName: l.DisplayName,
			Latitude:    l.Latitude,
			Longitude:   l.Longitude,
		}

		response = append(response, struct {
			types.ResponseFileInfo
			Location types.Location
			Photo    entity.Photo
		}{
			responseFileInfo,
			responseLocation,
			p,
		})
	}

	return response, nil
}
