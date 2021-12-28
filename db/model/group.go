package model

type Group struct {
	GroupId     uint32
	CreatedAt   int64
	UpdatedAt   int64
	OwnerId     uint32
	GroupName   string
	GroupNotice string
}
