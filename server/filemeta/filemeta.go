package filemeta

import (
	"github.com/GolemHerry/easyFT/server/db"
)

type FileMeta struct {
	FileSha1 string
	FileMD5  string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta, 10)
}

func UpdateFileMeta(f FileMeta) bool {
	return db.UploadAndFinished(f.FileSha1, f.FileName, f.Location, f.FileSize)
}

func GetFileMeta(filesha1 string) (filemeta FileMeta, err error) {
	tablefile, err := db.GetFileMeta(filesha1)
	if err != nil {
		return FileMeta{}, err
	}
	filemeta = FileMeta{
		FileSha1: tablefile.FileHash,
		FileName: tablefile.FileName.String,
		FileSize: tablefile.FileSize.Int64,
		Location: tablefile.FileAddr.String,
	}
	return
}

func DelteFileMeta(fileSha1 string) {
	db.DeleteFileMeta(fileSha1)
}
