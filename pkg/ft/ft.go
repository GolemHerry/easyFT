package ft

import (
	"context"
	"easyfiler/pkg/proto"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
)

func List(target, dir string) ([]proto.ListResponse, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	c := proto.NewTransferServiceClient(conn)
	reqData := &proto.ListRequest{
		Directory: dir,
	}
	listClient, err := c.List(context.Background(), reqData)
	var lists = make([]proto.ListResponse, 0)
	if err != nil {
		return nil, err
	}
	for {
		res, err := listClient.Recv()
		if err != nil {
			break
		}
		lists = append(lists, *res)
	}
	return lists, nil
}

func Upload(target, filename string) error {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
		return err
	}
	defer conn.Close()

	c := proto.NewTransferServiceClient(conn)

	b, err := ioutil.ReadFile("./" + filename)
	upReqStreamData := &proto.UploadRequest{
		FileName: filename,
		Data:     b,
	}

	upClient, err := c.Upload(context.Background())
	if err != nil {
		return err
	}
	err = upClient.Send(upReqStreamData)
	if err != nil {
		fmt.Printf("failed to upload file err:%v\n", err)
		return err
	}
	return nil
}

func Download(target, filename string) error {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
		return err
	}
	defer conn.Close()

	c := proto.NewTransferServiceClient(conn)

	reqData := &proto.DownloadRequest{
		Name: filename,
	}

	dnClient, err := c.Download(context.Background(), reqData)
	if err != nil {
		return err
	}
	_, err = os.Create(filename)
	if err != nil {
		fmt.Printf("failed to create file, err:%v\n", err)
		return err
	}
	for {
		res, err := dnClient.Recv()
		if err != nil {
			break
		}
		err = ioutil.WriteFile(filename, res.Data, 0666)
		if err != nil {
			fmt.Printf("failed to write file, err:%v\n", err)
			return err
		}
	}
	return nil
}
