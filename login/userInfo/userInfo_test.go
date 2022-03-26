package userInfo

import (
	"fmt"
	"testing"
)

func TestUserInfo_CheckUserInfo(t *testing.T) {
	u := new(UserInfo)
	u.Username = "user1"
	u.Password = "12345"
	fmt.Println(u.CheckUserInfo())
}

func TestBuildJWT(t *testing.T) {
	u := new(UserInfo)
	u.Username = "user1"
	u.Password = "12345"
	Init()
	fmt.Println(BuildJWT(u.Username, u.Password))
}

func TestUserInfo_Publish(t *testing.T) {
	u := new(UserInfo)
	u.Username = "user1"
	u.Password = "123456"
	u.Publish()
}

func TestUserInfo_SaveToken(t *testing.T) {
	u := new(UserInfo)
	u.Username = "user1"
	u.Password = "123456"
	token, err := BuildJWT(u.Username, u.Password)
	if err != nil {
		panic(err)
	}
	u.SaveToken(token)
}
