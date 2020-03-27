package info

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
)

func GetCpu() []cpu.InfoStat {
	info, _ := cpu.Info()  //总体信息
	fmt.Println(info)
	return info
}

type CpuInfo struct {
	cpu.InfoStat
}