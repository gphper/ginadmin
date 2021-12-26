/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-12-22 13:55:23
 */
package jwt

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {

	var payload Payload

	payload.Name = "gphper"
	payload.Uid = 1
	payload.Exp = time.Now().Local().Add(5 * time.Minute)

	jtoken, _ := Generate("HS256", payload)
	fmt.Println(jtoken)
}

func TestCheck(t *testing.T) {
	payload, err := Check("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHAiOiIyMDIxLTEyLTIyVDE1OjE0OjQ4LjE3NTM2NjkrMDg6MDAiLCJOYW1lIjoiZ3BocGVyIiwiVWlkIjoxfQ==.516ed47243cdd47231db838d9e7d58345b49c0da3f681450e5ed4116e9a1ce50")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(payload)
}
