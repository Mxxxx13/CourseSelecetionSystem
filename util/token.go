// @Title : token
// @Description :生成jwt
// @Author : MX
// @Update : 2022/4/15 22:22

package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"
)

type Header struct {
	Alg string `json:"alg"` // 加密算法
	Typ string `json:"typ"` // 类别
}

type Sub struct {
	Uid      uint   `json:"id"`
	Username string `json:"username"` // 用户名
	Role     string `json:"role"`     // 身份
}

type Payload struct {
	Iss string `json:"iss"` // 发布者
	Exp int64  `json:"exp"` // 过期时间
	Iat int64  `json:"iat"` // 颁发时间
	Sub
}

type JWT struct {
	Header    Header
	Payload   Payload
	Signature string
	Token     string
}

var (
	key = "mx" // 密钥
)

// NewJWT 创建jwt
func NewJWT(id uint, username string, role string) (jwt JWT) {
	// 初始化header
	jwt.Header = Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	// 初始化Payload
	jwt.Payload = Payload{
		Iss: "mx",
		Exp: time.Now().Add(7 * 24 * time.Hour).Unix(),
		Iat: time.Now().Unix(),
		Sub: Sub{
			Uid:      id,
			Username: username,
			Role:     role,
		},
	}

	// 将header和payload序列化并编码
	h, err := json.Marshal(jwt.Header)
	if err != nil {
		log.Println("NewJWT:marshal header:", err)
	}
	baseH := base64.StdEncoding.EncodeToString(h)
	p, err := json.Marshal(jwt.Payload)
	if err != nil {
		log.Println("NewJWT:marshal payload:", err)
	}
	baseP := base64.StdEncoding.EncodeToString(p)

	source := baseH + "." + baseP

	// 将Signature加密编码
	s := Encryption(source, key)
	jwt.Signature = base64.StdEncoding.EncodeToString(s)
	jwt.Token = source + "." + jwt.Signature
	return jwt
}

// Encryption 加密
func Encryption(source string, key string) []byte {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(source))
	return mac.Sum(nil)
}

// CheckJWT 验证token的合法性并更新jwt
func CheckJWT(token string) (jwt JWT, err error) {
	err = errors.New("token error")
	arr := strings.Split(token, ".") // 将header, payload, Signature 分开
	if len(arr) < 3 {
		log.Println("CheckJWT:len < 3")
		return
	}
	// Header解码
	baseH := arr[0]
	h, err := base64.StdEncoding.DecodeString(baseH)
	if err != nil {
		log.Println("CheckJWT:decode header:", err)
		return
	}

	// Header反序列化
	err = json.Unmarshal(h, &jwt.Header)
	if err != nil {
		log.Println("CheckJWT:unmarshal header:", err)
		return
	}

	// Payload解码
	baseP := arr[1]
	p, err := base64.StdEncoding.DecodeString(baseP)
	if err != nil {
		log.Println("CheckJWT:decode payload:", err)
		return
	}

	// Payload反序列化
	err = json.Unmarshal(p, &jwt.Payload)
	if err != nil {
		log.Println("CheckJWT:unmarshal payload:", err)
		return
	}

	// Signature解码
	baseS := arr[2]
	s1, err := base64.StdEncoding.DecodeString(baseS)
	if err != nil {
		log.Println("decode bases", err)
		return
	}

	// 加密source后与signature比较验证token合法性
	source := baseH + "." + baseP

	s2 := Encryption(source, key)
	if string(s1) != string(s2) {
		log.Println("token is illegal")
		return
	} else {
		jwt.Signature = arr[2]
		jwt.Token = token
	}

	// 验证token是否过期
	if jwt.Payload.Exp < time.Now().Unix() {
		log.Println("CheckJWT:token has been expired")
		err = errors.New("CheckJWT:token has been expired")
		return
	}
	return jwt, nil
}
