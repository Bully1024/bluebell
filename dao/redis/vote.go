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
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

func TestZAdd(postID int64) (err error) {
	err = client.Set("123", postID, 0).Err()
	return
}

func CreatePost(postID int64) (err error) {
	pipeline := client.TxPipeline()
	//帖子时间
	err = pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	}).Err()
	if err != nil {
		return err
	}
	//帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	_, err = pipeline.Exec()
	return err
}

func VoteForPost(userID string, postID string, value float64) (err error) {
	//1.判断投票的限制
	//取redis取帖子发布时间
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSecond {
		//不允许投票
		return ErrVoteTimeExpire
	}
	//2.需要放到一个pipeline事务中操作
	//2.更新帖子分数
	//先查当前用户给当前帖子的投票记录
	key := getRedisKey(KeyPostVotedZSetPrefix + postID)
	ov := client.ZScore(key, userID).Val()

	//更新：如果这一次投票的值和之前保存的值一致，就提示不允许重复投票
	if value == ov {
		return ErrVoteRepeated
	}

	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}

	diff := math.Abs(ov - value)    //计算俩次投票的差值
	pipeline := client.TxPipeline() //开启事务
	_, err = pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID).Result()
	if err != nil {
		return err
	}
	//3，记录用户为该帖子投过票
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPrefix+postID), postID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPrefix+postID), redis.Z{
			Score:  value, //赞成票还是反对票
			Member: userID,
		})
	}
	_, err = pipeline.Exec()
	return err
}
