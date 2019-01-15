package outputformat

import (
	"crypto/md5"
	"fmt"

	"http_sql_api/config"
)
func Md5(str string,key ...string) string {
	key_val := config.AppConfig.AppKey
	if len(key) > 0 {
		key_val = key[0]
	}
	str = str + key_val
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str1
}