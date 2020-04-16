package info

import (
	"github.com/shirou/gopsutil/cpu"
)

func GetCpu() CpuInfo {
	info, _ := cpu.Info()  //总体信息
	return CpuInfo{Cpus:info}
}

//type CpuInfo []cpu.InfoStat
type CpuInfo struct {
	Cpus []cpu.InfoStat `json:"cpus"`
}

