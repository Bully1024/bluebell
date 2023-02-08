package logic

import (
	"GoWebCode/bluebell/dao/redis"
	"GoWebCode/bluebell/models"
	"strconv"

	"go.uber.org/zap"
)

/*
投票的几种情况
direction = 1,有俩种情况：
	1，之前没有投过票，现在投赞成票	-->更新分数和投票纪录 差值的绝对值：1 +432
	2，之前投反对票，现在改投赞成票	-->更新分数和投票纪录 差值的绝对值：2 +432*2
direction= 0，有俩种情况：
	1，之前投反对票，现在取消投票	-->更新分数和投票纪录 差值的绝对值：1 +432
	2，之前投赞成票，现在取消投票	-->更新分数和投票纪录 差值的绝对值：1 -432
direction= -1，有俩种情况：
	1，之前没有投过票，现在改投反对票	-->更新分数和投票纪录 差值的绝对值：1 -432
	2，之前投赞成票，现在投改投反对票	-->更新分数和投票纪录 差值的绝对值：2 -432*2

投票的限制：
每个帖子自发表之日起一星期之内允许用户投票，超过一个星期就不允许再投票了
	1，到期之后将redis中保存的赞成票及反对票存储到mysql中
	2，到期之后删除KeyPostVotedZSetPrefix
*/
// 投票功能

//本项目使用简化版的投票分数
//投一票就加432分 根据时间戳：86400/200 -> 需要200张赞成票才能可以给你的帖子续一天 ->redis实战

// VoteForPost 给帖子投票函数
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
