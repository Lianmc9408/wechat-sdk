package pay

import "encoding/xml"

type WxClient struct {
	AppID   string
	MchID   string
	MchKey  string
	SandBox bool
}

// ====统一下单====
type unifiedOrderData struct {
	XMLName        xml.Name `xml:"xml"`
	AppID          string   `xml:"appid"`
	MchID          string   `xml:"mch_id"`
	NotifyUrl      string   `xml:"notify_url"`
	NonceStr       string   `xml:"nonce_str"`
	Sign           string   `xml:"sign"`
	Body           string   `xml:"body"`
	OutTradeNo     string   `xml:"out_trade_no"`
	TotalFee       uint32   `xml:"total_fee"`
	SpbillCreateIP string   `xml:"spbill_create_ip"`
	TradeType      string   `xml:"trade_type"`
	OpenID         string   `xml:"openid,omitempty"`
	ProductId      string   `xml:"product_id,omitempty"`
	Attach         string   `xml:"attach,omitempty"`
}

type unifiedOrderResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	AppID      string `xml:"appid"`
	MchID      string `xml:"mch_id"`
	DeviceInfo string `xml:"device_info"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
	TradeType  string `xml:"trade_type"`
	PrepayID   string `xml:"prepay_id"`
	CodeUrl    string `xml:"code_url"`
}

// ====H5和小程序支付====
type JsApiRequest struct {
	Body           string
	OutTradeNo     string
	TotalFee       uint32
	SpbillCreateIP string
	OpenID         string
	Attach         string
	NotifyUrl      string
}

type jsApiResponse struct {
	AppID     string `json:"appId,omitempty"`
	Timestamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

// Native支付
type NativeRequest struct {
	Body           string
	OutTradeNo     string
	TotalFee       uint32
	SpbillCreateIP string
	ProductId      string
	Attach         string
}

// ====App支付====
type AppRequest struct {
	Body           string
	OutTradeNo     string
	TotalFee       uint32
	SpbillCreateIP string
	Attach         string
	NotifyUrl      string
}

type appResponse struct {
	AppID     string `json:"appId"`
	PartnerID string `json:"partnerId"`
	PrepayID  string `json:"prepayId"`
	Package   string `json:"package"`
	NonceStr  string `json:"nonceStr"`
	Timestamp string `json:"timeStamp"`
	Sign      string `json:"sign"`
}

// ====订单查询====
type OrderQueryRequest struct {
	TransactionID string
	OutTradeNo    string
}

type orderQueryData struct {
	XMLName       xml.Name `xml:"xml"`
	AppID         string   `xml:"appid"`
	MchID         string   `xml:"mch_id"`
	TransactionID string   `xml:"transaction_id,omitempty"`
	OutTradeNo    string   `xml:"out_trade_no,omitempty"`
	NonceStr      string   `xml:"nonce_str"`
	Sign          string   `xml:"sign"`
}

type orderQueryResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	AppID      string `xml:"appid"`
	MchID      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
	// 以下字段在return_code,result_code,trade_state都为SUCCESS时有返回
	// 如trade_state不为SUCCESS,则只返回out_trade_no(必传)和attach(选传)
	DeviceInfo     string `xml:"device_info"`
	OpenID         string `xml:"openid"`
	IsSubscribe    string `xml:"is_subscribe"`
	TradeType      string `xml:"trade_type"`
	TradeState     string `xml:"trade_state"`
	BankType       string `xml:"bank_type"`
	TotalFee       uint32 `xml:"total_fee"`
	CashFee        uint32 `xml:"cash_fee"`
	TransactionID  string `xml:"transaction_id"`
	OutTradeNo     string `xml:"out_trade_no"`
	Attach         string `xml:"attach"`
	TimeEnd        string `xml:"time_end"`
	TradeStateDesc string `xml:"trade_state_desc"`
}

//====充值回调====
type CallBackBaseRequest struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
}

type CallBackRequest struct {
	ReturnCode    string `xml:"return_code"`
	ReturnMsg     string `xml:"return_msg"`
	AppID         string `xml:"appid"`
	MchID         string `xml:"mch_id"`
	NonceStr      string `xml:"nonce_str"`
	Sign          string `xml:"sign"`
	ResultCode    string `xml:"result_code"`
	OpenID        string `xml:"openid"`
	IsSubscribe   string `xml:"is_subscribe"`
	TradeType     string `xml:"trade_type"`
	BankType      string `xml:"bank_type"`
	TotalFee      uint32 `xml:"total_fee"`
	CashFee       uint32 `xml:"cash_fee"`
	TransactionID string `xml:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	Attach        string `xml:"attach,omitempty"`
	TimeEnd       string `xml:"time_end"`
	FeeType       string `xml:"fee_type"`
}

// ====关闭订单====
type orderCloseData struct {
	XMLName    xml.Name `xml:"xml"`
	AppID      string   `xml:"appid"`
	MchID      string   `xml:"mch_id"`
	OutTradeNo string   `xml:"out_trade_no"`
	NonceStr   string   `xml:"nonce_str"`
	Sign       string   `xml:"sign"`
}

type orderCloseResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	AppID      string `xml:"appid"`
	MchID      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
	// 以下字段在return_code,result_code,trade_state都为SUCCESS时有返回
	// 如trade_state不为SUCCESS,则只返回out_trade_no(必传)和attach(选传)
	DeviceInfo     string `xml:"device_info"`
	OpenID         string `xml:"openid"`
	IsSubscribe    string `xml:"is_subscribe"`
	TradeType      string `xml:"trade_type"`
	TradeState     string `xml:"trade_state"`
	BankType       string `xml:"bank_type"`
	TotalFee       uint32 `xml:"total_fee"`
	CashFee        uint32 `xml:"cash_fee"`
	TransactionID  string `xml:"transaction_id"`
	OutTradeNo     string `xml:"out_trade_no"`
	Attach         string `xml:"attach"`
	TimeEnd        string `xml:"time_end"`
	TradeStateDesc string `xml:"trade_state_desc"`
}

//====申请退款====
type RefundRequest struct {
	TransactionID string
	OutTradeNo    string
	OutRefundNo   string
	TotalFee      int
	RefundFee     int
	NotifyUrl     string
	CertPath      string
}

type refundData struct {
	XMLName       xml.Name `xml:"xml"`
	AppID         string   `xml:"appid"`
	MchID         string   `xml:"mch_id"`
	NonceStr      string   `xml:"nonce_str"`
	Sign          string   `xml:"sign"`
	TransactionID string   `xml:"transaction_id,omitempty"`
	OutTradeNo    string   `xml:"out_trade_no,omitempty"`
	OutRefundNo   string   `xml:"out_refund_no"`
	TotalFee      int      `xml:"total_fee"`
	RefundFee     int      `xml:"refund_fee"`
	NotifyUrl     string   `xml:"notify_url,omitempty"`
}

type refundResponse struct {
	ReturnCode    string `xml:"return_code"`
	ReturnMsg     string `xml:"return_msg"`
	ResultCode    string `xml:"result_code"`
	AppID         string `xml:"appid"`
	MchID         string `xml:"mch_id"`
	NonceStr      string `xml:"nonce_str"`
	Sign          string `xml:"sign"`
	TransactionID string `xml:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	OutRefundNo   string `xml:"out_refund_no"`
	RefundID      string `xml:"refund_id"`
	RefundFee     int    `xml:"refund_fee"`
	TotalFee      uint32 `xml:"total_fee"`
	CashFee       uint32 `xml:"cash_fee"`
	ErrCode       string `xml:"err_code"`
	ErrCodeDes    string `xml:"err_code_des"`
}

//====退款查询====
type RefundQueryRequest struct {
	TransactionID string
	OutTradeNo    string
	OutRefundNo   string
	RefundID      string
}

type refundQueryData struct {
	XMLName       xml.Name `xml:"xml"`
	AppID         string   `xml:"appid"`
	MchID         string   `xml:"mch_id"`
	NonceStr      string   `xml:"nonce_str"`
	Sign          string   `xml:"sign"`
	TransactionID string   `xml:"transaction_id,omitempty"`
	OutTradeNo    string   `xml:"out_trade_no,omitempty"`
	OutRefundNo   string   `xml:"out_refund_no,omitempty"`
	RefundID      string   `xml:"refund_id,omitempty"`
}

//type refundQueryResponse struct {
//	ReturnCode    string `xml:"return_code"`
//	ReturnMsg     string `xml:"return_msg"`
//	ResultCode    string `xml:"result_code"`
//	ErrCode       string `xml:"err_code"`
//	ErrCodeDes    string `xml:"err_code_des"`
//	AppID         string `xml:"appid"`
//	MchID         string `xml:"mch_id"`
//	NonceStr      string `xml:"nonce_str"`
//	Sign          string `xml:"sign"`
//	TransactionID string `xml:"transaction_id"`
//	OutTradeNo    string `xml:"out_trade_no"`
//	TotalFee      uint32 `xml:"total_fee"`
//	CashFee       uint32 `xml:"cash_fee"`
//	RefundCount   int    `xml:"refund_count"`
//	// TODO 部分数据因为键无法变成属性， 待解决
//}

