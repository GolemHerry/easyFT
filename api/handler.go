package main

import (
	"easyfiler/pkg/ft"
	"encoding/json"
	"github.com/kataras/iris"
	"io"
	"io/ioutil"
	"os"
)

const target = ":7788"

func listRootHandler(ctx iris.Context) {
	lists, err := ft.List(target, "")
	if err != nil {
		ctx.WriteString("directory not exist")
		return
	}
	data, err := json.Marshal(lists)
	ctx.Write(data)
}

func listHandler(ctx iris.Context) {
	dir := ctx.Params().Get("directory")
	lists, err := ft.List(target, dir)
	if err != nil {
		ctx.WriteString("directory not exist")
		return
	}
	data, err := json.Marshal(lists)
	if err != nil {
		ctx.WriteString("failed to marshal")
		return
	}
	ctx.Write(data)
}

func downloadHandler(ctx iris.Context) {
	path := ctx.Params().Get("path")
	ft.Download(target, path)
	file, err := os.Open("/var/log/" + path)
	if err != nil {
		ctx.WriteString("file not exist")
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		ctx.WriteString("failed to read file")
		return
	}
	ctx.Write(data)
}

func uploadHandler(ctx iris.Context) {
	if ctx.Request().Method == "POST" {
		file, head, err := ctx.FormFile("filename")
		if err != nil {
			ctx.WriteString("failed to get file")
			return
		}
		newFile, err := os.Create("/var/log/" + head.Filename)
		if err != nil {
			ctx.WriteString("failed to create file")
			return
		}
		_, err = io.Copy(newFile, file)
		if err != nil {
			ctx.WriteString("failed to save file")
			return
		}

		err = ft.Upload(target, ctx.PostValue(head.Filename))
		if err != nil {
			ctx.WriteString("failed to upload")
			return
		}
	}
}
