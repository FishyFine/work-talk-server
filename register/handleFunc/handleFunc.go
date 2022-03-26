package handleFunc

import (
	"2225322_WorkTalk/register/userInfo"
	"github.com/gin-gonic/gin"
	"log"
)

func Register(ctx *gin.Context){
	username := ctx.Query("username")
	password := ctx.Query("password")
	user := new(userInfo.UserInfo)
	user.Username = username
	user.Password = password

	//Check whether the information can be registered
	ret,err := user.Check()
	if err!=nil{
		log.Println(err)
		ctx.JSON(500,gin.H{"result":"服务器发生了异常，请联系管理员qq：1923325014"})
		return
	}
	if !ret{
		ctx.JSON(400,gin.H{"result":"用户名已存在"})
		return
	}

	//success
	user.Register()
	ctx.JSON(200,gin.H{"result":"success"})
}
