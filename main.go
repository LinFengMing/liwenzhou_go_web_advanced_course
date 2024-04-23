package main

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB

type user struct {
	ID   int    `db:"id"`
	Age  int    `db:"age"`
	Name string `db:"name"`
}

func (u user) Value() (driver.Value, error) {
	return []driver.Value{u.Name, u.Age}, nil
}

func BatchInsertUsers(users []interface{}) error {
	query, args, _ := sqlx.In(
		"INSERT INTO user (name, age) VALUES (?), (?), (?)",
		users...,
	)
	fmt.Println(query)
	fmt.Println(args)
	_, err := db.Exec(query, args...)
	return err
}

func BatchInsertUsers2(users []user) error {
	_, err := db.NamedExec("INSERT INTO user (name, age) VALUES (:name, :age)", users)
	return err
}

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

func QueryByIDs(ids []int) (users []user, err error) {
	query, args, err := sqlx.In("SELECT id, name, age FROM user WHERE id IN (?)", ids)
	if err != nil {
		return
	}
	query = db.Rebind(query)
	err = db.Select(&users, query, args...)
	return
}

func QueryAndOrderByIDs(ids []int) (users []user, err error) {
	strIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		strIDs = append(strIDs, fmt.Sprintf("%d", id))
	}
	query, args, err := sqlx.In("SELECT id, name, age FROM user WHERE id IN (?) ORDER BY FIND_IN_SET(id, ?)", ids, strings.Join(strIDs, ","))
	if err != nil {
		return
	}
	query = db.Rebind(query)
	err = db.Select(&users, query, args...)
	return
}

func main() {
	if err := initMySQL(); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to db success")
	// u1 := user{Name: "terry", Age: 18}
	// u2 := user{Name: "tom", Age: 20}
	// u3 := user{Name: "jerry", Age: 22}
	// users := []interface{}{u1, u2, u3}
	// BatchInsertUsers(users)
	users, err := QueryByIDs([]int{4, 5, 6})
	if err != nil {
		fmt.Printf("QueryByIDs failed, err:%v\n", err)
		return
	}
	for _, user := range users {
		fmt.Printf("user:%#v\n", user)
	}
	users, err = QueryAndOrderByIDs([]int{7, 8, 4, 2})
	if err != nil {
		fmt.Printf("QueryByIDs failed, err:%v\n", err)
		return
	}
	for _, user := range users {
		fmt.Printf("user:%#v\n", user)
	}
}
