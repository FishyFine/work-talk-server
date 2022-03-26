package main

import (
	"2225322_WorkTalk/work-talk-server/login/userInfo"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func init() {
	logFile, err := os.OpenFile("./login.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0664)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Ldate)
}
func main() {
	userInfo.Init()
	gin.SetMode(gin.ReleaseMode)
	route()
}
