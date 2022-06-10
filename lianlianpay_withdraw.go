package lianlianpay

import (
	"github.com/imloama/go-lianlianpay-sdk/schema"
)

func(client *LianlianPayClient)WithdrawRequest(reqdata *schema.WithdrawalRequest,result *schema.WithdrawalResponse)error{
	return client.Exec(reqdata, result)
}