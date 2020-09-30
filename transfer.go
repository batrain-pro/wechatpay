package wechatpay

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

//企业付款到零钱
func (this *WechatPay) Transfer(param PromotionTransfersReq) (*PromotionTransfersRes, error) {
	param.AppId = this.AppId
	param.MchId = this.MchId
	param.NonceStr = randomNonceStr()

	var m map[string]interface{}
	m = make(map[string]interface{}, 0)
	m["mch_appid"] = param.AppId
	m["mchid"] = param.MchId
	m["device_info"] = param.DeviceInfo
	m["partner_trade_no"] = param.PartnerTradeNo
	m["openid"] = param.Openid
	m["check_name"] = param.CheckName
	m["re_user_name"] = param.ReUserName
	m["amount"] = param.Amount
	m["desc"] = param.Desc
	m["spbill_create_ip"] = param.SpbillCreateIp
	m["nonce_str"] = param.NonceStr
	param.Sign = GetSign(m, this.ApiKey)

	bytes_req, err := xml.Marshal(param)
	if err != nil {
		return nil, err
	}
	str_req := string(bytes_req)
	str_req = strings.Replace(str_req, "PromotionTransfersReq", "xml", -1)

	req, err := http.NewRequest("POST", PROMOTION_TRANSFERS_URL, bytes.NewReader([]byte(str_req)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")

	w_req := http.Client{
		Transport: WithCertBytes(this.ApiclientCert, this.ApiclientKey),
	}

	resp, err := w_req.Do(req)
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)

	var transfer_result PromotionTransfersRes
	err = xml.Unmarshal(body, &transfer_result)
	if err != nil {
		return nil, err
	} else if transfer_result.ReturnCode != "SUCCESS" {
		return nil, errors.New(transfer_result.ReturnMsg)
	} else if transfer_result.ResultCode != "SUCCESS" {
		return nil, errors.New(transfer_result.ErrCode)
	}

	return &transfer_result, nil
}
