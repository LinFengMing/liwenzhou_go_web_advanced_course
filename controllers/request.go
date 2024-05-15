package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const CtxtUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用戶未登錄")

func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxtUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
