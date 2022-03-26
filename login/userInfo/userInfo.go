package userInfo

import (
	"2225322_WorkTalk/work-talk-server/login/mysqlFunc"
	"2225322_WorkTalk/work-talk-server/login/nsqFunc"
	"2225322_WorkTalk/work-talk-server/login/redisFunc"
	"context"
	"encoding/json"
	"log"
	"time"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserInfo) CheckUserInfo() (bool, error) {
	conf := mysqlFunc.NewMysqlConf()
	helper, err := mysqlFunc.NewMysqlHelper(conf)
	if err != nil {
		return false, err
	}
	ret := new(UserInfo)
	helper.DB.Where("username=?", u.Username).First(ret)
	if ret.Password == u.Password {
		return true, nil
	}
	return false, nil
}

func (u *UserInfo) Publish() error {
	data, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
	}
	conf := nsqFunc.NewNSQConf()
	helper, err := nsqFunc.NewNSQHelper(conf)
	if err != nil {
		return err
	}
	helper.Producer.Publish(conf.Topic, data)
	return nil
}
func (u *UserInfo) SaveToken(token string) error {
	redisConf := redisFunc.NewRedisConf()
	redisHelper, err := redisFunc.NewRedisHelper(redisConf)
	if err != nil {
		return err
	}
	//将token的保存时间设置的比其过期时间更长一些
	err = redisHelper.RedisClient.Set(context.Background(), u.Username, token, time.Hour*time.Duration(jwtConfig.ExpireTime+1)).Err()
	if err != nil {
		return err
	}
	return nil
}
