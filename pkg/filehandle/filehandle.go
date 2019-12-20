package filehandle

import (
	"easyfiler/pkg/filemeta"
	"easyfiler/pkg/proto"
	"easyfiler/pkg/util"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type FileTransferService struct {
	Root   string
	WithDB bool
}

func (fts *FileTransferService) List(r *proto.ListRequest, stream proto.TransferService_ListServer) error {
	err := filepath.Walk(fts.Root+r.Directory, func(p string, info os.FileInfo, err error) error {
		name, err := filepath.Rel(fts.Root, p)
		if err != nil {
			return err
		}
		name = filepath.ToSlash(name)
		f := &proto.ListResponse{
			Name: filepath.ToSlash(name),
			Size: info.Size(),
			Mode: uint32(info.Mode()),
		}
		return stream.Send(f)
	})
	return err
}

func (fts *FileTransferService) Download(r *proto.DownloadRequest, stream proto.TransferService_DownloadServer) error {
	file, err := os.Open(fts.Root + r.Name)
	if err != nil {
		fmt.Println("no file", err)
		return err
	}
	defer file.Close()

	var b [1024 * 1024]byte
	for {
		n, err := file.Read(b[:])
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		err = stream.Send(&proto.DownloadResponse{
			Data: b[:n],
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (fts *FileTransferService) Upload(stream proto.TransferService_UploadServer) error {
	r, err := stream.Recv()
	if err != nil {
		fmt.Printf("failed to recive file, err:%v\n", err)
		return err
	}
	fileMeta := filemeta.FileMeta{
		Location: fts.Root + r.FileName,
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	newFile, err := os.Create(fileMeta.Location)
	if err != nil {
		fmt.Printf("failed to create file, err:%v\n", err)
		return err
	}
	defer newFile.Close()

	size, err := newFile.Write(r.Data)
	if fts.WithDB {
		fileMeta.FileSize = int64(size)
		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		ok := filemeta.UpdateFileMeta(fileMeta)
		if !ok {
			fmt.Println("already exist")
			return fmt.Errorf("failed to upload")
		}
	}
	stream.SendAndClose(&proto.UploadResponse{
		Finished: true,
		FileHash: fileMeta.FileSha1,
	})
	return nil
}
