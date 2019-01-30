package tools

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

func GenerateNonceStr() string {
	ctx := md5.New()
	ctx.Write([]byte(strconv.Itoa(int(time.Now().UnixNano()))))
	return hex.EncodeToString(ctx.Sum(nil))
}
