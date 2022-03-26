package handleFunc

import (
	"2225322_WorkTalk/work-talk-server/login/userInfo"
	"github.com/gin-gonic/gin"
	"log"
)

func Login(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	user := new(userInfo.UserInfo)
	user.Username = username
	user.Password = password

	ret, err := user.CheckUserInfo()
	if err != nil {
		ctx.JSON(500, gin.H{"result": "服务器发生了异常，请联系管理员qq：1923325014"})
		return
	}
	if !ret {
		ctx.JSON(400, gin.H{"result": "错误的帐号或密码"})
		return
	}
	token, err := userInfo.BuildJWT(user.Username, user.Password)
	if err != nil {
		ctx.JSON(500, gin.H{"result": "服务器发生了异常，请联系管理员qq：1923325014"})
		return
	}

	//success
	go func() {
		if err := user.Publish(); err != nil {
			log.Println(err)
		}
	}()
	go func() {
		if err := user.SaveToken(token); err != nil {
			log.Println(err)
		}
	}()
	ctx.JSON(200, gin.H{"result": "success",
		"token": token,
	})
}
