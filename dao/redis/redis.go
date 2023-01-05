package redis

import (
	"GoWebCode/bluebell/settings"
	"fmt"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

// Init 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password, // 密码
		DB:       cfg.DB,       // 数据库
		PoolSize: cfg.PoolSize, // 连接池大小
	})
	_, err = rdb.Ping().Result()
	return
}
func Close() {
	_ = rdb.Close()
}
