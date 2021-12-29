package user_center

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"time"
	"web/db/model"
	"web/db/mysql"
	"web/log"
	user_center "web/proto/user-center"
	"web/utils/helper"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	db = mysql.Instance()
	log.Init("/Users/pengdarong/Desktop/Personal/web/logs", "uc", "[uc] 🎄 ", "info")
}

func Register(ctx context.Context, req *user_center.RegisterRequest) (resp *user_center.RegisterResponse, err error) {
	switch req.Type {
	case 0:
		return registerWithPhone(ctx, req.User.Account, req.User.Password, req.VerifyCode)
	case 1:
		return registerWithEmail(ctx, req.User.Account, req.User.Password, req.VerifyCode)
	case 2:
		return registerWithWechat(ctx, req.User.Account, req.User.Password)
	}

	return nil, nil
}

func Login(ctx context.Context, req *user_center.LoginRequest) (resp *user_center.LoginResponse, err error) {
	switch req.Type {
	case 0:
		return loginWithPhone(ctx, req.Account, req.Password)
	case 1:
		return loginWithEmail(ctx, req.Account, req.Password)
	case 2:
		return loginWithWechat(ctx, req.Account, req.Password)
	}
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

func registerWithPhone(ctx context.Context, phone, password, verify string) (resp *user_center.RegisterResponse, err error) {
	// 检验电话号码格式
	if err := helper.CheckPhone(phone); err != nil {
		return nil, err
	}
	// 验证码

	uid, err := register(ctx, phone, password, 0)
	if err != nil {
		return nil, err
	}
	resp = &user_center.RegisterResponse{
		UserId: uid,
	}
	return resp, nil
}

func registerWithEmail(ctx context.Context, email, password, verify string) (resp *user_center.RegisterResponse, err error) {
	// 检验邮箱格式
	if err := helper.CheckEmail(email); err != nil {
		return nil, err
	}

	// 验证码

	uid, err := register(ctx, email, password, 1)
	if err != nil {
		return nil, err
	}
	resp = &user_center.RegisterResponse{
		UserId: uid,
	}
	return resp, nil
}

func registerWithWechat(ctx context.Context, wechat, password string) (resp *user_center.RegisterResponse, err error) {
	// 三方登陆OR微信号登陆

	return nil, nil
}

func register(ctx context.Context, account, password string, typ int) (uid string, err error) {
	err = db.WithContext(ctx).Transaction(func(tx *gorm.DB) (err error) {
		uid = ksuid.New().String()
		user := model.User{
			UserId:   uid,
			Password: fmt.Sprintf("%x", md5.Sum([]byte(password))),
			LoginAt:  time.Now().Unix(),
		}
		switch typ {
		case 0:
			user.Phone = account
			res := db.Where("phone = ?", account).First(&model.User{})
			if res.RowsAffected == 0 {
				result := db.Table("user").Create(&user)
				if result.Error != nil {
					log.Error(`Create user error: %v`, result.Error)
					return
				}
			} else {
				log.Error(`Account already exists: %v`, account)
			}
		case 1:
			user.Email = account
			res := db.Where("email = ?", account).First(&model.User{})
			if res.RowsAffected == 0 {
				result := db.Table("user").Create(&user)
				if result.Error != nil {
					log.Error(`Create user error: %v`, result.Error)
					return
				}
			} else {
				log.Error(`Account already exists: %v`, account)
			}
		case 2:
			user.Wechat = account
			res := db.Where("wechat = ?", account).First(&model.User{})
			if res.RowsAffected == 0 {
				result := db.Table("user").Create(&user)
				if result.Error != nil {
					log.Error(`Create user error: %v`, result.Error)
					return
				}
			} else {
				log.Error(`Account already exists: %v`, account)
			}
		}

		return
	})
	if err != nil {
		log.Warn(`Register error: %v`, err)
	}
	return
}

func loginWithPhone(ctx context.Context, phone, password string) (resp *user_center.LoginResponse, err error) {
	user, err := getUser(ctx, model.User{Phone: phone})
	if err != nil {
		return nil, errors.New("Account Not Exist")
	}

	if user.Password != fmt.Sprintf("%x", md5.Sum([]byte(password))) {
		return nil, errors.New("Account Or Password Wrong")
	}
	return nil, nil
}

func loginWithEmail(ctx context.Context, email, password string) (resp *user_center.LoginResponse, err error) {
	user, err := getUser(ctx, model.User{Email: email})
	if err != nil {
		log.Debug(`User Not exists: %v`, email)
		return nil, err
	}

	if user.Password != fmt.Sprintf("%x", md5.Sum([]byte(password))) {
		return nil, errors.New("Account Or Password Wrong")
	}
	return nil, nil
}

func loginWithWechat(ctx context.Context, wechat, password string) (resp *user_center.LoginResponse, err error) {
	user, err := getUser(ctx, model.User{Wechat: wechat})
	if err != nil {
		log.Debug(`User Not exists: %v`, wechat)
		return nil, err
	}

	if user.Password != fmt.Sprintf("%x", md5.Sum([]byte(password))) {
		return nil, errors.New("Account Or Password Wrong")
	}
	return nil, nil
}

func login() {}

func getUser(ctx context.Context, user model.User) (res model.User, err error) {
	db := db.WithContext(ctx)
	err = db.Model(&user).Where(&user).Take(&res).Error
	return
}
