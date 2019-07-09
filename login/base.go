package login

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	redirectOauthURL      = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect"
	accessTokenURL        = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	globalAccessTokenURL  = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	refreshAccessTokenURL = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s"
	userInfoURL           = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"
	checkAccessTokenURL   = "https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s"

	jsApiTicketURL = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"

	mediaGetUrl = "https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s"
)

func CreateJsApiConfigSign(nonceStr, ticket, url string, timestamp int64) string {
	dataMap := map[string]string{
		"noncestr":     nonceStr,
		"jsapi_ticket": ticket,
		"timestamp":    strconv.Itoa(int(timestamp)),
		"url":          url,
	}
	signList := make([]string, 0)
	for k, v := range dataMap {
		if v != "" {
			signList = append(signList, fmt.Sprintf("%s=%s", k, v))
		}
	}
	sort.Strings(signList)
	signStr := strings.Join(signList, "&")

	h := sha1.New()
	h.Write([]byte(signStr))
	return hex.EncodeToString(h.Sum(nil))
}
