package dada

import (
	"encoding/hex"
	"crypto/md5"
)

func MD5(str string) (res string) {
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串为 str
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr) // 输出加密结果
}
