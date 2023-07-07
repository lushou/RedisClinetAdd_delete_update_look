package openredis

import (
	"RedisClinetAdd_delete_update_look/determine"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func RedisAdd(c *gin.Context) {
	// 获取传入的参数
	redismstart := c.DefaultPostForm("redisMstart", "")
	redisadd := c.DefaultPostForm("redisAdd", "")
	password := c.DefaultPostForm("passWord", "")
	// 进行判断数据是否是空值，如果是空返回报错
	if redismstart == "" && redisadd == "" && password == "" {
		log.Error("传入的数据为空")
		c.JSON(404, gin.H{"code": 404, "msg": "请进行检查传入的数据，传入的数据出现空的现象"})
		return
	}
	const db = 0
	// 进行检查是否可以进行连接
	rdb := determine.Redisplayed(redismstart, password, db)
	err := determine.RedisPing(rdb)
	if err != nil {
		Rediserr := fmt.Sprintf("当前集群%s无法进行连接请进行检查", redismstart)
		log.Error(Rediserr)
		c.JSON(404, gin.H{"code": 404, "msg": Rediserr})
		return
	}
	// 进行测试要添加的节点是否可以进行连接：

	rdbCong := determine.Redisplayed(redisadd, password, db)
	err = determine.RedisPing(rdbCong)
	if err != nil {
		Rediserr := fmt.Sprintf("当前集群%s无法进行连接请进行检查", redisadd)
		log.Error(Rediserr)
		c.JSON(404, gin.H{"code": 404, "msg": Rediserr})
		return
	}
	// 进行检测要进行添加的节点是否有槽，通过有就不让进行添加
	err = determine.Trough(rdbCong, redismstart, password, redisadd)
	if err != nil {
		log.Error(err)
		c.JSON(404, gin.H{"code": 404, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "添加完成", "time": time.Now().Format("20060102_15:04:05")})
	log.Info("添加完成")
}
