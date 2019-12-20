package main

import "easyfiler/pkg/ft"

func main() {
	ft.Download(":7788", "wifi.log.4.bz2")
	ft.List(":7788", "displaypolicy/")
}
