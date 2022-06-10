package schema

//提现申请
// https://open.lianlianpay.com/docs/accp/accpstandard/withdrawal.html

type WithdrawalRequest struct {
	Timestamp    string            `json:"timestamp"`
	OidPartner   string            `json:"oid_partner"`
	NotifyURL    string            `json:"notify_url"`
	LinkedAcctno string            `json:"linked_acctno"` //银行账号
	LinkedAgrtno string            `json:"linked_agrtno"` // 绑卡协议号
	FundsFlag    string            `json:"funds_flag"`
	CheckFlag    string            `json:"check_flag"`    // 审核标识。标识该笔订单是否需要审核，默认:N	Y:需要提现确认	N：无需提现确认
	PayTimeType  string            `json:"pay_time_type"` // 到账类型。	TRANS_THIS_TIME :实时到账	TRANS_NORMAL :普通到账（2小时内）	TRANS_NEXT_TIME :次日到账 	默认：实时到账
	OrderInfo    WithdrawOrderInfo `json:"orderInfo"`
	PayerInfo    PayerInfo         `json:"payerInfo"`
}

func (s *WithdrawalRequest) GetUrl() string {
	return "/v1/txn/withdrawal"
}

type WithdrawOrderInfo struct {
	TxnSeqno    string  `json:"txn_seqno"`
	TxnTime     string  `json:"txn_time"`
	TotalAmount string  `json:"total_amount"`
	FeeAmount   float64 `json:"fee_amount"` //可选,手续费金额，只支持充值交易。单位为元，精确到小数点后两位。会自动收取到平台商户的自有资金账户。支持低于订单总金额的20%，或者低于10元但不高于订单金额。
	Memo        string  `json:"memo"`       // Postscript or memo
	OrderInfo   string  `json:"order_info"` // 订单信息。用于订单说明，透传返回。
}

type WithdrawalResponse struct {
	Ret
	OidPartner  string `json:"oid_partner"`
	UserID      string `json:"user_id"`
	TxnSeqno    string `json:"txn_seqno"`
	TotalAmount string `json:"total_amount"`
	AccpTxno    string `json:"accp_txno"`
}

// 连连支付提现回调数据
type LianLianWithdrawNotifyResult struct {
	OidPartner     string `json:"oid_partner"`     // 商户号
	UserID         string `json:"user_id"`         // 用户在商户系统中的唯一编码
	AccountingDate string `json:"accounting_date"` // 账务日期
	FinishTime     string `json:"finish_time"`     // 提现成功时间。格式：yyyyMMddHHmmss
	AccpTxno       string `json:"accp_txno"`       // ACCP系统交易单号
	ChnlTxno       string `json:"chnl_txno"`       // 渠道交易单号
	TxnStatus      string `json:"txn_status"`      // 提现交易状态。TRADE_SUCCESS:交易成功 TRADE_FAILURE:交易失败 TRADE_CANCEL：退汇
	FailureReason  string `json:"failure_reason"`  // 提现失败原因
	Bankcode       string `json:"bankcode"`        // 银行编码
	ChnlReason     string `json:"chnl_reason"`     // 提现失败渠道原始原因
	PayerInfo      struct {
		PayerType string `json:"payer_type"`
		PayerId   string `json:"payer_id"`
	}
	OrderInfo struct {
		TxnSeqno    string `json:"txn_seqno"`
		TxnTime     string `json:"txn_time"`
		TotalAmount string `json:"total_amount"`
		OrderInfo   string `json:"order_info"` // 订单信息。用于订单说明，透传返回。
	}
}
