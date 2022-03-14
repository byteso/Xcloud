package entity

var (
	DatabaseNameFileInfoName = "fileInfo"
)

type FileInfo struct {
	BucketName     string         `json:"bucketName" bson:"bucketName"`
	Id             string         `json:"id" bson:"id"`
	File           File           `json:"file" bson:"file"`
	FileSystemInfo FileSystemInfo `json:"fileSystemInfo" bson:"fileSystemInfo"`
	Name           string         `json:"name" bson:"name"`
	// Location       Location       `json:"location" bson:"location"` // ?
	// Photo      Photo      `json:"photo" bson:"photo"` // ?
	ParentReference ParentReference `json:"parentReference" bson:"parentReference"`
	Thumbnails      Thumbnails      `json:"thumbnails" bson:"thumbnails"`
	Size            float64         `json:"size" bson:"size"`
	// Video      Video      `json:"video" bson:"video"` // ?
}

type File struct {
	MimeType string `json:"mimeType" bson:"mimeType"`
}

type FileSystemInfo struct {
	CreatedDateTime      uint64 `json:"createdDateTime"`
	LastModifiedDateTime uint64 `json:"lastModifiedDateTime"`
}

type Image struct {
	Height float64 `json:"" bson:"height"`
	Width  float64 `json:"" bson:"width"`
}

type Location struct {
	// Address     Address `json:"address" bson:"address"`
	Altitude    float64 `json:"altitude" bson:"altitude"`
	DisplayName string  `json:"displayName" bson:"displayName"`
	Latitude    float64 `json:"latitude" bson:"latitude"`
	Longitude   float64 `json:"longitude" bson:"longitude"`
}

type Address struct {
	City            string `json:"city" bson:"city"`
	CountryOrRegion string `json:"countryOrRegion" bson:"countryOrRegion"`
	Locality        string `json:"locality" bson:"locality"`
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
	Large  Large  `json:"large" bson:"large"`
	Medium Medium `json:"medium" bson:"medium"`
	Small  Small  `json:"small" bson:"small"`
}

type Large struct {
	Height float64 `json:"height" bson:"height"`
	Id     string  `json:"id" bson:"id"`
	Width  float64 `json:"width" bson:"width"`
}

type Medium struct {
	Height float64 `json:"height" bson:"height"`
	Id     string  `json:"id" bson:"id"`
	Width  float64 `json:"width" bson:"width"`
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
