package logic

import (
	"GoWebCode/bluebell/dao/mysql"
	"GoWebCode/bluebell/models"
	"GoWebCode/bluebell/pkg/snowflake"
	"fmt"

	"go.uber.org/zap"
)

func CreatPost(p *models.Post) (err error) {
	//1.生成post id
	p.ID = snowflake.GenID()
	//2.保存到数据库
	//3.返回
	return mysql.CreatePost(p)
}

// GetPostId 根据帖子ID查询帖子详情数据
func GetPostId(pid int64) (data *models.ApiPostDetail, err error) {
	//查询并拼接我们接口想要的数据
	post, err := mysql.GetPostByID(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(pid) failed",
			zap.Int64("pid", pid),
			zap.Error(err))
		return
	}

	//根据作者id查询作者信息
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById(post.AuthorID)",
			zap.Int64("authorID", post.AuthorID),
			zap.Error(err))
		return
	}
	//根据社区id查询社区详细信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetUserById(post.AuthorID)",
			zap.Int64("communityID", post.CommunityID),
			zap.Error(err))
		return
	}
	//接口数据拼接
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

// 获取帖子列表
func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		fmt.Println("6453542346")
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))

	for _, post := range posts {
		//根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID)",
				zap.Int64("authorID", post.AuthorID),
				zap.Error(err))
			continue
		}
		//根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID)",
				zap.Int64("communityID", post.CommunityID),
				zap.Error(err))
			continue
		}
		postdetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postdetail)
		fmt.Println(data)
	}

	return
}
