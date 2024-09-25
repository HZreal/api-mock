package service

/**
 * @Author elastic·H
 * @Date 2024-09-23
 * @File: requestTaskService.go
 * @Description:
 */

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
)

// TODO 从登录认证获取认证信息
func authenticate() string {
	//
	return "1TCATNOC1REMOTSUC%7CYPyZYhGVtuj5q1MrFHkk%2BhZCDJh%2FWYeIQS8xBdTvBO%2BAnnHgjEj%2FiBTNH6g3RVgZjlmRMZCNbK7EwW6K%2Bd709atJ3r8Zwif3jCMnhkGnJUqG%2FFmbT0zNb%2FZMVvSdYHoqMdJcJcRxIW6Q%2FLU3oi14JJ6Q3vRzC7AHDMlqJByjoJziGqtxkJ0ZI1w0CNliTnq4t67PEmutJKr6z68%2Bj6RgkTDvJ%2BkGhefQvSjqn5YE6%2BhJYgwsTx8EcSAQOULeXBY1G1c5%2Bj901rETK6F%2B3uc2p0UTFhemf2WaJz412FRZC8U9sr%2BEDc0dBTkbgMeom0zo6FTH%2FlyRDIRMOX8q7rnTfMoJY1YRm4YVhxXfb0hqKz8UNWD9q1DMKIkBJ67DsV1OLYRlCLx419Ck4ej7z30Zgz80HmK0lnDjGeE%2BZMfxT%2BgNVNo0fxMcuofmSnZWPAKKS90GE1Ul9u%2FD6C9CNofhfCemVvAFpngrjxKNnEhTTLlJ6g0mW4tMFfc0r0wKkhBhMlKGHhVB4luUruRp3KlMDRPAP8CjebrQWzQ2Wbn2q2ZDAqMFS95MChvbKkHAF%2B8WY0dNETw6K45mQshH4bCB%2Be92CJnbGtvMwE2fwjXNZwHfl9aLNm5J7PSSxXMAl5O5lt0jVfBEIvJnJfhi9AsjM366WISjJHuyJo01VS4spJYJfhOW0cGyGpWsb04s0OITaTGaLgqKoa3fnaJYsFG67Q%3D%3D"
}

func getHost() string {
	return "https://172.16.28.93"
}

func addCookies(req *http.Request) {

}

// 发送请求
func sendReq(url string, method string, body string) {
	//
	req, err := http.NewRequest(method, url, bytes.NewBufferString(body))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// Content-Type
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Cookie
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: cookieJar}
	cookie := &http.Cookie{
		Name:  "srv_session_id",
		Value: authenticate(),
	}
	req.AddCookie(cookie)

	//
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}
	fmt.Println("响应内容:", string(bodyBytes))

}
