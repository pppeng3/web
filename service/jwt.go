package service

import (
	"time"
	"web/db/model"
	"web/log"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	accessMaxAge  = 100
	refreshMaxAge = 60 * 60 * 24 * 7
)

var (
	accessSecret  string
	refreshSecret string
)

type jwtCustomClaims struct {
	jwt.StandardClaims

	UserId            string `gorm:"primaryKey;column:user_id;type:varchar(32);not null" json:"user_id"`
	Sex               uint32 `gorm:"column:sex;type:tinyint(1)" json:"sex"`
	Nickname          string `gorm:"column:nickname;type:varchar(32)" json:"nickname"`
	Email             string `gorm:"index;column:email;type:varchar(32)" json:"email"`
	Phone             string `gorm:"index;column:phone;type:varchar(32)" json:"phone"`
	Wechat            string `gorm:"index;column:wechat;type:varchar(32)" json:"wechat"`
	AvatarUrl         string `gorm:"column:avatar_url;type:varchar(32)" json:"avatar_url"`
	PersonalSignature string `gorm:"column:personal_signature;type:varchar(32)" json:"personal_signature"`
}

func init() {
	accessSecret, refreshSecret = getJWTConf()
}

func getJWTConf() (accessKey, refreshKey string) {
	viper.SetConfigName("jwt_conf")
	viper.AddConfigPath("/Users/pengdarong/Desktop/Personal/web/conf")
	if err := viper.ReadInConfig(); err != nil {
		log.Error(`Fail To Load AccessKey And RefreshKey: %v`, err)
	}
	return viper.GetString("accessKey"), viper.GetString("refreshKey")
}

func createToken(user model.User, timeLength int, key string) (tokenString string, exp int64, err error) {
	claims := jwtCustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(timeLength) * time.Second).Unix(),
			Issuer:    "user_center",
		},
		UserId:            user.UserId,
		Sex:               user.Sex,
		Nickname:          user.Nickname,
		Email:             user.Email,
		Phone:             user.Phone,
		Wechat:            user.Wechat,
		AvatarUrl:         user.AvatarUrl,
		PersonalSignature: user.PersonalSignature,
	}
	exp = claims.ExpiresAt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(key))
	if err != nil {
		return
	}
	return
}

func CreateAccessToken(user model.User) (tokenString string, exp int64, err error) {
	return createToken(user, accessMaxAge, accessSecret)
}

func CreateRefreshToken(user model.User) (tokenString string, exp int64, err error) {
	return createToken(user, refreshMaxAge, refreshSecret)
}

func ParseToken(tokenString string, isRefresh bool) (jwtCustomClaims, error) {
	var key string
	if isRefresh {
		key = refreshSecret
	} else {
		key = accessSecret
	}
	token, err := jwt.ParseWithClaims(tokenString, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf(`Unexpected signing method: %v`, token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		log.Warn(`Parse Token Error: %v`, tokenString)
		return jwtCustomClaims{}, errors.WithStack(err)
	}
	if claims, ok := token.Claims.(*jwtCustomClaims); ok && token.Valid {
		return *claims, nil
	} else {
		log.Info(`%+v`, token.Claims)
		log.Warn(`%+v`, token.Valid)
		return jwtCustomClaims{}, errors.WithStack(errors.New("invalid claims or token"))
	}
}
