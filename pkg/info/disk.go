package info

import "github.com/shirou/gopsutil/disk"

func GetDisk() ([]disk.PartitionStat , map[string]disk.IOCountersStat) {
	info, _ := disk.Partitions(true)  //所有分区
	info3, _ := disk.IOCounters()  //所有硬盘的io信息
	return info , info3
}

type DiskInfo struct {
	disk.PartitionStat
}

type IOCountersInfo struct {
	disk.IOCountersStat
}