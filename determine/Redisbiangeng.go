package determine

import (
	"context"
	"fmt"
	"strings"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

// // Redisbiangeng 进行做redis的主从复制
// func Redisbiangeng(RedisClineNnode, redismstart, redisslave string, rdbslave *redis.Client) error {
// 	// 进行判断RedisClineNnode 中的redismster是不是主和redisslav是不是主
// 	var RestartID *string
// 	redismstartList := strings.Split(RedisClineNnode, "\n") // 这个按照回车进行切割
// 	log.Debug(redismstartList)                              // 把刚刚查询的结果添加道日志中
// 	for _, redisIpstr := range redismstartList {
// 		if (strings.Contains(redisIpstr, "myself,master") || strings.Contains(redisIpstr, "master")) && strings.Contains(redisIpstr, "redismstart") {
// 			// 进行获取主节点的id
// 			redisNnodid := strings.Split(redisIpstr, " ")
// 			RestartID = &redisNnodid[0]
// 		} else {
// 			if strings.Contains(redisIpstr, redisslave) {
// 				// 进行主/从复制
// 				err := rdbslave.ClusterReplicate(context.Background(), *RestartID).Err()
// 				if err != nil {
// 					return err
// 				}
// 				return nil
// 			}
// 		}
// 	}
// }

// Redisbiangeng 用于进行 Redis 的主从复制
func Redisbiangeng(RedisClineNnode, redismstart, redisslave string, rdbsave *redis.Client) error {
	redismstartList := strings.Split(RedisClineNnode, "\n") // 按换行符进行切割
	log.Debug(redismstartList)                              // 将查询结果添加到日志中

	var restartID string
	for _, redisIpstr := range redismstartList {
		if (strings.Contains(redisIpstr, "myself,master") || strings.Contains(redisIpstr, "master")) && strings.Contains(redisIpstr, redismstart) {
			redisNnodid := strings.Split(redisIpstr, " ")[0] // 提取第一个字段
			restartID = redisNnodid
		} else if strings.Contains(redisIpstr, redisslave) {
			err := rdbsave.ClusterReplicate(context.Background(), restartID).Err()
			if err != nil {
				return err
			}
			return nil
		}
		redislogs := fmt.Errorf("节点", redisslave, "是主节点")
		return redislogs
	}
	return nil // 处理条件都不满足的情况，返回错误或采取其他处理方式
}
