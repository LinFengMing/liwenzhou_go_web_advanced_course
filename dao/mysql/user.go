package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"gin_demo/models"
)

const secret = "gin_demo"

var (
	ErroeUserExist       = errors.New("用戶已存在")
	ErroeUserNotExist    = errors.New("用戶不存在")
	ErroeInvalidPassword = errors.New("用戶名或密碼錯誤")
)

func CheckUserExist(username string) error {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErroeUserExist
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

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id, username, password from user where username = ?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErroeUserNotExist
	}
	if err != nil {
		return err
	}
	// 判斷密碼是否正確
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErroeInvalidPassword
	}
	return
}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
