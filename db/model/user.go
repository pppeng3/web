package model

type User struct {
	UserId            uint32  ``
	Password          string  `gorm:"column:password;type:varchar(32);not null"`
	Sex               uint32  `gorm:"column:sex;type:tinyint(1)"`
	Nickname          string  `gorm:"column:nickname;type:varchar(32)"`
	Email             string  `gorm:"column:email;type:varchar(32)"`
	Phone             string  `gorm:"column:phone;type:varchar(32)"`
	Wechat            string  `gorm:"column:wechat;type:varchar(32)"`
	AvatarUrl         string  `gorm:"column:avatar_url;type:varchar(32)"`
	Follow            []int32 `gorm:"column:follow;type:text"`
	Fans              []int32 `gorm:"column:fans;type:text"`
	PersonalSignature string
	Auth              []int32 `gorm:"column:auth;type:text"`
}
