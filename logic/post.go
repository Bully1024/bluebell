package logic

import (
	"GoWebCode/bluebell/dao/mysql"
	"GoWebCode/bluebell/models"
	"GoWebCode/bluebell/pkg/snowflake"
)

func CreatPost(p *models.Post) (err error) {
	//1.生成post id
	p.ID = snowflake.GenID()
	//2.保存到数据库
	//3.返回
	return mysql.CreatePost(p)
}

// GetPostId 根据帖子ID查询帖子详情数据
func GetPostId(pid int64) (data *models.Post, err error) {
	return mysql.GetPostByID(pid)
}
