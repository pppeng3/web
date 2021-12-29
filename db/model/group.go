package model

type Group struct {
	GroupId     uint32 `gorm:"primaryKey;column:group_id;type:int(10) unsigned;not null" json:"group_id"`
	CreatedAt   int64  `gorm:"column:created_at;type:int(10) unsigned;not null" json:"created_at"`
	UpdatedAt   int64  `gorm:"column:updated_at;type:int(10) unsigned;not null" json:"updated_at"`
	OwnerId     uint32 `gorm:"index;column:owner_id;type:int(10) unsigned;not null" json:"owner_id"`
	GroupName   string `gorm:"column:group_name;type:varchar(32)" json:"group_name"`
	GroupNotice string `gorm:"column:group_notice;type:varchar(32)" json:"group_notice"`
}
