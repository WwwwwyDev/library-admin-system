package util

import (
	"crypto/md5"
	"fmt"
)

func Str2Md5(str string) string{
	strb := []byte(str)
	has := md5.Sum(strb)
	strMd5 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return strMd5
}
