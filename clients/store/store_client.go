package store

import (
	pb "bs/protoc"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"time"
	"context"
)

type UploadFileClient struct {
	Address string ;
}

func NewUploadFileClient(address string) *UploadFileClient  {
	log.Println("new client")
	return &UploadFileClient{Address:address}
}

func (cli *UploadFileClient)UploadFile(filePath string , fileName string)  error{
	log.Println("upload file %v" , cli.Address)
	conn , err := grpc.Dial(cli.Address ,  grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("after conn")
	defer conn.Close()
	out := pb.NewUploadClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req , err := out.UploadFile(ctx)

	file , err := ioutil.ReadFile(filePath)
	log.Println(len(file))
	if err != nil {
		log.Fatal(err)
		return err
	}
	// /Users/bynn/uploadFileTest.txt

	err = req.Send(&pb.UploadFileReq{
		Name: fileName,
		File: file,
	})
	if err != nil {
		log.Fatal("upload file err %v" , err)
		return err
	}

	return nil

}

func (cli *UploadFileClient)DownloadFile(source string , target string , tag string)  error{

	conn , err := grpc.Dial(cli.Address ,  grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("after conn")
	defer conn.Close()
	out := pb.NewUploadClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req , err := out.DoloadFile(ctx)

	err = req.Send(&pb.DownloadFileReq{
		Path: source,
	})
	if err != nil {
		log.Println(err)
		return err
	}

	resp , err  := req.Recv()
	if err != nil {
		log.Println(err)
		return err
	}
	file := resp.File
	err = ioutil.WriteFile(target,file,0644)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}