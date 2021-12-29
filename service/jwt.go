package service

import (
	"time"
	"web/db/model"

	"github.com/dgrijalva/jwt-go"
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
	exp = claims.StandardClaims.ExpiresAt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(key))
	if err != nil {
		return
	}
	return
}

func CreateAccessToken() {}

func CreateRefreshToken() {}

func ParseToken(token string, isRefresh bool) (jwtCustomClaims, error) {
	// var key string
	return jwtCustomClaims{}, nil
}
