package logic

import (
	"gin_demo/dao/mysql"
	"gin_demo/models"
	"gin_demo/pkg/snowflake"
)

func SignUp(p *models.ParamSigUp) (err error) {
	// 判斷用戶是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 生成 UID
	userID := snowflake.GenID()
	// 創建一個 User 實體
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 儲存至資料庫
	return mysql.InserUser(user)
}

func Login(p *models.ParamLogin) error {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.Login(user)
}
