package pay

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/yeezyi/wechat-sdk/tools"
	"strconv"
	"time"
)

// 统一下单
// 详细规则参考  https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
func unifiedOrder(data unifiedOrderData, mchKey string, sandbox bool) (*unifiedOrderResponse, error) {
	if data.OutTradeNo == "" {
		return nil, outTradeNoIsNull
	}
	if data.Body == "" {
		return nil, bodyIsNull
	}
	if data.TotalFee == 0 {
		return nil, totalFeeIsNull
	}
	if data.TradeType == "" {
		return nil, tradeTypeIsNull
	}
	if data.SpbillCreateIP == "" {
		return nil, spbillCreateIPIsNull
	}
	if data.TradeType == "JSAPI" && data.OpenID == "" {
		return nil, tradeTypeJSAPIButOpenidIsNull
	}
	if data.TradeType == "NATIVE" && data.ProductId == "" {
		return nil, tradeTypeNativeButProductIdIsNull
	}
	dataMap := map[string]string{
		"appid":            data.AppID,
		"mch_id":           data.MchID,
		"nonce_str":        data.NonceStr,
		"body":             data.Body,
		"out_trade_no":     data.OutTradeNo,
		"total_fee":        strconv.Itoa(int(data.TotalFee)),
		"spbill_create_ip": data.SpbillCreateIP,
		"notify_url":       data.NotifyUrl,
		"trade_type":       data.TradeType,
	}
	if data.TradeType == "JSAPI" {
		dataMap["openid"] = data.OpenID
	} else if data.TradeType == "NATIVE" {
		dataMap["product_id"] = data.ProductId
	}
	if data.Attach != "" {
		dataMap["attach"] = data.Attach
	}
	sign := createSign(dataMap, mchKey)
	data.Sign = sign
	url := unifiedOrderUrl
	if sandbox {
		url = sandboxUnifiedOrderUrl
	}
	xmlData, err := xml.Marshal(data)
	if err != nil {
		return nil, err
	}
	body, err := tools.Post(url, xmlData)
	if err != nil {
		return nil, err
	}
	response := new(unifiedOrderResponse)
	err = xml.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}
	if response.ReturnCode != "SUCCESS" {
		return response, errors.New(response.ReturnMsg)
	}
	return response, err
}

// H5 支付
// 详细规则参考 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=7_7&index=6
func (c *WxClient) JsApiPay(req JsApiRequest) (*jsApiResponse, error) {
	data := unifiedOrderData{
		AppID:          c.AppID,
		MchID:          c.MchID,
		NonceStr:       tools.GenerateNonceStr(),
		NotifyUrl:      req.NotifyUrl,
		Body:           req.Body,
		OutTradeNo:     req.OutTradeNo,
		TotalFee:       req.TotalFee,
		SpbillCreateIP: req.SpbillCreateIP,
		TradeType:      "JSAPI",
		OpenID:         req.OpenID,
		Attach:         req.Attach,
	}
	resp, err := unifiedOrder(data, c.MchKey, c.SandBox)
	if err != nil {
		return nil, err
	}
	response := jsApiResponse{
		AppID:     c.AppID,
		Timestamp: strconv.Itoa(int(time.Now().Unix())),
		NonceStr:  tools.GenerateNonceStr(),
		Package:   fmt.Sprintf("prepay_id=%s", resp.PrepayID),
		SignType:  "MD5",
	}
	raw := map[string]string{
		"appId":     c.AppID,
		"timeStamp": response.Timestamp,
		"nonceStr":  response.NonceStr,
		"package":   response.Package,
		"signType":  "MD5",
	}
	paySign := createSign(raw, c.MchKey)
	response.PaySign = paySign
	return &response, nil
}

// 小程序支付
// 详细规则参考 https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_1&index=1
func (c *WxClient) XCXPay(req JsApiRequest) (*jsApiResponse, error) {
	if response, err := c.JsApiPay(req); err != nil {
		return nil, err
	} else {
		response.AppID = ""
		return response, nil
	}
}

// 生成给App调用的数据
// 详细规则参考 https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_12&index=2
func (c *WxClient) AppPay(req AppRequest) (*appResponse, error) {
	data := unifiedOrderData{
		AppID:          c.AppID,
		MchID:          c.MchID,
		NonceStr:       tools.GenerateNonceStr(),
		NotifyUrl:      req.NotifyUrl,
		Body:           req.Body,
		OutTradeNo:     req.OutTradeNo,
		TotalFee:       req.TotalFee,
		SpbillCreateIP: req.SpbillCreateIP,
		TradeType:      "APP",
		Attach:         req.Attach,
	}
	resp, err := unifiedOrder(data, c.MchKey, c.SandBox)
	if err != nil {
		return nil, err
	}
	response := appResponse{
		AppID:     c.AppID,
		PartnerID: c.MchID,
		PrepayID:  resp.PrepayID,
		Package:   "Sign=WXPay",
		NonceStr:  tools.GenerateNonceStr(),
		Timestamp: strconv.Itoa(int(time.Now().Unix())),
	}
	raw := map[string]string{
		"appid":     c.AppID,
		"partnerid": c.MchID,
		"prepayid":  response.PrepayID,
		"package":   response.Package,
		"noncestr":  response.NonceStr,
		"timestamp": response.Timestamp,
	}
	paySign := createSign(raw, c.MchKey)
	response.Sign = paySign
	return &response, nil
}
