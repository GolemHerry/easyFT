package main

import (
	"context"
	"fmt"
	"github.com/GolemHerry/easyFT/proto"
	"google.golang.org/grpc"
	"io/ioutil"
)

func main() {
	conn, err := grpc.Dial(":7788", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewTransferServiceClient(conn)

	filename := "test.txt"
	b, err := ioutil.ReadFile("./" + filename)
	fmt.Println("data is ", string(b))
	reqStreamData := &proto.UploadRequest{
		FileName: filename,
		Data:     b,
	}
	putRes, _ := c.Upload(context.Background())
	err = putRes.Send(reqStreamData)
	if err != nil {
		fmt.Printf("failed to upload file err:%v\n", err)
	}
	fmt.Println("succeed")
}
