package info

import "github.com/shirou/gopsutil/mem"

func GetMem() AllMemInfo {
	virInfo, _ := mem.VirtualMemory()
	swapInfo, _ := mem.SwapMemory()
	return AllMemInfo {
		swapInfo,
		virInfo,
	}
}

type VMemInfo struct {
	mem.VirtualMemoryStat
}

type SMemInfo struct {
	mem.SwapMemoryStat
}

type AllMemInfo struct {
	Sinfo *mem.SwapMemoryStat
	Vinfo *mem.VirtualMemoryStat
}