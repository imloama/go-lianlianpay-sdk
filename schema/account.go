package schema

//绑定手机验证码申请 全必填
type GetPhoneVerifyCodeRequest struct {
	Timestamp  string `json:"timestamp"`   //时间戳，格式yyyyMMddHHmmss HH以24小时为准，如20170309143712。timestamp 与连连服务器的时间(北京时间)之间的误差不能超过30分钟。
	OidPartner string `json:"oid_partner"` //商户号，ACCP系统分配给平台商户的唯一编号
	UserID     string `json:"user_id"`     //用户在商户系统中的唯一编号，要求该编号在商户系统能唯一标识用户。由商户自定义。
	RegPhone   string `json:"reg_phone"`   //绑定手机号。用户开户注册绑定手机号。
}

type GetPhoneVerifyCodeResponse struct {
	//RetCode string `json:"ret_code"` // 请求结果代码
	//RetMsg string `json:"ret_msg"`
	Ret
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	RegPhone   string `json:"reg_phone"`
}

func (req *GetPhoneVerifyCodeRequest) GetUrl() string {
	return "/v1/acctmgr/regphone-verifycode-apply"
}

// ========= 验证手机号验证码 https://open.lianlianpay.com/docs/accp/accpstandard/regphone-verifycode-verify.html

type PhoneVerifyCodeRequest struct {
	Timestamp  string `json:"timestamp"`   // 时间戳，格式yyyyMMddHHmmss HH以24小时为准，如20170309143712。timestamp 与连连服务器的时间(北京时间)之间的误差不能超过30分钟。
	OidPartner string `json:"oid_partner"` // 商户号，ACCP系统分配给平台商户的唯一编号
	UserID     string `json:"user_id"`     // 用户在商户系统中的唯一编号，要求该编号在商户系统能唯一标识用户。由商户自定义。
	VerifyCode string `json:"verify_code"` // 绑定手机号。用户开户注册绑定手机号。
	RegPhone   string `json:"reg_phone"`   // 绑定手机号验证码。通过绑定手机验证码申请接口申请发送给用户绑定手机的验证码。
}

type PhoneVerifyCodeResponse struct {
	//RetCode string `json:"ret_code"` // 请求结果代码
	//RetMsg string `json:"ret_msg"`
	Ret
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	RegPhone   string `json:"reg_phone"`
}

func (req *PhoneVerifyCodeRequest) GetUrl() string {
	return "/v1/acctmgr/regphone-verifycode-apply"
}

//=====开户申请  https://open.lianlianpay.com/docs/accp/accpstandard/openacct-apply-individual.html
type ApplyIndividualRequest struct {
	Timestamp      string         `json:"timestamp"`     // 时间戳，格式yyyyMMddHHmmss HH以24小时为准，如20170309143712。timestamp 与连连服务器的时间(北京时间)之间的误差不能超过30分钟。
	OidPartner     string         `json:"oid_partner"`   // 商户号，ACCP系统分配给平台商户的唯一编号
	UserID         string         `json:"user_id"`       // 用户在商户系统中的唯一编号，要求该编号在商户系统能唯一标识用户。由商户自定义。
	TxnSeqno       string         `json:"txn_seqno"`     // 商户系统唯一交易流水号。由商户自定义。
	TxnTime        string         `json:"txn_time"`      // 商户系统交易时间。	格式：yyyyMMddHHmmss。
	NotifyURL      string         `json:"notify_url"`    // 交易结果异步通知接收地址，建议HTTPS协议。
	OpenSmsFlag    string         `json:"open_sms_flag"` // 开户成功短信通知。Y:发送 N：不发送（默认）可选  该功能需要合作银行支持，暂只满足富民银行
	RiskItem       string         `json:"risk_item"`     //风险控制参数。连连风控部门要求商户统一传入风险控制参数字段，字段值为json 字符串的形式
	BasicInfo      BasicInfo      `json:"basicInfo"`
	LinkedAcctInfo LinkedAcctInfo `json:"linkedAcctInfo"`
	AccountInfo    ACSAccountInfo `json:"accountInfo"`
}

// 开户基本信息
type BasicInfo struct {
	RegPhone           string `json:"reg_phone"` // 必传 绑定手机号。用户开户注册绑定手机号。
	UserName           string `json:"user_name"` // 必传 用户名称。企业用户传企业名称，个体工商户传营业执照的名称（营业执照上名称是****或者无的，传经营者姓名），个人用户传姓名。
	IDType             string `json:"id_type"`   // 必传 证件类型  身份证：ID_CARD。
	IDNo               string `json:"id_no"`     // 必传 证件号码。
	IDExp              string `json:"id_exp"`    // id_exp	必传	String	8	证件有效期。证件到期日，格式：yyyyMMdd 长期有效则设置为：99991231
	RegPhoneVerifycode string `json:"reg_phone_verifycode"`
	Occupation         string `json:"occupation"`
}

// 开户绑卡信息
type LinkedAcctInfo struct {
	LinkedAcctno string `json:"linked_acctno"`
	LinkedPhone  string `json:"linked_phone"`
}

type ACSAccountInfo struct {
	AccountType      string `json:"account_type"`       // 账户属性。
	AccountNeedLevel string `json:"account_need_level"` // 需求等级。当账户属性为个人支付账户时，需要上传账户等级。	注：该字段表示商户请求的个人支付账户需求等级，真实开户结果以ACCP返回等级为准。
}

// 职业类型 https://open.lianlianpay.com/docs/accp/accpstandard/openacct-apply-individual.html#%E8%81%8C%E4%B8%9A%E7%B1%BB%E5%9E%8B
const OCCUPATION_CIVIL_SERVANT = "01" // 公务员
const OCCUPATION_EMP = "10"           // 公司员工
const OCCUPATION_FREELANCER = "10"    // 自由职业者
const OCCUPATION_OTHER = "19"         // 其他

type ApplyIndividualResponse struct {
	//RetCode string `json:"ret_code"`
	//RetMsg string `json:"ret_msg"`
	Ret
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	TxnSeqno   string `json:"txn_seqno"` // 商户系统唯一交易流水号
	AccpTxno   string `json:"accp_txno"` // ACCP系统交易单号。
	Token      string `json:"token"`     // 授权令牌。有效期30分钟。
}

func (s *ApplyIndividualRequest) GetUrl() string {
	return "/v1/acctmgr/openacct-apply-individual"
}

// 个人用户开户验证 https://open.lianlianpay.com/docs/accp/accpstandard/openacct-verify-individual.html

type VerifyApplyIndividualRequest struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	TxnSeqno   string `json:"txn_seqno"`
	Token      string `json:"token"`
	VerifyCode string `json:"verify_code"` // 银行预留手机短信验证码。开户申请时提供的绑卡信息中的银行预留手机发送的短信验证码，用于验证开户绑卡。
	Password   string `json:"password"`
	RandomKey  string `json:"random_key"`
}

func (s *VerifyApplyIndividualRequest) GetUrl() string {
	return "/v1/acctmgr/openacct-verify-individual"
}

// 开户验证返回结果
type VerifyApplyIndividualResponse struct {
	//RetCode string `json:"ret_code"`
	//RetMsg string `json:"ret_msg"`
	Ret
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	TxnSeqno   string `json:"txn_seqno"` // 商户系统唯一交易流水号
	AccpTxno   string `json:"accp_txno"`
	OidUserno  string `json:"oid_userno"`
	UserStatus string `json:"user_status"`
}

/*
user_statusString长度32请求受理成功时返回
用户状态。
ACTIVATE_PENDING :已登记或开户失败（原待激活）
CHECK_PENDING :审核中（原待审核）
REMITTANCE_VALID_PENDING :审核通过，待打款验证（企业用户使用，暂未要求）
NORMAL :正常
CANCEL :销户
PAUSE :暂停
ACTIVATE_PENDING_NEW ：待激活
*/
const USER_STATUS_PENDING = "ACTIVATE_PENDING"
const USER_STATUS_NORMAL = "NORMAL"
const USER_STATUS_CHECK = "CHECK_PENDING"

//==========个人用户信息修改=======================
// https://open.lianlianpay.com/docs/accp/accpstandard/modify-userinfo-individual.html
type ModifyUserInfoRequest struct {
	Timestamp  string    `json:"timestamp"`
	OidPartner string    `json:"oid_partner"`
	UserID     string    `json:"user_id"`
	TxnSeqno   string    `json:"txn_seqno"`  // 必传	String	64	商户系统唯一交易流水号。由商户自定义。
	TxnTime    string    `json:"txn_time"`   // 必传	String	14	商户系统交易时间。	格式：yyyyMMddHHmmss
	NotifyURL  string    `json:"notify_url"` // 必传 交易结果异步通知接收地址，建议HTTPS协议。
	BasicInfo  BasicInfo `json:"basicInfo"`
}

func (s *ModifyUserInfoRequest) GetUrl() string {
	return "/v1/acctmgr/openacct-verify-individual"
}

type ModifyUserInfoResponse struct {
	//RetCode string `json:"ret_code"`
	//RetMsg string `json:"ret_msg"`
	Ret
	Token     string `json:"token"`
	OidPartne string `json:"oid_partne"`
	UserID    string `json:"user_id"`
	TxnSeqno  string `json:"txn_seqno"`
	AccpTxno  string `json:"accp_txno"`
}

//==========个人用户新增绑卡=================
// https://open.lianlianpay.com/docs/accp/accpstandard/individual-bindcard-apply.html
type IndividualBindcardApplyRequest struct {
	Timestamp    string `json:"timestamp"`
	OidPartner   string `json:"oid_partner"`
	UserID       string `json:"user_id"`
	TxnSeqno     string `json:"txn_seqno"` // 商户系统唯一交易流水号。由商户自定义。
	TxnTime      string `json:"txn_time"`  // 商户系统交易时间。	格式：yyyyMMddHHmmss
	NotifyURL    string `json:"notify_url"`
	LinkedAcctno string `json:"linked_acctno"` // linked_acctno	可选	String	32	绑定银行账号。个人用户绑定的银行卡号。
	LinkedPhone  string `json:"linked_phone"`  // linked_phone	必传	String	11	银行预留手机号。
	Password     string `json:"password"`
	RandomKey    string `json:"random_key"`
}

func (s *IndividualBindcardApplyRequest) GetUrl() string {
	return "/v1/acctmgr/individual-bindcard-apply"
}

type IndividualBindcardApplyResponse struct {
	Ret
	Token     string `json:"token"` // 请求受理成功时返回	授权令牌。有效期30分钟。
	OidPartne string `json:"oid_partne"`
	UserID    string `json:"user_id"`
	TxnSeqno  string `json:"txn_seqno"`
	AccpTxno  string `json:"accp_txno"` // ACCP系统交易单号。
}

//绑卡验证
//https://open.lianlianpay.com/docs/accp/accpstandard/individual-bindcard-verify.html
type IndividualBindcardVerifyRequest struct {
	Timestamp  string `json:"timestamp"`
	OidPartner string `json:"oid_partner"`
	UserID     string `json:"user_id"`
	TxnSeqno   string `json:"txn_seqno"`
	VerifyCode string `json:"verify_code"`
	Token      string `json:"token"`
}

func (s *IndividualBindcardVerifyRequest) GetUrl() string {
	return "/v1/acctmgr/individual-bindcard-verify"
}

type IndividualBindcardVerifyResponse struct {
	Ret
	LinkedAgrtno string `json:"linked_agrtno"`
	OidPartne    string `json:"oid_partne"`
	UserID       string `json:"user_id"`
	TxnSeqno     string `json:"txn_seqno"`
	AccpTxno     string `json:"accp_txno"`
}

// 重置密码
// https://open.lianlianpay.com/docs/accp/accpstandard/change-password.html

type ChangePasswordRequest struct {
	Timestamp   string `json:"timestamp"` // 时间戳
	OidPartner  string `json:"oid_partner"`
	Password    string `json:"password"`     // 必传	String	12	密码。6-12位的字母、数字，不可以是连续或者重复的数字和字母，正则：^[a-zA-Z0-9]{6,12}$。
	PasswordNew string `json:"password_new"` // 必传	String	12	新支付密码。6-12位的字母、数字，不可以是连续或者重复的数字和字母，正则：^[a-zA-Z0-9]{6,12}$。
	RandomKey   string `json:"random_key"`   //必传	String	不限	密码随机因子key。随机因子获取接口返回。
	UserId      string `json:"user_id"`
}

func (s *ChangePasswordRequest) GetUrl() string {
	return "/v1/acctmgr/change-password"
}

type ChangePasswordResponse struct {
	Ret
	OidPartner string `json:"oid_partner"`
	serId      string `json:"user_id"`
}
