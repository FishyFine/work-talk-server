package redisFunc

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
)

type RedisHelper struct {
	RedisClient *redis.Client
}

func NewRedisHelper(conf *RedisConf) (*RedisHelper, error) {
	helper := new(RedisHelper)
	helper.RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + conf.Port,
		Password: "",
		DB:       conf.DB,
	})
	pong, err := helper.RedisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	if pong == "" {
		return nil, errors.New("there are not error at Ping(),but the pong is empty")
	}
	return helper, nil
}
