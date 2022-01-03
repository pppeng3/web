package main

import (
	"context"
	"fmt"
	"web/log"
	user_center "web/proto/user-center"

	"google.golang.org/grpc"
)

const PORT = "9001"
var uCtx = context.Background()

func main() {
	log.Init("/Users/pengdarong/Desktop/Personal/web/logs", "uc", "[uc] 🎄 ", "info")
	log.Info(`Client Start`)
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatal(`Fail to dial grpc.Client: %v`, err)
	}
	defer conn.Close()

	client := user_center.NewUserCenterClient(conn)

	// 注册
	resp, err := client.Register(uCtx, &user_center.RegisterRequest{User: &user_center.User{Account: "549822881@qq.com", Password: "ropzzzz"}, Type: 1})
	if err != nil {
		log.Error(`error:%v`, err)
	}
	fmt.Println(resp)
	// 登陆
	resp2, err := client.Login(uCtx, &user_center.LoginRequest{Account: "549822881@qq.com", Password: "ropzzzz", Type: 1})
	if err != nil {
		log.Error(`error:%v`, err)
	}
	fmt.Println(resp2)
	// checkToken
	resp3, err := client.CheckToken(uCtx, &user_center.CheckTokenRequest{Token: resp2.RefreshToken, IsRefresh: true})
	if err != nil {
		log.Error(`error:%v`, err)
	}
	fmt.Println(resp3)
	// refreshToken
	resp4, err := client.RefreshToken(uCtx, &user_center.RefreshTokenRequest{RefreshToken: resp2.RefreshToken})
	if err != nil {
		log.Error(`error:%v`, err)
	}
	fmt.Println(resp4)
	// ModifyUser
	resp5, err := client.ModifyUserInfo(uCtx, &user_center.ModifyUserInfoRequest{
		UserInfo: &user_center.UserInfo_{
			Sex: 1,
			Nickname: "pp",
			Email: "549822881@qq.com",
			Phone: "13377712345",
			Wechat: "qweyuwqy",
			AvatarUrl: "",
			PersonalSignature: "hu",
		},
	})
	if err != nil {
		log.Error(`error:%v`, err)
	}
	fmt.Println(resp5)
}
