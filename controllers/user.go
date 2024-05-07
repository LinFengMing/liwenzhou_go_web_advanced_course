package controllers

import (
	"gin_demo/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	// 1. 驗證參數
	// 2. 業務邏輯處理
	logic.SignUp()
	// 3. 回傳 Response
	c.JSON(http.StatusOK, "ok")
}
