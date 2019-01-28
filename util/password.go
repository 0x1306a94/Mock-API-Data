package util

import (
	"crypto/md5"
	"encoding/hex"
)

// 加密用户密码
func EncryptionPassword(password string) string {
	ctx := md5.New()
	ctx.Write([]byte(password))
	encryptionPassword := hex.EncodeToString(ctx.Sum(nil))
	return encryptionPassword
}

// 验证用户密码
func ValidationPassword(password, encryption string) bool {
	ctx := md5.New()
	ctx.Write([]byte(password))
	encryptionPassword := hex.EncodeToString(ctx.Sum(nil))
	return encryptionPassword == encryption
}
