package main

import (
	"github.com/kataras/iris"
	"net"
)

func main() {
	app := iris.New()
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	app.Get("/listroot", listRootHandler)
	app.Get("/list/{directory}", listHandler)
	app.Get("/download/{path}", downloadHandler)
	app.Post("/upload", uploadHandler)

	app.Run(iris.Listener(l))
}
