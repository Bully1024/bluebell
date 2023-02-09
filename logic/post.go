package logic

import (
	"GoWebCode/bluebell/dao/mysql"
	"GoWebCode/bluebell/dao/redis"
	"GoWebCode/bluebell/models"
	"GoWebCode/bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatPost(p *models.Post) (err error) {
	//1.生成post id
	p.ID = snowflake.GenID()
	//2.保存到数据库
	//3.返回
	err = mysql.CreatePost(p)
	if err != nil {
		zap.L().Error("mysql.CreatPost(&p) failed", zap.Error(err))
		return err
	}
	//ToDo BUG!!!!!!!!!!!!!!!!!!!!!!!!!
	if err = redis.CreatePost(p.ID); err != nil {
		zap.L().Error("redis.CreatePost failed", zap.Error(err))
		return err
	}
	return
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

// GetPostList 获取帖子列表
func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
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
	}

	return
}

func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	//2.去redis查询id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder() return 0 data")
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("ids", ids))
	//3.根据id去Mysql数据库查询帖子的详细信息
	//返回的数据还要按照我给定的id的顺序返回->mysql:order by
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("posts", posts))
	//将帖子的作者以及分区信息查询出来填充到帖子
	//ids和posts相对应，没必要在循环中查询，提前查询好数据
	voteData, err := redis.GetPostVoteDatas(ids)
	if err != nil {
		return
	}
	for idx, post := range posts {
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
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postdetail)
	}
	return
}
