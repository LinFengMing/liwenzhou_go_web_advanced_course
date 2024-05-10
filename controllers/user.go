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
	logic.SignUp(p)
	// 3. 回傳 Response
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
