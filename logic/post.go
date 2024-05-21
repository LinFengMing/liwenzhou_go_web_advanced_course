package logic

import (
	"gin_demo/dao/mysql"
	"gin_demo/models"
	"gin_demo/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	// 1. 生成帖子ID
	p.ID = snowflake.GenID()
	// 2. 新增帖子
	return mysql.CreatePost(p)
}

func GetPostById(pid int64) (data *models.Post, err error) {
	return mysql.GetPostById(pid)
}
