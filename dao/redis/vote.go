package redis

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
)

const (
	oneWeekInSecond = 7 * 24 * 3600
	scorePerVote    = 432 //每一票的分数
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
)

func CreatePost(postID int64) error {
	pipeline := client.TxPipeline()
	//帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	//帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	_, err := pipeline.Exec()
	return err
}

func VoteForPost(userID, postID string, value float64) error {
	//1.判断投票的限制
	//取redis取帖子发布时间
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSecond {
		return ErrVoteTimeExpire
	}
	//2,3需要放到一个pipeline事务中操作
	//2，更新帖子分数
	//先查当前用户给当前帖子的投票记录
	ov := client.ZScore(getRedisKey(KeyPostVotedZSetPrefix+postID), userID).Val()
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) //计算俩次投票的差值
	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostTimeZSet), op*diff*scorePerVote, postID)
	//3，记录用户为该帖子投过票
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPrefix+postID), userID)
	}
	pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPrefix+postID), redis.Z{
		Score:  value, //赞成票还是反对票
		Member: userID,
	})
	_, err := pipeline.Exec()
	return err
}
