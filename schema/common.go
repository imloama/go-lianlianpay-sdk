package schema

// 随机因子获取
//https://open.lianlianpay.com/docs/accp/accpstandard/get-random.html

type RandomKeyRequest struct {
	Timestamp        string `json:"timestamp"`         //时间戳，格式yyyyMMddHHmmss HH以24小时为准，如20170309143712。timestamp 与连连服务器的时间(北京时间)之间的误差不能超过30分钟。
	OidPartner       string `json:"oid_partner"`       //商户号，ACCP系统分配给平台商户的唯一编号
	UserId           string `json:"user_id"`           // 用户在商户系统中的唯一编号，要求该编号在商户系统能唯一标识用户。由商户自定义。
	FlagChnl         string `json:"flag_chnl"`         // flag_chnl	必传	String	10	交易发起渠道。 	ANDROID	IOS	H5	PC
	PkgName          string `json:"pkg_name"`          //pkg_name	可选	String	256	APP包名。flag_chnl为H5时，送商户一级域名
	AppName          string `json:"app_name"`          //app_name	可选	String	256	APP应用名。flag_chnl为H5时，送商户一级域名
	EncryptAlgorithm string `json:"encrypt_algorithm"` //encrypt_algorithm	可选	String	16	加密算法。 RSA：国际通用RSA算法 SM2 ：国密算法 默认 RSA算法
}

type RandomKeyResponse struct {
	Ret
	OidPartner       string `json:"oid_partner"`        //oid_partnerString长度16必返回 	ACCP系统分配给平台商户的唯一编号。
	UserId           string `json:"user_id"`            // user_idString长度64必返回 商户用户唯一编号。用户在商户系统中的唯一编号，要求该编号在商户系统能唯一标识用户。
	RandomKey        string `json:"random_key"`         // random_key String长度不限请求受理成功时返回	随机因子key，有效期30分钟。
	RandomValue      string `json:"random_value"`       //random_value String长度32请求受理成功时返回	随机因子值，有效期30分钟。
	License          string `json:"license"`            //license  String长度1024请求受理成功时返回	license。flag_chnl为ANDROID、IOS、H5时返回
	MapArr           string `json:"map_arr"`            // String长度1024请求受理成功且flag_chnl为H5时返回	映射数组
	RsaPublicContent string `json:"rsa_public_content"` //String长度1024请求受理成功时返回	连连RSA公钥。
	Sm2KeyHex        string `json:"sm2_key_hex"`        //String长度3使用国密算法返回	使用国密算法返回
}


func(req *RandomKeyRequest)GetUrl()string{
	return "/v1/acctmgr/get-random"
}

