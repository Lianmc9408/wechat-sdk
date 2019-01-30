package pay

import (
	"encoding/xml"
	"errors"
	"strconv"
	"wechat-sdk/tools"
)

// 订单查询
// 详细规则参考 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_2
func (c *WxClient) OrderQuery(req OrderQueryRequest) (*orderQueryResponse, error) {
	if req.OutTradeNo == "" && req.TransactionID == "" {
		return nil, outTradeNoAndTransactionIdIsNull
	}
	data := orderQueryData{
		AppID:         c.AppID,
		MchID:         c.MchID,
		NonceStr:      tools.GenerateNonceStr(),
		TransactionID: req.TransactionID,
		OutTradeNo:    req.OutTradeNo,
	}
	dataMap := map[string]string{
		"appid":     c.AppID,
		"mch_id":    c.MchID,
		"nonce_str": data.NonceStr,
	}
	if req.TransactionID != "" {
		dataMap["transaction_id"] = req.TransactionID
	} else {
		dataMap["out_trade_no"] = req.OutTradeNo
	}
	data.Sign = createSign(dataMap, c.MchKey)
	xmlData, err := xml.Marshal(data)
	if err != nil {
		return nil, err
	}
	url := orderQueryUrl
	if c.SandBox {
		url = sandboxOrderQueryUrl
	}
	body, err := tools.Post(url, xmlData)
	if err != nil {
		return nil, err
	}
	response := new(orderQueryResponse)
	err = xml.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}
	if response.ReturnCode != "SUCCESS" {
		return response, errors.New(response.ReturnMsg)
	}
	return response, err
}

//关闭订单
//详细规则参考 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_3
//注意：订单生成后不能马上调用关单接口，最短调用时间间隔为5分钟
func (c *WxClient) OrderClose(outTradeNo string) (*orderCloseResponse, error) {
	if outTradeNo == "" {
		return nil, outTradeNoIsNull
	}
	data := orderCloseData{
		AppID:      c.AppID,
		MchID:      c.MchID,
		NonceStr:   tools.GenerateNonceStr(),
		OutTradeNo: outTradeNo,
	}
	dataMap := map[string]string{
		"appid":        c.AppID,
		"mch_id":       c.MchID,
		"nonce_str":    data.NonceStr,
		"out_trade_no": outTradeNo,
	}
	data.Sign = createSign(dataMap, c.MchKey)
	xmlData, err := xml.Marshal(data)
	if err != nil {
		return nil, err
	}
	url := closeOrderUrl
	if c.SandBox {
		url = sandboxCloseOrderUrl
	}
	body, err := tools.Post(url, xmlData)
	if err != nil {
		return nil, err
	}
	response := new(orderCloseResponse)
	err = xml.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}
	if response.ReturnCode != "SUCCESS" {
		return response, errors.New(response.ReturnMsg)
	}
	return response, err
}

//申请退款
//详细规则参考 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_4
func (c *WxClient) Refund(req RefundRequest) (*refundResponse, error) {
	if req.OutTradeNo == "" && req.TransactionID == "" {
		return nil, outTradeNoAndTransactionIdIsNull
	}
	if req.OutRefundNo == "" {
		return nil, outRefundNoIsNull
	}
	if req.TotalFee == 0 {
		return nil, totalFeeIsNull
	}
	if req.RefundFee == 0 {
		return nil, refundFeeIsNull
	}
	data := refundData{
		AppID:         c.AppID,
		MchID:         c.MchID,
		NonceStr:      tools.GenerateNonceStr(),
		TransactionID: req.TransactionID,
		OutTradeNo:    req.OutTradeNo,
		OutRefundNo:   req.OutRefundNo,
		TotalFee:      req.TotalFee,
		RefundFee:     req.RefundFee,
		NotifyUrl:     req.NotifyUrl,
	}
	dataMap := map[string]string{
		"appid":      c.AppID,
		"mch_id":     c.MchID,
		"nonce_str":  data.NonceStr,
		"total_fee":  strconv.Itoa(data.TotalFee),
		"refund_fee": strconv.Itoa(data.RefundFee),
	}
	if req.TransactionID != "" {
		dataMap["transaction_id"] = req.TransactionID
	} else {
		dataMap["out_trade_no"] = req.OutTradeNo
	}
	if req.NotifyUrl != "" {
		dataMap["notify_url"] = req.NotifyUrl
	}
	data.Sign = createSign(dataMap, c.MchKey)
	xmlData, err := xml.Marshal(data)
	if err != nil {
		return nil, err
	}
	url := refundUrl
	if c.SandBox {
		url = sandboxRefundUrl
	}
	body, err := tools.PostWithCert(url, c.MchID, xmlData, req.CertPath)
	if err != nil {
		return nil, err
	}
	response := new(refundResponse)
	err = xml.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}
	if response.ReturnCode != "SUCCESS" {
		return response, errors.New(response.ReturnMsg)
	}
	return response, err
}

//申请退款
//详细规则参考 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_4
func (c *WxClient) RefundQuery(req RefundQueryRequest) (map[string]string, error) {
	if req.OutRefundNo == "" && req.TransactionID == "" && req.OutTradeNo == "" && req.RefundID == "" {
		return nil, AllParamsIsNull
	}
	data := refundQueryData{
		AppID:         c.AppID,
		MchID:         c.MchID,
		NonceStr:      tools.GenerateNonceStr(),
		TransactionID: req.TransactionID,
		OutTradeNo:    req.OutTradeNo,
		OutRefundNo:   req.OutRefundNo,
		RefundID:      req.RefundID,
	}
	dataMap := map[string]string{
		"appid":     c.AppID,
		"mch_id":    c.MchID,
		"nonce_str": data.NonceStr,
	}
	if req.RefundID != "" {
		dataMap["refund_id"] = req.RefundID
	}else if req.OutRefundNo != "" {
		dataMap["out_refund_no"] = req.RefundID
	} else if req.TransactionID != "" {
		dataMap["transaction_id"] = req.TransactionID
	} else {
		dataMap["out_trade_no"] = req.OutTradeNo
	}
	data.Sign = createSign(dataMap, c.MchKey)
	xmlData, err := xml.Marshal(data)
	if err != nil {
		return nil, err
	}
	url := refundQueryUrl
	if c.SandBox {
		url = sandboxRefundQueryUrl
	}
	body, err := tools.Post(url, xmlData)
	if err != nil {
		return nil, err
	}
	responseMap := XmlToMap(body)
	if responseMap["return_code"] != "SUCCESS" {
		return responseMap, errors.New(responseMap["return_msg"])
	}
	return responseMap, nil
}
