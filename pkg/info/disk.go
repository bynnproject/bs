package info

import "github.com/shirou/gopsutil/disk"

func GetDisk() AllDiskInfo {
	info, _ := disk.Partitions(true)  //所有分区
	info3, _ := disk.IOCounters()  //所有硬盘的io信息

	return AllDiskInfo{
		info ,
		info3,
	}
}

type DiskInfo struct {
	disk.PartitionStat
}

type IOCountersInfo struct {
	disk.IOCountersStat
}

type AllDiskInfo struct {
	Info []disk.PartitionStat
	Ioinfo map[string]disk.IOCountersStat
}