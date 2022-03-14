package types

type ServerInfo struct {
	CentralProcessingUnit CentralProcessingUnit `json:"CentralProcessingUnit"`
	Memory                Memory                `json:"memory"`
	Disk                  Disk                  `json:"disk"`
	Net                   Net                   `json:"net"`
}

type CentralProcessingUnit struct {
	Percent float64 `json:"percent"`
}

type Memory struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"usedPercent"`
}

type Disk struct {
	Total float64 `json:"total"`
	Free  float64 `json:"free"`
}

type Net struct {
}
