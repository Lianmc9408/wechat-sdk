package pay

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"sort"
	"strings"
)

const (
	microPayUrl                = "https://api.mch.weixin.qq.com/pay/micropay"
	unifiedOrderUrl            = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	orderQueryUrl              = "https://api.mch.weixin.qq.com/pay/orderquery"
	reverseUrl                 = "https://api.mch.weixin.qq.com/secapi/pay/reverse"
	closeOrderUrl              = "https://api.mch.weixin.qq.com/pay/closeorder"
	refundUrl                  = "https://api.mch.weixin.qq.com/secapi/pay/refund"
	refundQueryUrl             = "https://api.mch.weixin.qq.com/pay/refundquery"
	downloadBillUrl            = "https://api.mch.weixin.qq.com/pay/downloadbill"
	downloadFundFlowUrl        = "https://api.mch.weixin.qq.com/pay/downloadfundflow"
	reportUrl                  = "https://api.mch.weixin.qq.com/payitil/report"
	shortUrl                   = "https://api.mch.weixin.qq.com/tools/shorturl"
	authCodeToOpenidUrl        = "https://api.mch.weixin.qq.com/tools/authcodetoopenid"
	sandboxMicroPayUrl         = "https://api.mch.weixin.qq.com/sandboxnew/pay/micropay"
	sandboxUnifiedOrderUrl     = "https://api.mch.weixin.qq.com/sandboxnew/pay/unifiedorder"
	sandboxOrderQueryUrl       = "https://api.mch.weixin.qq.com/sandboxnew/pay/orderquery"
	sandboxReverseUrl          = "https://api.mch.weixin.qq.com/sandboxnew/secapi/pay/reverse"
	sandboxCloseOrderUrl       = "https://api.mch.weixin.qq.com/sandboxnew/pay/closeorder"
	sandboxRefundUrl           = "https://api.mch.weixin.qq.com/sandboxnew/secapi/pay/refund"
	sandboxRefundQueryUrl      = "https://api.mch.weixin.qq.com/sandboxnew/pay/refundquery"
	sandboxDownloadBillUrl     = "https://api.mch.weixin.qq.com/sandboxnew/pay/downloadbill"
	sandboxDownloadFundFlowUrl = "https://api.mch.weixin.qq.com/sandboxnew/pay/downloadfundflow"
	sandboxReportUrl           = "https://api.mch.weixin.qq.com/sandboxnew/payitil/report"
	sandboxShortUrl            = "https://api.mch.weixin.qq.com/sandboxnew/tools/shorturl"
	sandboxAuthCodeToOpenidUrl = "https://api.mch.weixin.qq.com/sandboxnew/tools/authcodetoopenid"
)

var (
	outTradeNoIsNull                  = errors.New("缺少接口必填参数out_trade_no")
	bodyIsNull                        = errors.New("缺少接口必填参数body")
	totalFeeIsNull                    = errors.New("缺少接口必填参数total_fee")
	tradeTypeIsNull                   = errors.New("缺少接口必填参数trade_type")
	spbillCreateIPIsNull              = errors.New("缺少接口必填参数spbill_create_ip")
	tradeTypeJSAPIButOpenidIsNull     = errors.New("trade_type为JSAPI时, openid为必填参数")
	tradeTypeNativeButProductIdIsNull = errors.New("trade_type为NATIVE时, product_id为必填参数")
	outTradeNoAndTransactionIdIsNull  = errors.New("out_trade_no, transaction_id至少填一个")
	AllParamsIsNull                   = errors.New("out_trade_no, transaction_id, out_refund_no, refund_id至少填一个")
	outRefundNoIsNull                 = errors.New("缺少接口必填参数out_refund_no")
	refundFeeIsNull                   = errors.New("缺少接口必填参数refund_fee")
)

func createSign(dataMap map[string]string, mchKey string) string {
	signList := make([]string, 0)
	for k, v := range dataMap {
		if v != "" {
			signList = append(signList, fmt.Sprintf("%s=%s", k, v))
		}
	}
	sort.Strings(signList)
	signStr := strings.Join(signList, "&")
	signStr = signStr + "&key=" + mchKey

	c := md5.New()
	c.Write([]byte(signStr))
	return strings.ToUpper(hex.EncodeToString(c.Sum(nil)))
}
