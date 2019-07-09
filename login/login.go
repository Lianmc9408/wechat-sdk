package login

import (
	"encoding/json"
	"fmt"
	"wechat-sdk/tools"
)

func (c *WxClient) GetAccessToken(code string) (aT *accessToken, err error) {
	resp, err := tools.Get(fmt.Sprintf(accessTokenURL, c.AppID, c.Secret, code))
	if err != nil {
		return nil, err
	}
	aT = new(accessToken)
	if err := json.Unmarshal(resp, aT); err != nil {
		return nil, err
	}
	if aT.ErrCode != 0 {
		return nil, fmt.Errorf("ErrCode:[%d], ErrMsg:[%s]", aT.ErrCode, aT.ErrMsg)
	}
	return
}

func (c *WxClient)GetGlobalAccessToken()(aT *globalAccessToken, err error){
	resp, err := tools.Get(fmt.Sprintf(globalAccessTokenURL, c.AppID, c.Secret))
	if err != nil {
		return nil, err
	}
	aT = new(globalAccessToken)
	if err := json.Unmarshal(resp, aT); err != nil {
		return nil, err
	}
	if aT.ErrCode != 0 {
		return nil, fmt.Errorf("ErrCode:[%d], ErrMsg:[%s]", aT.ErrCode, aT.ErrMsg)
	}
	return
}

func (c *WxClient)GetJsApiTicket(accessToken string)(ticket *jsApiTicket, err error){
	resp, err := tools.Get(fmt.Sprintf(jsApiTicketURL, accessToken))
	if err != nil {
		return nil, err
	}
	ticket = new(jsApiTicket)
	if err := json.Unmarshal(resp, ticket); err != nil {
		return nil, err
	}
	if ticket.ErrCode != 0 {
		return nil, fmt.Errorf("ErrCode:[%d], ErrMsg:[%s]", ticket.ErrCode, ticket.ErrMsg)
	}
	return
}