package database

import (
	"log"

	"github.com/byteso/Xcloud/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	ObjectClient *minio.Client
)

func InitObjectServerEngine() {
	ObjectClient = initObjectServer()
}

func initObjectServer() *minio.Client {
	c := config.Config

	minioClient, err := minio.New(c.ObjectServer.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.ObjectServer.AccessKeyID, c.ObjectServer.SecretAccessKey, ""),
		Secure: c.ObjectServer.UseSSL,
	})
	if err != nil {
		log.Fatal(err)
	}

	return minioClient
}
