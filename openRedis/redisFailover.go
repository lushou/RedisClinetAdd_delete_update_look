package openredis

import (
	"RedisClinetAdd_delete_update_look/determine"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// redisfailover 进行执行failover操作吧redis从节点变更为主节点
func Redisfailover(c *gin.Context) {
	redisslave := c.DefaultPostForm("redisSlave", "") // 集群地址
	password := c.DefaultPostForm("passWord", "")
	log.Info(redisslave, password)
	if redisslave == "" || password == "" {
		log.Error("传入的数据为空")
		c.JSON(404, gin.H{"code": 404, "msg": "请进行检查传入的数据，传入的数据出现空的现象"})
		return
	}
	// 进行连接redis
	const db = 0
	rdb := determine.Redisplayed(redisslave, password, db)
	err := determine.RedisPing(rdb)
	if err != nil {
		log.Error(err.Error())
		c.JSON(404, gin.H{"code": 404, "msg": err.Error()})
		return
	}
	RedisClineNnode, err := determine.Redispatching(rdb)
	if err != nil {
		Rediserr := fmt.Sprintf("当前%s不是clinet 集群模式", redisslave)
		log.Error(Rediserr)
		c.JSON(200, gin.H{"code": 404, "msg": Rediserr})
		return
	}
	// 代表可以连接成功，进行判断连接的节点是否是主节点
	if !determine.RedisConnectAndCheckRole(redisslave, RedisClineNnode) {
		log.Error("节点", redisslave, "是主节点无需在进行操作")
		c.JSON(200, gin.H{"code": 200, "msg": "该节点是主节点，无需进行变更操作", "time": time.Now().Format("20060102_15:04:05")})
		return
	}
	// 进行变更曾从
	if err = determine.RedisMstart(rdb); err != nil {
		log.Error("节点", redisslave, "切换失败", err.Error())
		c.JSON(200, gin.H{"code": 200, "msg": err.Error(), "time": time.Now().Format("20060102_15:04:05")})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "切换成功", "time": time.Now().Format("20060102_15:04:05")})

}
