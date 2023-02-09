package redis

import "GoWebCode/bluebell/models"

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	//从redis中获取id
	//1.根据用户请求中携带的order参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	//2.确定查询的索引起始点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	//3.ZREVRANGE查询 按分数从大到小顺序查询指定数量
	return client.ZRevRange(key, start, end).Result()
}
