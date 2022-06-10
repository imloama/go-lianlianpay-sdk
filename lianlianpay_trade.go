package lianlianpay

import (
	"github.com/imloama/go-lianlianpay-sdk/schema"
)

func(client *LianlianPayClient)TradeCreate(reqdata *schema.TradeCreateRequest,result *schema.TradeCreateResponse)error{
	return client.Exec(reqdata, result)
}
// 查询商户详情
func(client *LianlianPayClient)QueryPayment(reqdata *schema.QueryPaymentRequest,result *schema.QueryPaymentResponse)error{
	return client.Exec(reqdata, result)
}
func(client *LianlianPayClient)PaymentGw(reqdata *schema.PaymentGWRequest,result *schema.PaymentGWResponse)error{
	return client.Exec(reqdata, result)
}
//type PaymentCallbackFn func(req *http.Request)
//func(client *LianlianPayClient)PaymentCallback(fn PaymentCallbackFn){
//
//}