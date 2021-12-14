package source

// request source
type RequestSource struct {
	Token  string   `json:"token" binding:"required"`
	Type   string   `json:"type" binding:"required"`
	Target []string `json:"target" binding:"required"`
}

// response source
type ResponseSource struct {
	Target interface{} `json:"target"`
}
