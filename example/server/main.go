package main

import (
	"easyfiler"
	"easyfiler/pkg/db/mysql"
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

func main() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("fail to load config, err: %v", err)
		os.Exit(1)
	}
	withdb, _ := cfg.Section("server").Key("withdb").Bool()
	if withdb {
		err = mysql.InitDB(mysql.DBconf{
			User:     cfg.Section("mysql").Key("user").String(),
			Password: cfg.Section("mysql").Key("password").String(),
			Host:     cfg.Section("mysql").Key("host").String(),
			DB:       cfg.Section("mysql").Key("db").String(),
		})
		if err != nil {
			fmt.Printf("failed to connect mysql err:%v\n", err)
			os.Exit(1)
		}
	}
	srv := &easyfiler.Server{
		Port:   cfg.Section("server").Key("port").String(),
		Root:   cfg.Section("server").Key("root").String(),
		WithDB: withdb,
	}

	err = srv.Start()
	if err != nil {
		fmt.Println("failed to start")
		os.Exit(1)
	}
}
