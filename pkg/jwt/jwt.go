/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-12-22 13:19:04
 */
package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

var secret string = "12345678"

func Generate(alg string, payload Payload) (string, error) {

	var (
		jtoken      string
		header      Header
		headerJson  string
		payloadJson string
		err         error
		sign        string
		builder     strings.Builder
	)

	header.Alg = alg
	header.Typ = "JWT"
	headerJson, err = header.Gen()
	if err != nil {
		return jtoken, err
	}

	payloadJson, err = payload.Gen()
	if err != nil {
		return jtoken, err
	}

	sign, err = Signature(headerJson, payloadJson, header.Alg)
	if err != nil {
		fmt.Println(err)
	}

	builder.WriteString(headerJson)
	builder.WriteString(".")
	builder.WriteString(payloadJson)
	builder.WriteString(".")
	builder.WriteString(sign)

	jtoken = builder.String()

	return jtoken, err
}

/**
* 校验jtoken
 */
func Check(jtoken string) (payload Payload, err error) {
	var (
		secs     []string
		header   Header
		jsonByte []byte
		sign     string
	)
	secs = strings.Split(jtoken, ".")

	//解析header
	jsonByte, err = base64.StdEncoding.DecodeString(secs[0])
	if err != nil {
		return payload, err
	}

	err = json.Unmarshal(jsonByte, &header)
	if err != nil {
		return payload, err
	}

	//验证签名
	sign, err = Signature(secs[0], secs[1], header.Alg)
	if err != nil {
		return payload, err
	}

	if sign != secs[2] {
		return payload, errors.New("jtoken签名验证失败")
	}

	//解析payload
	jsonByte, err = base64.StdEncoding.DecodeString(secs[1])
	if err != nil {
		return payload, err
	}

	err = json.Unmarshal(jsonByte, &payload)
	if err != nil {
		return payload, err
	}

	//校验是否过期
	if time.Until(payload.Exp).Minutes() < 0 {
		return payload, errors.New("token已失效")
	}

	return payload, err
}
