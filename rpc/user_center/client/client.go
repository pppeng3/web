package main

import (
	"context"
	"fmt"
	"web/log"
	user_center "web/proto/user-center"

	"google.golang.org/grpc"
)

const PORT = "9001"

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
	resp, err := client.Register(context.Background(), &user_center.RegisterRequest{User: &user_center.User{Account: "549822881@qq.com", Password: "ropzzzz"}, Type: 1})
	if err != nil {
		log.Error(`error:%v`, err)
	}
	fmt.Println(resp)
	// 登陆
	resp_, err := client.Login(context.Background(), &user_center.LoginRequest{Account: "549822881@qq.com", Password: "ropzzzz", Type: 1})
	if err != nil {
		log.Error(`error:%v`, err)
	}
	fmt.Println(resp_)
}
