package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var mysqldb *sql.DB

func init() {
	var err error
	dsn := "${user}:${pwd}@tcp(${host})/${db}"
	mysqldb, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Failed to open connection to mysql, error:%v\n", err)
		os.Exit(1)
	}
	err = mysqldb.Ping()
	if err != nil {
		fmt.Printf("Failed to connect to mysql, error:%v\n", err)
		os.Exit(1)
	}
	fmt.Println("connected to mysql")
	mysqldb.SetMaxOpenConns(1000)
}

func DBConn() *sql.DB {
	return mysqldb
}
