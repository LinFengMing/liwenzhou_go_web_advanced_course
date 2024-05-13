package controllers

import (
	"fmt"
	"gin_demo/logic"
	"gin_demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	// 1. 驗證參數
	p := new(models.ParamSigUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判斷 err 是不是 validator.ValidationErrors 類型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	fmt.Println(p)
	// 2. 業務邏輯處理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "註冊失敗",
		})
		return
	}
	// 3. 回傳 Response
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func LoginHandler(c *gin.Context) {
	// 1. 取得參數及驗證
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Login with invalid param", zap.String("username", p.Username), zap.Error(err))
		// 判斷 err 是不是 validator.ValidationErrors 類型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	// 2. 業務邏輯處理
	if err := logic.Login(p); err != nil {
		zap.L().Error("logic.Login failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "會員名或密碼錯誤",
		})
		return
	}
	// 3. 回傳 Response
	c.JSON(http.StatusOK, gin.H{
		"msg": "登陸成功",
	})
}
