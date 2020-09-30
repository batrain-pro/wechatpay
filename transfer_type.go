package wechatpay

const (
	//企业付款到零钱 - wiki:https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2
	PROMOTION_TRANSFERS_URL = "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers"
	//查询企业付款 - https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_3
	GET_TRANSFER_INFO_URL = "https://api.mch.weixin.qq.com/mmpaymkttransfers/gettransferinfo"
)

type PromotionTransfersReq struct {
	AppId    string `xml:"mch_appid"`
	MchId    string `xml:"mchid"`
	NonceStr string `xml:"nonce_str"`
	Sign     string `xml:"sign"`

	PartnerTradeNo string `xml:"partner_trade_no"`
	Openid         string `xml:"openid"`
	//NO_CHECK FORCE_CHECK
	CheckName string `xml:"check_name"`
	//FORCE_CHECK，则必填用户真实姓名
	ReUserName string `xml:"re_user_name"`
	Amount     int    `xml:"amount"`
	Desc       string `xml:"desc"`
	//非必填
	DeviceInfo     string `xml:"device_info"`
	SpbillCreateIp string `xml:"spbill_create_ip"`
}

type PromotionTransfersRes struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`

	AppId    string `xml:"mch_appid"`
	MchId    string `xml:"mchid"`
	NonceStr string `xml:"nonce_str"`

	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`

	//非必填
	DeviceInfo string `xml:"device_info"`

	//以下字段在return_code 和result_code都为SUCCESS的时候有返回
	PartnerTradeNo string `xml:"partner_trade_no"`
	//企业付款成功，返回的微信付款单号
	PaymentNo   string `xml:"payment_no"`
	PaymentTime string `xml:"payment_time"`
}
