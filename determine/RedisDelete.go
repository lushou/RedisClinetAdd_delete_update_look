package determine

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

// Redispatching ,进行返回redis clinet 数据
func Redispatching(rdb *redis.Client) (RedisClineNnode string, err error) {
	cmd := rdb.ClusterNodes(context.Background())
	result, err := cmd.Result()
	if err != nil {
		return "", err
	}
	return result, nil

}

// RedisDeleteNode 删除相关的
func RedisDeleteNode(redismstart, password, redisDelete, RedisClineNnode string) error {
	//首先首先进行判断集群中是否存在要进行删除的节点
	if strings.Count(RedisClineNnode, redisDelete) == 0 {
		redisNowErr := fmt.Sprintf("集群%s中不存在节点%s", redismstart, redisDelete)
		log.Error(redisNowErr)
		return errors.New(redisNowErr)
	}
	// 道这一步也也就代表代码中存在这个参数
	redismstartList := strings.Split(RedisClineNnode, "\n") // 这个按照回车进行切割
	log.Debug(redismstartList)                              // 把刚刚查询的结果添加道日志中
	// 进行验证是否存在主节点
	for _, redisIpstr := range redismstartList {
		if strings.Contains(redisIpstr, "myself,master") || strings.Contains(redisIpstr, "master") { // 进行验证主节点中是否包含该配置文件
			// fmt.Println(redisIpstr)
			if strings.Contains(redisIpstr, redisDelete) {
				fmt.Println("包含了")
				redisNowErr := fmt.Sprintf("集群%s中%s节点是主节点不支持删除,请进行修改成从节点，在进行删除，", redismstart, redisDelete)
				log.Error(redisNowErr)
				return errors.New(redisNowErr)
			}
		} else { // 这部分相当于都是从节点
			// fmt.Println(redisIpstr)
			if strings.Contains(redisIpstr, redisDelete) { // 代表从节点中有这个要删除的节点
				// 这边进行获取要删除的节点中的redisid
				log.Debug("要删除的节点", redisIpstr, "已经可以进行找到")
				fmt.Println(redisIpstr) // 2feaaab9489295d7e66d26c35d235a4c34fccb42 192.168.0.163:7307@17307 master - 0 1688520933000 0 c
				// 进行获取nodeid
				redisNnodid := strings.Split(redisIpstr, " ")
				log.Debug(redisNnodid[0])
				// 进行删除该节点
				err := Redisdeletid(redismstart, redisNnodid[0], password)
				fmt.Println(err,"___")
				if err != nil {
					return err
				}
				return nil
			}
		}
	}

	// 要删除的节点是否是主节点，如果是主节点返回不能删除

	return nil
}
