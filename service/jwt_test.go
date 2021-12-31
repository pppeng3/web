package service

import (
	"testing"
	"web/db/model"
)

func TestCreateAndParseToken(t *testing.T) {
	user := model.User{
		UserId:            "3213gsq",
		Password:          "ewqewqasda",
		Sex:               1,
		Nickname:          "ropz",
		Email:             "user@example.com",
		Phone:             "12312312",
		Wechat:            "wyeg12",
		AvatarUrl:         "http://www.example.com",
		PersonalSignature: "jj",
	}
	token, exp, err := CreateAccessToken(user)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token, exp)

	j, err := ParseToken(token, false)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(j)
}
