package schema


type IRequest interface{
	GetUrl()string
}

// 请求结果状态码 https://open.lianlianpay.com/docs/accp/accpstandard/accp-returncodes.html

type Ret struct {
	RetCode string `json:"ret_code"` //0000
	RetMsg string `json:"ret_msg"` // String  //长度32必返回 	请求结果描述
}

func(ret Ret)GetRetCode()string{
	return ret.RetCode
}
func(ret Ret)GetRetMsg()string{
	return ret.RetMsg
}

type IResult interface {
	GetRetCode() string
	GetRetMsg() string
}