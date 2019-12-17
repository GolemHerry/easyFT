package filehandle

import (
	"fmt"
	"github.com/GolemHerry/easyFT/proto"
	"github.com/GolemHerry/easyFT/server/filemeta"
	"github.com/GolemHerry/easyFT/server/util"
	"io"
	"os"
	"path/filepath"
	"time"
)

type FileTransferService struct {
	Root string
}

func (fts *FileTransferService) Download(r *proto.DownloadRequest, stream proto.TransferService_DownloadServer) error {
	file, err := os.Open(filepath.Join(fts.Root, r.Name))
	if err != nil {
		return err
	}
	defer file.Close()

	var b [4096 * 1000]byte
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
		Location: "./tmp/" + r.FileName,
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	newFile, err := os.Create(fileMeta.Location)
	if err != nil {
		fmt.Printf("failed to create file, err:%v\n", err)
		return err
	}
	defer newFile.Close()

	size, err := newFile.Write(r.Data)
	fileMeta.FileSize = int64(size)
	newFile.Seek(0, 0)
	fileMeta.FileSha1 = util.FileSha1(newFile)
	ok := filemeta.UpdateFileMeta(fileMeta)
	if !ok {
		return fmt.Errorf("failed to upload")
	}
	stream.SendAndClose(&proto.UploadResponse{
		Finished: true,
		FileHash: fileMeta.FileSha1,
	})
	return nil
}
