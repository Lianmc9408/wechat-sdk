package login

import (
	"encoding/json"
	"fmt"
	"github.com/yeezyi/wechat-sdk/tools"
)

func GetMedia(accessToken, mediaId string)(media []byte, err error){
	resp, err := tools.Get(fmt.Sprintf(mediaGetUrl, accessToken, mediaId))
	if err != nil {
		return nil, err
	}
	ret := new(Err)
	if err := json.Unmarshal(resp, ret); err == nil {
		return nil, fmt.Errorf("ErrCode:[%d], ErrMsg:[%s]", ret.ErrCode, ret.ErrMsg)
	}
	return resp, nil
}
