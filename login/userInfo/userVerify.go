package userInfo

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserInfo
	jwt.StandardClaims
}

//CheckJWT 解码Token，获取用户信息;当error不为空时，说明token有问题
func CheckJWT(token string) (string, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtConfig.Secret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims.checkClaims()
		}
	}
	return "", err
}

func (c Claims) checkClaims() (string, error) {
	if time.Now().Unix() > c.ExpiresAt {
		return "", errors.New("token已经过期")
	} else {
		return c.Username, nil
	}
}

func BuildJWT(userName string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(jwtConfig.ExpireTime) * time.Hour)
	var claims = Claims{
		UserInfo: UserInfo{
			Username: userName,
			Password: password,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    jwtConfig.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtConfig.Secret)
}
