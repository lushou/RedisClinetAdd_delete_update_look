package openredis

import (
	"RedisClinetAdd_delete_update_look/determine"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// RedisDelete 进行删除节点，有一个前提就是不允许删除主节点
func RedisDelete(c *gin.Context) {
	redismstart := c.DefaultPostForm("redisMstart", "") // 集群地址
	redisDelete := c.DefaultPostForm("redisDelete", "") // 要删除的节点
	password := c.DefaultPostForm("passWord", "")
	log.Info(redismstart, redisDelete, password)
	if redismstart == "" && redisDelete == "" && password == "" {
		log.Error("传入的数据为空")
		c.JSON(200, gin.H{"code": 404, "msg": "请进行检查传入的数据，传入的数据出现空的现象", "time": time.Now().Format("20060102_15:04:05")})
		return
	}
	// 进行连接redis
	const db = 0
	rdb := determine.Redisplayed(redismstart, password, db)
	err := determine.RedisPing(rdb)
	if err != nil {
		log.Error(err.Error())
		c.JSON(200, gin.H{"code": 404, "msg": err.Error(), "time": time.Now().Format("20060102_15:04:05")})
		return
	}

	// 然后进行获取cluster nodes 的数据，
	RedisClineNnode, err := determine.Redispatching(rdb)
	if err != nil {
		Rediserr := fmt.Sprintf("当前%s不是clinet 集群模式", redismstart)
		log.Error(Rediserr)
		c.JSON(200, gin.H{"code": 404, "msg": Rediserr, "time": time.Now().Format("20060102_15:04:05")})
		return
	}
	// fmt.Println(RedisClineNnode)
	// 代表上面是集群模式
	err = determine.RedisDeleteNode(redismstart, password, redisDelete, RedisClineNnode)
	if err != nil {
		c.JSON(200, gin.H{"code": 404, "mag": err.Error(), "time": time.Now().Format("20060102_15:04:05")})
		return
	}
	log.Info("节点：", redisDelete, "删除成功")
	c.JSON(200, gin.H{"code": 200, "mag": "节点删除成功", "time": time.Now().Format("20060102_15:04:05")})
}
