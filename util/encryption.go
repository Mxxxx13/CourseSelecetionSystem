// @Title : encryption
// @Description :加密密码
// @Author : MX
// @Update : 2022/4/21 20:10

package util

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func EncryptionPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("加密失败")
	}
	password = string(hashPassword)
	return password, err
}
