package service

import (
	"context"

	"github.com/byteso/Xcloud/internal/database"
	"github.com/minio/minio-go/v7"
)

func ListBuckets() ([]minio.BucketInfo, error) {
	client := database.ObjectClient

	buckets, err := client.ListBuckets(context.Background())
	if err != nil {
		return buckets, err
	}

	return buckets, nil
}
