package info

type AllInfo struct {
	CPU CpuInfo `json:"cpu"`
	DISK AllDiskInfo `json:"disk"`
	HOST HostInfo `json:"host"`
	MEM AllMemInfo `json:"mem"`
}
