package util

import (
	"crypto/md5"
	"encoding/hex"
)

// md5文件名加密
func EncodingMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
