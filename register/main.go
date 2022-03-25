package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func init() {
	logFile,err:= os.OpenFile("./register.log",os.O_CREATE|os.O_APPEND|os.O_RDWR,0664)
	if err!=nil{
		panic(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile|log.Ldate)
}
func main(){
	gin.SetMode(gin.ReleaseMode)
	route()
}
