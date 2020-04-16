package info

import (
	pb "bs/protoc"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
	"bs/pkg/info"
)

type InfoClient struct {
	Address string ;
}

func (cli *InfoClient)GetInfo() (error , info.AllInfo)  {
	conn , err := grpc.Dial(cli.Address ,  grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("after conn")
	defer conn.Close()
	out := pb.NewInfoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req , err := out.GetInfo(ctx)
	allinfo := info.AllInfo{}
	if err != nil {
		log.Fatal(err)
		return err,allinfo
	}
	// /Users/bynn/uploadFileTest.txt

	err = req.Send(&pb.InfoReq{
		Msg:                  "",
	})
	if err != nil {
		log.Fatal("client get info err %v" , err)
		return err,allinfo

	}

	resp , err := req.Recv()
	cpuinfo := resp.Cpu
	meninfo := resp.Mem
	hostinfo := resp.Host
	diskinfo := resp.Disk
	var cpu info.CpuInfo
	var mem info.AllMemInfo
	var host info.HostInfo
	var disk info.AllDiskInfo
	if err := json.Unmarshal([]byte(cpuinfo), &cpu); err == nil {
		allinfo.CPU = cpu
		fmt.Println(cpu)

	} else {
		fmt.Println(err)
	}
	if err := json.Unmarshal([]byte(meninfo), &mem); err == nil {
		allinfo.MEM = mem
		fmt.Println(mem)
	} else {
		fmt.Println(err)
	}
	if err := json.Unmarshal([]byte(hostinfo), &host); err == nil {
		allinfo.HOST = host
		fmt.Println(host)
	} else {
		fmt.Println(err)
	}
	if err := json.Unmarshal([]byte(diskinfo), &disk); err == nil {
		allinfo.DISK = disk
		fmt.Println(disk)
	} else {
		fmt.Println(err)
	}
	return nil , allinfo
}
