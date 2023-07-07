package determine

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

// RedisPing 进行验证是否可以连接
func RedisPing(rdb *redis.Client) error {
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return errors.New("无法进行连接redis请进行检查是否有问题")
	}
	return nil
}

// Redisplayed 连接redis 的函数
func Redisplayed(redisIpPost, password string, DB int) (rdb *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisIpPost,
		Password: password, // no password set
		DB:       DB,       // use default DB
	})
	return
}
