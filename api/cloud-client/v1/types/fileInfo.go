package types

type ResponseFileInfo struct {
	Id             string         `json:"id"`
	File           File           `json:"file"`
	FileSystemInfo FileSystemInfo `json:"fileSystemInfo"`
	Name           string         `json:"name"`
	// Location       Location       `json:"location"` // ?
	// Photo          Photo          `json:"photo"`       // ?
	ParentReference ParentReference `json:"parentReference"`
	Thumbnails      Thumbnails      `json:"thumbnails"`
	Size            float64         `json:"size"`
	// Video          Video          `json:"video"` // ?
}

type File struct {
	MimeType string `json:"mimeType"`
}

type FileSystemInfo struct {
	CreatedDateTime      uint64 `json:"createdDateTime"`
	LastModifiedDateTime uint64 `json:"lastModifiedDateTime"`
}

type Image struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

type Location struct {
	Address     Address `json:"address"`
	Altitude    float64 `json:"altitude"`
	DisplayName string  `json:"displayName"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type Address struct {
	City            string `json:"city"`
	CountryOrRegion string `json:"countryOrRegion"`
	Locality        string `json:"locality"`
	State           string `json:"state" bson:"state"`
}

type ParentReference struct {
	Path string `json:"path" bson:"path"`
}

type Photo struct {
	CameraMake          string  `json:"cameraMake" bson:"cameraMake"`
	CameraModel         string  `json:"cameraModel" bson:"cameraModel"`
	ExposureDenominator float64 `json:"exposureDenominator" bson:"exposureDenominator"`
	ExposureNumerator   float64 `json:"exposureNumerator" bson:"exposureNumerator"`
	FNumber             float64 `json:"fNumber" bson:"fNumber"`
	FocalLength         float64 `json:"focalLength" bson:"focalLength"`
	Iso                 float64 `json:"iso" bson:"iso"`
	Orientation         float64 `json:"orientation" bson:"orientation"`
	TakenDateTime       uint64  `json:"takenDateTime" bson:"takenDateTime"`
}

type Thumbnails struct {
	Large  Large  `json:"large"`
	Medium Medium `json:"medium"`
	Small  Small  `json:"small"`
}

type Large struct {
	Height float64 `json:"height"`
	Id     string  `json:"id"`
	Width  float64 `json:"width"`
}

type Medium struct {
	Height float64 `json:"height"`
	Id     string  `json:"id"`
	Width  float64 `json:"width"`
}

type Small struct {
	Height float64 `json:"height" bson:"height"`
	Id     string  `json:"id" bson:"id"`
	Width  float64 `json:"width" bson:"width"`
}

type Video struct {
	AudioBitsPerSample    float64 `json:"audioBitsPerSample" bson:"audioBitsPerSample"`
	AudioChannels         float64 `json:"audioChannels" bson:"audioChannels"`
	AudioFormat           string  `json:"audioFormat" bson:"audioFormat"`
	AudioSamplesPerSecond float64 `json:"audioSamplesPerSecond" bson:"audioSamplesPerSecond"`
	Bitrate               float64 `json:"bitrate" bson:"bitrate"`
	Duration              float64 `json:"duration" bson:"duration"`
	FourCC                string  `json:"fourCC" bson:"fourCC"`
	FrameRate             float64 `json:"frameRate" bson:"frameRate"`
	Height                float64 `json:"height" bson:"height"`
	Width                 float64 `json:"width" bson:"width"`
}

//items
type Items struct {
	Id            int64  `json:"id"`
	TakenDateTime string `json:"takenDateTime"`
	TotalCount    uint64 `json:"totalCount"`
}
