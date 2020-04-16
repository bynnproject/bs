package info

import (
	"github.com/shirou/gopsutil/host"
)

func GetHosts() HostInfo {
	info, _ := host.Info()
	return HostInfo{info}
}

type HostInfo struct {
	Host *host.InfoStat
}