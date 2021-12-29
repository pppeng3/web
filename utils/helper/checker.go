package helper

import (
	"errors"
	"regexp"
	"web/log"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tcerrors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tiia "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tiia/v20190529"
)

func CheckEmail(email string) error {
	if email != "" {
		pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
		reg := regexp.MustCompile(pattern)
		if reg.MatchString(email) == false {
			log.Error(`Wrong Email: %v`, email)
			return errors.New("Please Enter A Right Email")
		}
	}
	return nil
}

func CheckPhone(phone string) error {
	if phone != "" {
		regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
		pho := regexp.MustCompile(regular)
		if pho.MatchString(phone) == false {
			log.Error(`Wrong Phone: %v`, phone)
			return errors.New("Please Enter A Right Phone Number")
		}
	}
	return nil
}

// 检查头像
func CheckAvatarUrl(avatarurl string) error {
	if avatarurl != "" {
		credential := common.NewCredential(
			"AKIDvmU0VyPo5oSTO5ghhYHtPm0BHSZaPDh2", //SecretId
			"02OrCUginrnbIpmfS8fDewFVdcXM2gCf",     //SecretKey
		)
		cpf := profile.NewClientProfile()
		cpf.HttpProfile.Endpoint = "tiia.tencentcloudapi.com"
		client, _ := tiia.NewClient(credential, "ap-guangzhou", cpf)

		request := tiia.NewDetectMisbehaviorRequest()

		request.ImageUrl = common.StringPtr(avatarurl)

		response, err := client.DetectMisbehavior(request)
		if _, ok := err.(*tcerrors.TencentCloudSDKError); ok {
			log.Warn(`An API error has returned: %s`, err)
			return err
		}
		if err != nil {
			panic(err)
		}

		mini := float64(0.6)
		response.ToJsonString()
		if *response.Response.Confidence > mini {
			log.Info(`picture is not be used`)
			return errors.New("Please Use Another Pic")
		}
		log.Info("%s", response.ToJsonString())
	}
	return nil
}
