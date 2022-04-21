// @Title : code
// @Description :生成验证码
// @Author : MX
// @Update : 2022/4/21 20:02

package util

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateCode 生成六位验证码
func GenerateCode() string {
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	return code
}
