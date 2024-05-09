package controllers

import (
	"fmt"
	"gin_demo/logic"
	"gin_demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	// 1. 驗證參數
	p := new(models.ParamSigUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	// 參數規則驗證
	if len(p.Username) == 0 || len(p.Password) == 0 || p.Password != p.RePassword {
		zap.L().Error("SignUp with invalid param")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "請求參數有誤",
		})
		return
	}
	fmt.Println(p)
	// 2. 業務邏輯處理
	logic.SignUp(p)
	// 3. 回傳 Response
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
