package main

import (
	"2225322_WorkTalk/work-talk-server/login/handleFunc"
	"github.com/gin-gonic/gin"
)

func route() {
	route := gin.Default()
	route.GET("/", handleFunc.Login)
	route.Run(":8082")
}
