package main

import (
	"2225322_WorkTalk/register/handleFunc"
	"github.com/gin-gonic/gin"
)

func route(){
	route := gin.Default()
	route.GET("/",handleFunc.Register)
	route.Run(":8081")
}
