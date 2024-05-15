package controllers

import (
	"errors"
	"gin_demo/middlewares"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用戶未登錄")

func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(middlewares.CtxtUserIDKey)
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
