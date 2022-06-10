package lianlianpay

import (
	"github.com/imloama/go-lianlianpay-sdk/schema"
)

// 执行随机因子获取接口
// https://open.lianlianpay.com/docs/accp/accpstandard/get-random.html
func(client *LianlianPayClient) ExecRandomKey(reqdata *schema.RandomKeyRequest, result *schema.RandomKeyResponse)error{
	return client.Exec(reqdata, result)

	//signature, err := client.SignData(reqdata)
	//if err!=nil{
	//	return nil, err
	//}
	//url := client.Gateway+reqdata.GetUrl()
	//var result schema.RandomKeyResponse
	//rsp, err := client.reqclient.R().SetHeader("Signature-Data", signature).SetBody(reqdata).
	//	SetResult(&result).EnableDump().
	//	Post(url)
	//if err!=nil {
	//	return nil, err
	//}
	//if rsp.IsSuccess() {
	//	if result.RetCode!= RET_CODE_OK {
	//		return nil, errors.New(fmt.Sprintf("[%s]%s", result.RetCode, RET_CODE_ERRORS[result.RetCode]))
	//	}
	//	return &result, nil
	//}
	//return nil, errors.New(rsp.Dump())
}
