package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initMySQL() (err error) {
	dsn := "root:1qaz2wsx3edc@tcp(127.0.0.1:3306)/liwenzhou_go_web_advanced_course"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("connect to db failed, err:%v\n", err)
		return
	}
	db.SetConnMaxLifetime(time.Second * 10)
	db.SetMaxOpenConns(200) // 最大連接數
	db.SetMaxIdleConns(10)  // 最大閒置連接數
	return
}

func main() {
	if err := initMySQL(); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
	}
	// 做完錯誤檢查之後，確保 db 不為 nil，才能執行 defer db.Close()
	defer db.Close() // 注意这行程式碼要寫在上面 err 判断的下面
	fmt.Println("connect to db success")
}
