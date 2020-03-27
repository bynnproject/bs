package info

import (
	"github.com/shirou/gopsutil/host"
)

func GetHosts() *host.InfoStat {
	info, _ := host.Info()
	return info
}

type HostInfo struct {
	host.InfoStat
}