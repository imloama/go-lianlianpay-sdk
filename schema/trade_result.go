package schema

//支付结果查询
//https://open.lianlianpay.com/docs/accp/accpstandard/query-payment.html
type QueryPaymentRequest struct{
	Timestamp  string `json:"timestamp"` // timestamp	必传	Date	14	时间戳，格式yyyyMMddHHmmss HH以24小时为准，如20170309143712。timestamp 与连连服务器的时间(北京时间)之间的误差不能超过30分钟。
	OidPartner string `json:"oid_partner"` //id_partner	必传	String	16	商户号，ACCP系统分配给平台商户的唯一编号。测试环境商户号
	TxnSeqno string `json:"txn_seqno"` // 商户系统唯一交易流水号。【三选一】，建议优先使用ACCP系统交易单号。
	AccpTxno   string `json:"accp_txno"` // ACCP系统交易单号。【三选一】，建议优先使用ACCP系统交易单号。
}

func(s *QueryPaymentRequest)GetUrl()string{
	return "/v1/txn/query-payment"
}

type QueryPaymentResponse struct {
	Ret
	OidPartner     string `json:"oid_partner"`
	AccountingDate string `json:"accounting_date"`
	FinishTime     string `json:"finish_time"`
	AccpTxno       string `json:"accp_txno"`
	TxnStatus      string `json:"txn_status"`
	Bankcode       string `json:"bankcode"`
	OrderInfo      struct { //商户订单信息orderInfo
		TxnSeqno    string `json:"txn_seqno"`
		TxnTime     string `json:"txn_time"`
		TotalAmount string `json:"total_amount"`
		OrderInfo   string `json:"order_info"`
	} `json:"orderInfo"`
	PayerInfo []struct { //付款方信息（组合支付场景返回付款方信息数组）payerInfo
		PayerType string `json:"payer_type"` //付款方类型。		用户：USER		平台商户：MERCHANT
		PayerID   string `json:"payer_id"`
		Method    string `json:"method"`
		Amount    string `json:"amount"`
	} `json:"payerInfo"`
	PayeeInfo []struct { //收款方信息（交易分账场景返回收款方信息数组）payeeInfo
		PayeeType string `json:"payee_type"`
		PayeeID   string `json:"payee_id"`
		Amount    string `json:"amount"`
	} `json:"payeeInfo"`
}
