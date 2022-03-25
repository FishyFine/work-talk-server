package userInfo

import (
	"2225322_WorkTalk/register/mysqlFunc"
	"2225322_WorkTalk/register/nsqFunc"
	"encoding/json"
	"log"
)

type UserInfo struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserInfo)Check()(bool,error){
	ret,err := u.checkExist()
	if err!=nil{
		return false,err
	}
	return ret,nil
}

// CheckExist if the username is existed,it returns the false
func (u *UserInfo)checkExist()(bool,error){
	conf := mysqlFunc.NewMysqlConf()
	helper,err := mysqlFunc.NewMysqlHelper(conf)
	if err !=nil{
		return false,err
	}
	ret := new(UserInfo)
	helper.DB.Where("username=?",u.Username).First(ret)
	if ret.Username==""{
		return true,nil
	}else{
		return false,nil
	}
}

func (u *UserInfo)Register(){
	go u.save()
	go u.publish()
}

func (u *UserInfo)save(){
	conf := mysqlFunc.NewMysqlConf()
	helper,err := mysqlFunc.NewMysqlHelper(conf)
	if err!=nil{
		log.Println(err)
	}
	helper.DB.Create(u)
}

func (u *UserInfo)publish(){
	data,err := json.Marshal(u)
	if err!=nil{
		log.Println(err)
	}
	conf := nsqFunc.NewNSQConf()
	helper,err:=nsqFunc.NewNSQHelper(conf)
	if err!=nil{
		log.Println(err)
	}
	helper.Producer.Publish(conf.Topic,data)
}