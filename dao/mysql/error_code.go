package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("用戶已存在")
	ErrorUserNotExist    = errors.New("用戶不存在")
	ErrorInvalidPassword = errors.New("用戶名或密碼錯誤")
	ErrorInvalidID       = errors.New("無效的ID")
)
