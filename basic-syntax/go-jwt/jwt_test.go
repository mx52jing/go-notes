package main

import (
	"fmt"
	"testing"
	"time"
)

func TestJWTToken(t *testing.T) {
	header := DefaultHeader
	payload := JwtPayload{
		ID:         "9527",
		Issuer:     "凌凌漆",
		Audience:   "ALL",
		Subject:    "找回国宝",
		IssuedAt:   time.Now().Unix(),
		Expiration: time.Now().Add(time.Second * 10).Unix(),
		UserDefined: map[string]any{
			"Name": "血灵",
			"Age":  22,
			"Sex":  "male",
		},
	}
	token, err := generateToken(header, payload)
	if err != nil {
		fmt.Printf("生成json web token失败: %v\n", err)
		return
	}
	fmt.Printf("生成Token成功，生a成的Token为【%s】\n", token)
	_, payloadRes, err := verifyToken(token)
	if err != nil {
		fmt.Printf("Token验证失败%s", err)
		return
	}
	fmt.Printf("Token验证成功，得到的payload部分为\n%v\n", payloadRes)
}
