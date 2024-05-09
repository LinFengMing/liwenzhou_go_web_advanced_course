package logic

import (
	"gin_demo/dao/mysql"
	"gin_demo/models"
	"gin_demo/pkg/snowflake"
)

func SignUp(p *models.ParamSigUp) {
	// 1.判斷用戶是否存在
	mysql.QueryUseByUsername()
	// 2.生成 UID
	snowflake.GenID()
	// 3.儲存至資料庫
	mysql.InserUser()
}
