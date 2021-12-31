package model

type User struct {
	UserId            string  `gorm:"primaryKey;column:user_id;type:varchar(32);not null" json:"user_id"`
	Password          string  `gorm:"column:password;type:varchar(32);not null" json:"password"`
	Sex               uint32  `gorm:"column:sex;type:tinyint(1)" json:"sex"`
	Nickname          string  `gorm:"column:nickname;type:varchar(32)" json:"nickname"`
	Email             string  `gorm:"index;column:email;type:varchar(32)" json:"email"`
	Phone             string  `gorm:"index;column:phone;type:varchar(32)" json:"phone"`
	Wechat            string  `gorm:"index;column:wechat;type:varchar(32)" json:"wechat"`
	AvatarUrl         string  `gorm:"column:avatar_url;type:varchar(32)" json:"avatar_url"`
	Follow            []int32 `gorm:"column:follow;type:text" json:"follow"`
	Fans              []int32 `gorm:"column:fans;type:text" json:"fans"`
	PersonalSignature string  `gorm:"column:personal_signature;type:varchar(32)" json:"personal_signature"`
	Auth              []int32 `gorm:"column:auth;type:text"`
	LoginAt           int64   `gorm:"column:login_at;type:int(10) unsigned;not null" json:"login_at"`
}
