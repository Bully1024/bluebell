package redis

import (
	"GoWebCode/bluebell/settings"
	"fmt"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

// Init 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password, // 密码
		DB:       cfg.DB,       // 数据库
		PoolSize: cfg.PoolSize, // 连接池大小
		//MinIdleConns: cfg.MinIdleConns,
	})
	_, err = client.Ping().Result()
	if err != nil {

	}
	return
}
func Close() {
	_ = client.Close()
}
