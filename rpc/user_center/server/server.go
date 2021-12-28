package ucserver

import (
	"context"
	"log"
	"net"

	huc "web/handler/user_center"
	user_center "web/proto/user-center"

	"google.golang.org/grpc"
)

const PORT = "9001"

type UserCenterServer struct {
}

func (s UserCenterServer) Register(ctx context.Context, req *user_center.RegisterRequest) (resp *user_center.RegisterResponse, err error) {
	return huc.Register(ctx, req)
}

func (s UserCenterServer) Login(ctx context.Context, req *user_center.LoginRequest) (resp *user_center.LoginResponse, err error) {
	return huc.Login(ctx, req)
}

func (s UserCenterServer) Follow(ctx context.Context, req *user_center.FollowRequest) (resp *user_center.FollowResponse, err error) {
	return huc.Follow(ctx, req)
}

func (s UserCenterServer) CheckToken(ctx context.Context, req *user_center.CheckTokenRequest) (resp *user_center.CheckTokenResponse, err error) {
	return huc.CheckToken(ctx, req)
}

func (s UserCenterServer) RefreshToken(ctx context.Context, req *user_center.RefreshTokenRequest) (resp *user_center.RefreshTokenResponse, err error) {
	return huc.RefreshToken(ctx, req)
}

func (u UserCenterServer) ModifyUserInfo(ctx context.Context, req *user_center.ModifyUserInfoRequest) (resp *user_center.ModifyUserInfoResponse, err error) {
	return huc.ModifyUserInfo(ctx, req)
}

func (u UserCenterServer) ModifyPassword(ctx context.Context, req *user_center.ModifyPasswordRequest) (resp *user_center.ModifyPasswordResponse, err error) {
	return huc.ModifyPassword(ctx, req)
}

func (s UserCenterServer) GetUserInfo(ctx context.Context, req *user_center.GetUserInfoRequest) (resp *user_center.GetUserInfoResponse, err error) {
	return huc.GetUserInfo(ctx, req)
}

func (s UserCenterServer) CreateGroup(ctx context.Context, req *user_center.CreateGroupRequest) (resp *user_center.CreateGroupResponse, err error) {
	return huc.CreateGroup(ctx, req)
}

func (s UserCenterServer) GetGroup(ctx context.Context, req *user_center.GetGroupRequest) (resp *user_center.GetGroupResponse, err error) {
	return huc.GetGroup(ctx, req)
}

func (s UserCenterServer) GetGroupMembers(ctx context.Context, req *user_center.GetGroupMembersRequest) (resp *user_center.GetGroupMembersResponse, err error) {
	return huc.GetGroupMembers(ctx, req)
}

func (s UserCenterServer) JoinGroup(ctx context.Context, req *user_center.JoinGroupRequest) (resp *user_center.JoinGroupResponse, err error) {
	return huc.JoinGroup(ctx, req)
}

func (s UserCenterServer) ExitGroup(ctx context.Context, req *user_center.ExitGroupRequest) (resp *user_center.ExitGroupResponse, err error) {
	return huc.ExitGroup(ctx, req)
}

func main() {
	server := grpc.NewServer()
	user_center.RegisterUserCenterServer(server, &UserCenterServer{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf(`net listen error: %v`, err)
	}

	server.Serve(lis)
}
