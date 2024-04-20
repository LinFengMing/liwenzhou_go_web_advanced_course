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

type user struct {
	id   int
	age  int
	name string
}

// 事務操作範例
func transactionDemo() {
	tx, err := db.Begin() // 開啟事務
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滾
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "Update user set age=30 where id=?"
	ret1, err := tx.Exec(sqlStr1, 2)
	if err != nil {
		tx.Rollback() // 回滾
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected() // 操作影響的行數
	if err != nil {
		tx.Rollback() // 回滾
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return

	}
	sqlStr2 := "Update user set age=40 where id=?"
	ret2, err := tx.Exec(sqlStr2, 3)
	if err != nil {
		tx.Rollback() // 回滾
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected() // 操作影響的行數
	if err != nil {
		tx.Rollback() // 回滾
		fmt.Printf("exec ret2.RowsAffected() failed, err:%v\n", err)
		return

	}
	// 當 affRow1 == 1 && affRow2 == 1 才提交事務
	fmt.Println(affRow1, affRow2)
	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("commit")
		tx.Commit() // 提交事務
	} else {
		tx.Rollback() // 回滾
		fmt.Println("affRow1 != 1 || affRow2 != 1, rollback")
	}
	fmt.Println("exec trans success")
}

func main() {
	if err := initMySQL(); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
	}
	// 做完錯誤檢查之後，確保 db 不為 nil，才能執行 defer db.Close()
	defer db.Close() // 注意这行程式碼要寫在上面 err 判断的下面
	fmt.Println("connect to db success")
	transactionDemo()
}
