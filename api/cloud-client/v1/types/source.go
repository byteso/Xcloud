package types

// request source
type RequestSource struct {
	FolderName string `json:"folderName" binding:"required"`
}

// response source
type ResponseSource struct {
	Target interface{} `json:"target"`
}
