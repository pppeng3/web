package user_center

import (
	"context"
	user_center "web/proto/user-center"
)

func Register(ctx context.Context, req *user_center.RegisterRequest) (resp *user_center.RegisterResponse, err error) {
	switch req.Type {
	case 0:
		registerWithPhone(ctx, req.User.Account, req.User.Password, req.VerifyCode)
	case 1:
		registerWithEmail(ctx, req.User.Account, req.User.Password, req.VerifyCode)
	case 2:
		registerWithWechat(ctx, req.User.Account, req.User.Password)
	}

	return nil, nil
}

func Login(ctx context.Context, req *user_center.LoginRequest) (resp *user_center.LoginResponse, err error) {

	return nil, nil
}

func Follow(ctx context.Context, req *user_center.FollowRequest) (resp *user_center.FollowResponse, err error) {
	return nil, nil
}

func CheckToken(ctx context.Context, req *user_center.CheckTokenRequest) (resp *user_center.CheckTokenResponse, err error) {
	return nil, nil
}

func RefreshToken(ctx context.Context, req *user_center.RefreshTokenRequest) (resp *user_center.RefreshTokenResponse, err error) {
	return nil, nil
}

func ModifyUserInfo(ctx context.Context, req *user_center.ModifyUserInfoRequest) (resp *user_center.ModifyUserInfoResponse, err error) {
	return nil, nil
}

func ModifyPassword(ctx context.Context, req *user_center.ModifyPasswordRequest) (resp *user_center.ModifyPasswordResponse, err error) {
	return nil, nil
}

func GetUserInfo(ctx context.Context, req *user_center.GetUserInfoRequest) (resp *user_center.GetUserInfoResponse, err error) {
	return nil, nil
}

func CreateGroup(ctx context.Context, req *user_center.CreateGroupRequest) (resp *user_center.CreateGroupResponse, err error) {
	return nil, nil
}

func GetGroup(ctx context.Context, req *user_center.GetGroupRequest) (resp *user_center.GetGroupResponse, err error) {
	return nil, nil
}

func GetGroupMembers(ctx context.Context, req *user_center.GetGroupMembersRequest) (resp *user_center.GetGroupMembersResponse, err error) {
	return nil, nil
}

func JoinGroup(ctx context.Context, req *user_center.JoinGroupRequest) (resp *user_center.JoinGroupResponse, err error) {
	return nil, nil
}

func ExitGroup(ctx context.Context, req *user_center.ExitGroupRequest) (resp *user_center.ExitGroupResponse, err error) {
	return nil, nil
}

func registerWithPhone(ctx context.Context, phone, password, verify string) error {
	// 短信验证
	return nil
}

func registerWithEmail(ctx context.Context, email, password, verify string) error {
	// 邮箱验证
	return nil
}

func registerWithWechat(ctx context.Context, wechat, password string) error {
	// 三方登陆OR微信号登陆
	return nil
}

func generateUserId() uint32 {
	return 1
}
