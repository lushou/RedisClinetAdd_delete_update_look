package determine

import (
	"context"
	"strings"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

// redisConnectAndCheckRole 这个进行查看这个进行判断redis主从
func RedisConnectAndCheckRole(redisslave, RedisClineNnode string) bool {
	redismstartList := strings.Split(RedisClineNnode, "\n") // 这个按照回车进行切割
	log.Debug(redismstartList)                              // 把刚刚查询的结果添加道日志中
	for _, v := range redismstartList {
		if strings.Contains(v, "myself,slave") || strings.Contains(v, "slave") {
			if strings.Contains(v, redisslave) { // 代表可以进行匹配到
				return true
			}

		}
	}
	return false

}

// RedisSlave 进行变更成主节点
func RedisMstart(rdb *redis.Client) error {
	if err := rdb.ClusterFailover(context.Background()).Err(); err != nil {
		return err
	}
	return nil
}
