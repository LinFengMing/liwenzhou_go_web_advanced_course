package controllers

import (
	"gin_demo/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHundler(c *gin.Context) {
	// 查詢社區列表
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
