package info

import "github.com/shirou/gopsutil/mem"

func GetMem() (*mem.VirtualMemoryStat , *mem.SwapMemoryStat) {
	virInfo, _ := mem.VirtualMemory()
	swapInfo, _ := mem.SwapMemory()
	return virInfo , swapInfo
}

type VMemInfo struct {
	mem.VirtualMemoryStat
}

type SMemInfo struct {
	mem.SwapMemoryStat
}