package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB

func initMySQL() (err error) {
	dsn := "root:1qaz2wsx3edc@tcp(127.0.0.1:3306)/liwenzhou_go_web_advanced_course?charset=utf8mb4&parseTime=True"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect to db failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(200) // 最大連接數
	db.SetMaxIdleConns(10)  // 最大閒置連接數
	return
}

func main() {
	if err := initMySQL(); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to db success")
}
