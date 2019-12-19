package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	dsn := "root:golem@tcp(47.100.114.83)/vigny"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Failed to open connection to mysql, error:%v\n", err)
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("Failed to connect to mysql, error:%v\n", err)
		os.Exit(1)
	}
	fmt.Println("connected to mysql")
	db.SetMaxOpenConns(1000)
}

func DBConn() *sql.DB {
	return db
}
