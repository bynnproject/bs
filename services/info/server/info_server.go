package server

import (
	"bs/pkg/info"
	pb "bs/protoc"
	"bytes"
	"encoding/json"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

type InfoServer struct {
	Port string
	pb.UnimplementedInfoServer
}

func (s *InfoServer)GetInfo(stream pb.Info_GetInfoServer) error{

	log.Println("into func")
	req , err := stream.Recv()
	log.Println(req)
	if err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}
	d ,_ := json.Marshal(info.GetDisk())
	h ,_ := json.Marshal(info.GetHosts())
	c ,err := json.Marshal(info.GetCpu())
	m ,_ := json.Marshal(info.GetMem())




	err = stream.Send(&pb.InfoResp{
		Host:    string(h),
		Cpu:     string(c),
		Mem:     string(m),
		Disk:    string(d),
	})

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (app *InfoServer)Run() error  {
	lis, err := net.Listen("tcp" , ":" + app.Port)
	log.Println("into server run")
	if err != nil {
		log.Fatal(err)
		return err
	}

	s := grpc.NewServer()
	pb.RegisterInfoServer(s , app)
		url := "http://127.0.0.1:8080/node/start?address=127.0.0.1:"+app.Port
		log.Println(url)
		Get(url)
		log.Println(url)
	err = s.Serve(lis)

	return nil
}


func Get(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}