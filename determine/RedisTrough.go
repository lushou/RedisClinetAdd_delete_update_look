package determine

import (
	"context"
	"errors"
	"fmt"
	"os/exec"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

// 进行判断是否有槽通过有不进行添加如果没在进行添加

func Trough(rdbCong *redis.Client, redisStat, password, redisAdd string) error {
	// 添加节点是时候进行验证之前是否有槽
	slots, _ := rdbCong.ClusterSlots(context.Background()).Result()
	fmt.Println(slots)
	if len(slots) > 1 {
		Rediscao := fmt.Sprintf("%s节点已经有槽了无需在添加已经村早槽", redisAdd)
		log.Error(Rediscao)
		err := errors.New(Rediscao)
		return err
	}
	fmt.Println("/root/redis-cli", "-a", password, "--cluster", "add-node", redisAdd, redisStat)
	//output, _ := exec.Command("./redis-service/bin/redis-cli", "-a", password, "--cluster", "add-node", redisAdd, redisStat).Output()
	output, err := exec.Command("/root/redis-cli", "-a", password, "--cluster", "add-node", redisAdd, redisStat).Output()
	if err != nil {
		return err
	}
	log.Info("添加成功", string(output))
	return nil
}

func Redisdeletid(redismstart, redisid, password string) error {
	fmt.Println("/root/redis-cli", "--cluster", "del-node", redismstart, redisid, "-a", password)
	output, _ := exec.Command("/root/redis-cli", "--cluster", "del-node", redismstart, redisid, "-a", password).Output()
	log.Debug(string(output))
	return nil
}
