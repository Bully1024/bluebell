package redis

//redis key

//redis key尽量使用命名空间的方式，区分不同的key，方便查询和拆分

const (
	KeyPreFix              = "bluebell:"
	KeyPostTimeZSet        = "post:time"   //zset:帖子及发帖时间
	KeyPostScoreZSet       = "post:score"  //zset:帖子及投票的分数
	KeyPostVotedZSetPrefix = "post:voted:" //zset:记录用户及投票的类型;参数是post_id

	KeyCommunitySetPrefix = "community:" //set:保存每个分区下帖子的id
)

// 给redis key加上前缀
func getRedisKey(key string) string {
	return KeyPreFix + key
}
