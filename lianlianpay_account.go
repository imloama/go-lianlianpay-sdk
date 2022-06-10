package lianlianpay

import (
	"github.com/imloama/go-lianlianpay-sdk/schema"
)

// 绑定手机验证码申请
func (client *LianlianPayClient) GetVerifyCode(reqdata *schema.GetPhoneVerifyCodeRequest, result *schema.GetPhoneVerifyCodeResponse) error {
	//var result schema.GetPhoneVerifyCodeResponse
	return client.Exec(reqdata, result)
}

func (client *LianlianPayClient) PhoneVerifyCode(reqdata *schema.PhoneVerifyCodeRequest, result *schema.PhoneVerifyCodeResponse) error {
	//var result schema.PhoneVerifyCodeResponse
	//err := client.Exec(reqdata, &result)
	//if err!=nil{
	//	return
	//}
	return client.Exec(reqdata, result)
}

func (client *LianlianPayClient) ApplyIndividual(reqdata *schema.ApplyIndividualRequest, result *schema.ApplyIndividualResponse) error {
	return client.Exec(reqdata, result)
}

// 验证用户开户申请
func (client *LianlianPayClient) VerifyApplyIndividual(reqdata *schema.VerifyApplyIndividualRequest, result *schema.VerifyApplyIndividualResponse) error {
	return client.Exec(reqdata, result)
}

// 修改用户信息
func (client *LianlianPayClient) ModifyUserInfo(reqdata *schema.ModifyUserInfoRequest, result *schema.ModifyUserInfoResponse) error {
	return client.Exec(reqdata, result)
}

//个人用户新增绑卡申请
func (client *LianlianPayClient) IndividualBindcardApply(reqdata *schema.IndividualBindcardApplyRequest, result *schema.IndividualBindcardApplyResponse) error {
	return client.Exec(reqdata, result)
}

//个人用户新增绑卡验证
func (client *LianlianPayClient) IndividualBindcardVerify(reqdata *schema.IndividualBindcardVerifyRequest, result *schema.IndividualBindcardVerifyResponse) error {
	return client.Exec(reqdata, result)
}

// 重置密码
func (client *LianlianPayClient) ChangePassword(reqdata *schema.ChangePasswordRequest, result *schema.ChangePasswordResponse) error {
	return client.Exec(reqdata, result)
}
