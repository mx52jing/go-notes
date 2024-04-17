package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"strings"
)

//JWT: Json Web Token

// jwt header
type JwtHeader struct {
	Algo string `json:"alg"` // 哈希算法，默认为HMAC SHA256(写为 HS256)
	Type string `json:"typ"` // 令牌(token)类型，统一写为JWT
}

// jwt payload
type JwtPayload struct {
	ID          string         `json:"jti"` //(JWT ID)编号，用于标识该JWT
	Issuer      string         `json:"iss"` // issuer签发人，比如微信
	Audience    string         `json:"aud"` // 受众人,比如英雄联盟
	Subject     string         `json:"sub"` // 主题
	IssuedAt    int64          `json:"iat"` //发布时间,精确到秒
	NotBefore   int64          `json:"nbf"` //生效时间，在此之前不可用,精确到秒
	Expiration  int64          `json:"exp"` //过期时间，精确到秒
	UserDefined map[string]any `json:"ud"`  //用户自定义的其他字段
}

var DefaultHeader = JwtHeader{
	Algo: "HS256",
	Type: "JWT",
}

const JWT_SECRET = "aa&^#*bjbH)(U)72_+894BHFA"

// Signature 部分是对前两部分(Header，Payload)的签名，防止数据篡改
// 需要指定一个密钥（secret），这个密钥只有服务器才知道，不能泄露给用户
// 然后，使用Header里面指定的签名算法（默认是 HMAC SHA256），按照下面的公式产生签名
// HMACSHA256(
// base64UrlEncode(header) + "." +
// base64UrlEncode(payload),
// secret
// )

// 生成Token原理
func generateToken(header JwtHeader, payload JwtPayload) (string, error) {
	var headerPart, payloadPart, signature string
	headerByte, err := sonic.Marshal(header)
	if err != nil {
		return "", err
	}
	// 得到第一部分 header 部分
	headerPart = base64.RawURLEncoding.EncodeToString(headerByte)
	payloadByte, err := sonic.Marshal(payload)
	if err != nil {
		return "", err
	}
	// 得到第二部分 payload 部分
	payloadPart = base64.RawURLEncoding.EncodeToString(payloadByte)
	// 通过密钥secret 生成jwt
	//基于sha256的哈希认证算法。任意长度的字符串，经过sha256之后长度都变成了256 bits
	h := hmac.New(sha256.New, []byte(JWT_SECRET))
	h.Write([]byte(fmt.Sprintf("%s.%s", headerPart, payloadPart)))
	signature = base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	return fmt.Sprintf("%s.%s.%s", headerPart, payloadPart, signature), nil
}

// 校验Token是否合法
func verifyToken(token string) (*JwtHeader, *JwtPayload, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, nil, errors.New("token 不合法\n")
	}
	headerStr := parts[0]
	payloadStr := parts[1]
	// 进行哈希签名的验证
	h := hmac.New(sha256.New, []byte(JWT_SECRET))
	h.Write([]byte(fmt.Sprintf("%s.%s", headerStr, payloadStr)))
	signature := base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	if signature != parts[2] {
		return nil, nil, errors.New("token 校验失败\n")
	}
	var headerPart, payloadPart []byte
	var err error
	headerPart, err = base64.RawURLEncoding.DecodeString(headerStr)
	if err != nil {
		return nil, nil, errors.New("header部分 base64反解失败\n")
	}
	payloadPart, err = base64.RawURLEncoding.DecodeString(payloadStr)
	if err != nil {
		return nil, nil, errors.New("payload部分 base64反解失败\n")
	}
	var header JwtHeader
	var payload JwtPayload
	err = sonic.Unmarshal(headerPart, &header)
	if err != nil {
		return nil, nil, errors.New("header部分 json反序列化失败\n")
	}
	err = sonic.Unmarshal(payloadPart, &payload)
	if err != nil {
		return nil, nil, errors.New("header部分 json反序列化失败\n")
	}
	return &header, &payload, nil
}
