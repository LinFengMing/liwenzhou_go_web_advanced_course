package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"gin_demo/models"
)

const secret = "gin_demo"

func CheckUserExist(username string) error {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用戶已存在")
	}
	return nil
}

func InserUser(user *models.User) (err error) {
	user.Password = encryptPassword(user.Password)
	fmt.Println(user.UserID)
	fmt.Println(user.Username)
	fmt.Println(user.Password)
	sqlStr := `insert into user(user_id, username, password) values(?, ?, ?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
