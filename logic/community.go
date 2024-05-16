package logic

import (
	"gin_demo/dao/mysql"
	"gin_demo/models"
)

func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}
