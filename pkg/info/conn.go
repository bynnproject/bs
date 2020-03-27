package info

import "github.com/shirou/gopsutil/net"

func GetConn() []net.ConnectionStat {
	info, _ := net.Connections("all")
	return info
}


type ConnInfo struct {
	net.ConnectionStat
}

