package main

import (
	"RedisClinetAdd_delete_update_look/route"
	"fmt"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

func init() {
	path := "./logs/Rcdupl"
	writer, _ := rotatelogs.New(
		path+".%Y%m%d.log",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(5*24*time.Hour),     // 5天的最长保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 每天分割一次日志
	)
	log.SetOutput(writer)
	log.SetFormatter(&log.JSONFormatter{})
}
func main() {
	route := route.Route()

	route.Run(":80")
	fmt.Println("sadf")
}
