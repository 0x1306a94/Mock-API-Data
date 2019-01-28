package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func NewProjectKey(uid int64) string {
	str := fmt.Sprintf("%v+abcde+%v+adcde", uid, time.Now().Unix())
	ctx := md5.New()
	ctx.Write([]byte(str))
	return hex.EncodeToString(ctx.Sum(nil))
}
