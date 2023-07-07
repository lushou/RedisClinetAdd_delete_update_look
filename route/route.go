package route

import (
	openredis "RedisClinetAdd_delete_update_look/openRedis"

	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	route := gin.Default()
	// 进行添加的接口
	route.POST("/redisadd/", openredis.RedisAdd)
	// 进行删除的接口
	route.POST("/redisdelete/", openredis.RedisDelete)
	// 进行把从服务器变更程主服务器的接口
	route.POST("/redisfailover/",openredis.Redisfailover)
	// 进行主从赋值的接口
	route.POST("/redisbiangeng/",openredis.Redisbiangeng)
	return route
}
