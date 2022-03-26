package userInfo

import (
	"github.com/spf13/viper"
	"log"
	"strconv"
)

type JWTConfig struct {
	Issuer     string //token发行人的名称
	ExpireTime int    //有效期，单位时间
	Secret     []byte //采用的密钥
}

const configName string = "jwt"

func (j *JWTConfig) SetDefaultConfig() {
	viper.SetDefault(configName+"."+"issur", "keli")
	viper.SetDefault(configName+"."+"expiretime", "48")
	viper.SetDefault(configName+"."+"secret", "chengfengpolang")
}

func (j *JWTConfig) GetConfig() {
	j.SetDefaultConfig()

	j.Issuer = viper.Get(configName + "." + "issur").(string)
	expireTime, err := strconv.Atoi(viper.Get(configName + "." + "expiretime").(string))
	if err != nil {
		log.Panicln(err)
		return
	}
	j.ExpireTime = expireTime
	j.Secret = []byte(viper.Get(configName + "." + "secret").(string))
}
