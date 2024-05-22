package mysql

import "gin_demo/models"

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(post_id, author_id, community_id, title, content) values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ID, p.AuthorID, p.CommunityID, p.Title, p.Content)
	return
}

func GetPostById(pid int64) (data *models.Post, err error) {
	data = new(models.Post)
	sqlStr := `select post_id, author_id, community_id, title, content, create_time from post where post_id = ?`
	err = db.Get(data, sqlStr, pid)
	return
}

func GetPostList(page, size int64) (data []*models.Post, err error) {
	data = make([]*models.Post, 0, 2)
	sqlStr := `select post_id, author_id, community_id, title, content, create_time from post limit ?,?`
	err = db.Select(&data, sqlStr, (page-1)*size, size)
	return
}
