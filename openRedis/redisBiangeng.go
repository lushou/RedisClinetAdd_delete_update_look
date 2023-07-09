package openredis

import (
	"RedisClinetAdd_delete_update_look/determine"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Redisbiangeng 进行手动复制redis节点
func Redisbiangeng(c *gin.Context) {
	redisslave := c.DefaultPostForm("redisslave", "")   // 集群地址
	redismstart := c.DefaultPostForm("redismstart", "") // 集群地址
	password := c.DefaultPostForm("passWord", "")
	log.Info(redisslave, "要进行复制", redismstart, "密码是：", password)
	if redisslave == "" || password == "" || redismstart == "" {
		log.Error("传入的数据为空")
		c.JSON(200, gin.H{"code": 404, "msg": "请进行检查传入的数据，传入的数据出现空的现象", "time": time.Now().Format("20060102_15:04:05")})
		return
	}
	// 进行连接redis
	const db = 0
	rdbslave := determine.Redisplayed(redisslave, password, db) // 从节点的连接
	err := determine.RedisPing(rdbslave)
	if err != nil {
		reidslogs := fmt.Sprintf("redis节点：%s无法进行连接", redisslave)
		log.Error(reidslogs)
		c.JSON(200, gin.H{"code": 404, "msg": reidslogs, "time": time.Now().Format("20060102_15:04:05")})
		return
	}
	rdbmstart := determine.Redisplayed(redismstart, password, db) // 从节点的连接
	err = determine.RedisPing(rdbmstart)
	if err != nil {
		reidslogs := fmt.Sprintf("redis节点：%s无法进行连接", redismstart)
		log.Error(reidslogs)
		c.JSON(200, gin.H{"code": 404, "msg": reidslogs, "time": time.Now().Format("20060102_15:04:05")})
		return
	}
	// 上面的可以进行连接成功，下面进行
	RedisClineNnode, err := determine.Redispatching(rdbslave)
	if err != nil {
		Rediserr := fmt.Sprintf("当前%s不是clinet 集群模式", redisslave)
		log.Error(Rediserr)
		c.JSON(200, gin.H{"code": 404, "msg": Rediserr, "time": time.Now().Format("20060102_15:04:05")})
		return
	}
	if !strings.Contains(RedisClineNnode, redismstart) {
		Rediserr := fmt.Sprintf("集群：%s中不存在节点：%s", redisslave, redismstart)
		log.Error(Rediserr)
		c.JSON(200, gin.H{"code": 404, "msg": Rediserr, "time": time.Now().Format("20060102_15:04:05")})
		return
	}
	// 下面表示存在
	err = determine.Redisbiangeng(RedisClineNnode, redismstart, redisslave, rdbslave)
	if err != nil {
		log.Error(err)
		c.JSON(200, gin.H{"code": 404, "msg": err.Error(), "time": time.Now().Format("20060102_15:04:05")})
		return
	}

	log.Info("切换成功")
	c.JSON(200, gin.H{"code": 200, "msg": "节点切换成功", "time": time.Now().Format("20060102_15:04:05")})
}
