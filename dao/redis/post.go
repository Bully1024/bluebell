package redis

import (
	"GoWebCode/bluebell/models"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func getIDSFromKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1
	return client.ZRevRange(key, start, end).Result()
}

// GetPostIDsInOrder 根据时间或者投票数有序查询ids
func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	//从redis中获取id
	//1.根据用户请求中携带的order参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	//2.确定查询的索引起始点
	//3.ZREVRANGE查询 按分数从大到小顺序查询指定数量
	return getIDSFromKey(key, p.Page, p.Size)
}

// GetPostVoteDatas 根据ids查询每篇帖子当前投票数
func GetPostVoteDatas(ids []string) (data []int64, err error) {
	//data = make([]int64, 0, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedZSetPrefix + id)
	//	//查找key中分数是1的元素的数量->统计每篇帖子的赞成票的数量
	//	v := client.ZCount(key, "1", "1").Val()
	//	data = append(data, v)
	//}
	//使用pipeline一次发送多条命令，减少RTT
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPrefix + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

// GetCommunityPostIDsInOrder 按社区根据时间或者投票数有序查询ids
func GetCommunityPostIDsInOrder(p *models.ParamCommunityPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}

	//使用zinterstore 把分区的帖子set与帖子分数的zset生成一个新的zset
	//根据新的zset按之前的逻辑取数据

	//社区的key
	communityKey := getRedisKey(KeyCommunitySetPrefix + strconv.Itoa(int(p.CommunityID)))
	//利用缓存key减少zinterstore执行的次数！！！！
	key := orderKey + strconv.Itoa(int(p.CommunityID))
	if client.Exists(key).Val() < 1 {
		//不存在，需要计算
		pipelines := client.Pipeline()
		pipelines.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, communityKey, orderKey) //zinterstore计算
		pipelines.Expire(key, 60*time.Second) //设置超时时间
		_, err := pipelines.Exec()
		if err != nil {
			return nil, err
		}
	}
	//存在的话直接根据kay查询
	return getIDSFromKey(key, p.Page, p.Size)
}
