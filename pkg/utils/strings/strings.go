/*
 * @Description:
 * @Author: gphper
 * @Date: 2022-03-27 10:58:44
 */
package strings

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

/**
生成随机字符串
*/
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

/**
密码加密
*/
func Encryption(password string, salt string) string {
	str := fmt.Sprintf("%s%s", password, salt)
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

/**
*首字母大写
**/
func StrFirstToUpper(str string) (string, string) {
	temp := strings.Split(str, "_")
	var upperStr string
	var firstStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				firstStr += string(vv[i])
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				upperStr += string(vv[i])
			}
		}
	}
	return upperStr, firstStr
}

/*
*比较第二个slice一第一个slice的区别
 */
func CompareSlice(first []string, second []string) (add []string, incre []string) {

	secondMap := make(map[string]struct{})

	for _, v := range second {
		secondMap[v] = struct{}{}
	}

	for _, v := range first {
		_, ok := secondMap[v]
		if !ok {
			incre = append(incre, v)
		} else {
			delete(secondMap, v)
		}
	}

	for k, _ := range secondMap {
		add = append(add, k)
	}

	return
}

/**
* 组装字符串
 */
func JoinStr(items ...interface{}) string {
	if len(items) == 0 {
		return ""
	}
	var builder strings.Builder
	for _, v := range items {
		builder.WriteString(v.(string))
	}
	return builder.String()
}
