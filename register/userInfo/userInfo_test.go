package userInfo

import (
	"testing"
)

func TestUserInfo_Register(t *testing.T) {
	info := new(UserInfo)
	info.Username = "testUsername"
	info.Password = "testPassword"
	//conf := mysqlFunc.NewMysqlConf()
	//helper,err := mysqlFunc.NewMysqlHelper(conf)
	//if err!=nil{
	//	log.Println(err)
	//}
	//helper.DB.AutoMigrate(&UserInfo{})
	info.Register()
	for{}
}
