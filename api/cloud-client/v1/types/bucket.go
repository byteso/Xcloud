package types

import "io"

type RequestCreateBucket struct {
	BucketName string `json:"bucketName" binding:"required"`
}

type ResponseGetSource struct {
	ContentType   string
	ContentLength int64
	ExtraHeaders  map[string]string
	Reader        io.Reader
}

type RequestUploadSource struct {
	FolderName string `json:"folderName"`
}

type RequestDownloadSource struct {
	FolderName string `json:"folderName" binding:"required"`
	Key        string `json:"key" binding:"required"`
}

type ResponseDownloadSource struct {
	ContentType string
	Data        []byte
}

type RequestDeleteSource struct {
	FolderName string `json:"folderName" binding:"required"`
	Key        string `json:"key" binding:"required"`
}

type ResponseDeleteSource struct {
}

type UploadSource struct {
	Name string
	Path string
	Size int64
	Data io.Reader
}
