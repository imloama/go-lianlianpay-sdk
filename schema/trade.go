package schema

// 支付统一创单
// https://open.lianlianpay.com/docs/accp/accpstandard/accp-tradecreate.html

const (
	TXN_TYPE_USERCHARGE = "USER_TOPUP" // 用户充值
	TXN_TYPE_MCHCHARGE = "MCH_TOPUP" // 商户充值
	TXN_TYPE_GENERAL_CONSUME = "GENERAL_CONSUME" // 普通消费
	TXN_TYPE_SECURED_CONSUME = "SECURED_CONSUME" // 担保消费
)


type TradeCreateRequest struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	UserId     string `json:"user_id"`
	UserType   string `json:"user_type"`
	TxnType    string `json:"txn_type"` // 交易类型 用户充值：USER_TOPUP 商户充值：MCH_TOPUP 普通消费：GENERAL_CONSUME 担保消费：SECURED_CONSUME
	NotifyURL  string `json:"notify_url"`
	PayExpire  string `json:"pay_expire"`
	OrderInfo  TradeCreateOrderInfo `json:"orderInfo"`
	PayeeInfo []TradeCreatePayeeInfo `json:"payeeInfo"`
}

type TradeCreateOrderInfo struct{
	TxnSeqno    string `json:"txn_seqno"`
	TxnTime     string `json:"txn_time"`
	TotalAmount float64 `json:"total_amount"`
	FeeAmount   float64 `json:"fee_amount"` //可选,手续费金额，只支持充值交易。单位为元，精确到小数点后两位。会自动收取到平台商户的自有资金账户。支持低于订单总金额的20%，或者低于10元但不高于订单金额。
	OrderInfo   string `json:"order_info"`// 可选,订单信息，在查询API和支付通知中原样返回，可作为自定义参数使用。
	GoodsName   string `json:"goods_name"`// 可选,商品描述信息。
	GoodUrl string `json:"goods_url"` // 可选,订单及商品展示地址。
}

type TradeCreatePayeeInfo struct {
	PayeeId     string `json:"payee_id"`
	PayeeType   string `json:"payee_type"`
	PayeeAccttype string `json:"payee_accttype"` // 收款方账户类型。交易类型为商户充值时必须指定充值入账账户类型 用户账户：USEROWN 平台商户自有资金账户：MCHOWN 平台商户优惠券账户：MCHCOUPON 平台商户手续费账户：MCHFEE
	PayeeAmount string `json:"payee_amount"`
	PayeeMemo   string `json:"payee_memo"`
}

func(s *TradeCreateRequest)GetUrl()string{
	return "/v1/txn/tradecreate"
}

type TradeCreateResponse struct {
	Ret
	OidPartner  string `json:"oid_partner"`
	UserId      string `json:"user_id"`
	TotalAmount string `json:"total_amount"`
	TxnSeqno    string `json:"txn_seqno"`
	AccpTxno    string `json:"accp_txno"`
}

//==========网关类支付
// https://open.lianlianpay.com/docs/accp/accpstandard/payment-gw.html

type PaymentGWRequest struct {
	Timestamp   string `json:"timestamp"`
	OidPartner  string `json:"oid_partner"`
	TxnSeqno    string `json:"txn_seqno"`
	TotalAmount string `json:"total_amount"`
	Appid       string `json:"appid"` // 应用ID。开发者在微信或支付宝开放平台申请的APPID。非网银类，除扫码支付，都必传。微信平台申请APPID
	Openid      string `json:"openid"` // 渠道用户标识。微信公众号和微信小程序支付时必传。此参数为微信用户在商户对应APPID下的唯一标识或支付宝买家的支付宝唯一用户号（2088开头的16位纯数字）。
	ClientIP    string `json:"client_ip"`
	Bankcode string `json:"bankcode"` // 银行编码。付款方式为网银类时可指定。
	DeviceInfo string `json:"device_info"` // 设备标识。自定义参数，可以为终端设备号。
	DeviceVersion string `json:"device_version"` // 版本标识。自定义参数，可以为支付宝的版本号。
	LimitCardType string `json:"limit_card_type"` // 限制卡类型。限制某种卡类型支付权限，支付宝和微信支付都适用，若限制多种类型以“,”分隔（暂时只支持限制信用卡支付）。credit：限制适用信用卡支付。
	ExtendParams string `json:"extend_params"` // 业务扩展字段。渠道扩展字段JSON串，若渠道为支付宝，则支持的键值对如下：sys_service_provider_id	hb_fq_num	hb_fq_seller_percent	industry_reflux_info	card_type	支付宝微信同时支持键值：accp_sub_mch_id 子商户号 微信H5扩展参数：req_domain 来源域
	PayerInfo   PayerInfo `json:"payerInfo"`
	PayMethods []PayMethod `json:"payMethods"`
}

type PayerInfo struct {
	PayerType string `json:"payer_type"`
	PayerId   string `json:"payer_id"`
	Password  string `json:"password"`
	RandomKey string `json:"random_key"`
	PayerAccttype string `json:"payer_accttype"`// payer_accttype	可选	String	不限	付款方账户类型。付款方类型为商户时需要指定平台商户账户类型。	//USEROWN：用户自有可用账户。	//MCHOWN：平台商户自有可用账户。
	PapAgreeNo string `json:"pap_agree_no"` // 委托代发协议号。用户在ACCP开通的委托提现协议号。	通过用户委托协议签约接口获取。	授权免密提现时必输。该字段需要RSA公钥加密传输。
}

type PayMethod struct {
	Method string `json:"method"`
	Amount string `json:"amount"`
}

func(s *PaymentGWRequest)GetUrl()string{
	return "/v1/txn/payment-gw"
}

type PaymentGWResponse struct {
	Ret
	OidPartner  string `json:"oid_partner"`
	UserId      string `json:"user_id"`
	TxnSeqno    string `json:"txn_seqno"`
	TotalAmount string `json:"total_amount"`
	AccpTxno    string `json:"accp_txno"`
	Token       string `json:"token"`
	Payload     string `json:"payload"`
}


const PAY_METHOD_BALANCE = "BALANCE"
const PAY_METHOD_WECHAT_APP_WXA = "WECHAT_APP_WXA"
const PAY_METHOD_AGGREGATE_CODE = "AGGREGATE_CODE"
const PAY_METHOD_WECHAT_APP = "WECHAT_APP"
const PAY_METHOD_WECHAT_NATIVE = "WECHAT_NATIVE"
const PAY_METHOD_WECHAT_H5 = "WECHAT_H5"
const PAY_METHOD_WECHAT_WXA = "WECHAT_WXA"
const PAY_METHOD_ALIPAY_NATIVE = "ALIPAY_NATIVE"
const PAY_METHOD_ALIPAY_APP = "ALIPAY_APP"
const PAY_METHOD_ALIPAY_H5 = "ALIPAY_H5"
const PAY_METHOD_CUP_YUNSHANFU = "CUP_YUNSHANFU"

/**

支持的付款方式
编码	描述
BALANCE	余额
COUPON	优惠券
WECHAT_APP_WXA	微信APP（小程序）
AGGREGATE_CODE	聚合码支付方式（微信、支付宝及云闪付）
BANK_CARD_PAY	银行卡收银台
EBANK_DEBIT_CARD	网银借记卡
EBANK_CREDIT_CARD	网银信用卡
EBANK_B2B	企业网银
WECHAT_APP	微信APP
WECHAT_JSAPI	微信公众号
WECHAT_NATIVE	微信扫码
WECHAT_H5	微信H5
WECHAT_WXA	微信小程序
WECHAT_JSAPI_OFFLINE	微信公众号-线下（富民）
WECHAT_NATIVE_OFFLINE	微信扫码-线下（富民）
WECHAT_WXA_OFFLINE	微信小程序-线下（富民）
ALIPAY_JSAPI	支付宝 JSAPI（富民）
ALIPAY_NATIVE	支付宝扫码
ALIPAY_APP	支付宝APP
ALIPAY_H5	支付宝H5
ALIPAY_WEB	支付宝WEB
ALIPAY_WXA	支付宝小程序
CUP_YUNSHANFU	银联云闪付（APP）

 */

