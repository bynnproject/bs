package info

import "github.com/shirou/gopsutil/process"

func GetProcesses() []int32 {
	info, _ := process.Pids()
	return info
}

func GetProcessInfo() []*process.Process {
	info2,_ := process.Processes();
	return info2
}


