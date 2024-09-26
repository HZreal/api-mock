package service

/**
 * @Author elastic·H
 * @Date 2024-09-23
 * @File: requestTaskService.go
 * @Description:
 */

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"io"
	"net/http"
	"strings"
)

var (
	HostBase                 = "https://172.16.28.93"
	GlobalSrvSessionIdCookie *http.Cookie
	GlobalPHPSESSIONCookie   *http.Cookie
	client                   *http.Client
)

func init() {
	//
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client = &http.Client{Transport: transport}
}

func Login() {
	method := "POST"
	body := "userName=admin&password=e4aed93ec8e0084edaef0cb945aa5acb885792dea7c115f6de9a96c77df0ca617761738c76e965f18aafc30eccbe0dacc9ec1788a7a1bbc0d6ef59b98047c099&rememberMe=true&yzm=&customerName=dipcc&client=1920*1080"
	uri := "/service/index.php?m=login&c=login&f=login"

	//
	req, err := http.NewRequest(method, HostBase+uri, strings.NewReader(body))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	//
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	// set Content-Type
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// 登录可以不用 cookie
	// req.Header.Add("Cookie", "srv_session_id=TCATNOCREMOTSUC%7CGtnwnxvTgKW4aY%2FBH%2Fe18FywA9IhxUSdmgMEos0ftdFB7fGK0hNG1gBVOPECb1CggV3VOpj1cmPL6fTgYOmHvmIAvkbqrGnXNay1s1tAMH9H5PCHaIgZXDjURDf5eDv%2F1pGrBmdHyboVD%2FkJiVi5NEvg8NzA%2FkISPgP3RsMim%2BwYU8mKm0S0MeArJw9ddxjBwnn1yBCxT3cUXTkHADxmkQ%3D%3D; PHPSESSID=mv4r9ro922eurj0r978tbgp3m3; rememberMe=true; customerName=dipcc; userName=admin; group1_enable=true; group2_enable=true; group3_enable=true; group4_enable=true; group5_enable=true; group1_time=7; group2_time=14; group3_time=30; srv_session_id=1TCATNOC1REMOTSUC%7CJiHTTvtP1kDHRoz9b%2BzFjt8cusktjFfxA%2BWMaFgrKPgB6Yy5twODd69TT9erT2WTFriTw4B9L5bSRaeyCWKm5v09RH3GDvGwgMiH%2FO7WY2xoS5It%2Fn4cPvdvOa6%2B%2F%2BN9oRujT291Cf7EYOluW%2B87%2F%2BS3lb95vLVcknEp7CqIBhMr%2FUEBomC3kv%2Fq502U5uBgbPqvdSFvnTXRx0tj%2FGQ0nt5r%2F9FwM7LwiPRjigTmeArnvY734CdSgyWB21diDjZ%2FHv7Pg6ZltSz1kAaRI%2BM%2B27VcO5iiNJdIeLQcttB%2Bd0JJ3Gu7jv6qfTmQJwEKNEiuPAvTV%2FAJOJ%2BnOKf4AZdb4a060f2PkumYa1VD7vldEzTeUyFub5mODkWgJnh7LcFdgYJ8m4ydELRi87P6JT8bdytbCFTdASKu7z8Wu2QA7QNsh5eCOPxVKcOsz9ofXynmEneIfhDCHLxTANni7Fp4DD78zn4bQKn68OMSlKfyu2Dt96yS20sUMoG2NnsN%2Fyl2tfGplWd2phsM6fYI%2FIzcq%2BUZj8CWOyYbSxmxGINIM1GRviVNZObkUP0GOUDYdbUanwXyqilCSqARH%2Fzs09MTwMGK6a99hT1DQY%2BBxion5SHBTP3sAMnaY63eVlC5MiO%2Fa3jvbWsTTwwul45QBjFxLi3yEYGmRRKmOI3TAN4tDJGyGCyhhI%2FmlTt3zgQYcrSHtyUHLwjpIq7JLDXShUr%2F2Q%3D%3D")

	//
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	//
	if resp.StatusCode != 200 {
		return
	}

	// TODO 可以检查一下 body
	// bodyBytes, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("读取响应失败:", err)
	// 	return
	// }
	// fmt.Println("响应内容:", string(bodyBytes))

	// 写全局 srv_session_id
	GlobalSrvSessionIdCookie, _ = lo.Find(resp.Cookies(), func(item *http.Cookie) bool {
		return item.Name == "srv_session_id"
	})
	GlobalPHPSESSIONCookie, _ = lo.Find(resp.Cookies(), func(item *http.Cookie) bool {
		return item.Name == "PHPSESSID"
	})

	//
	fmt.Println("login success  ---->  ", resp.StatusCode)
}

// 发送请求
func SendReq(method string, uri string, body string) {
	//
	// req, err := http.NewRequest(method, HostBase+uri, bytes.NewBufferString(body))
	req, err := http.NewRequest(method, HostBase+uri, strings.NewReader(body))
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// set Content-Type
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// add Cookie
	if GlobalSrvSessionIdCookie == nil {
		// TODO 登录认证 cookie srv_session_id 为空，需要重新登录
		fmt.Println("GlobalSrvSessionIdCookie is None  ---->  ")
	}
	req.AddCookie(GlobalSrvSessionIdCookie)
	// req.AddCookie(GlobalPHPSESSIONCookie)
	// req.AddCookie(&http.Cookie{Name: "rememberMe", Value: "false"})
	// req.AddCookie(&http.Cookie{Name: "customerName", Value: ""})
	// req.AddCookie(&http.Cookie{Name: "userName", Value: ""})

	//
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}
	var bodyData map[string]interface{}
	err = json.Unmarshal(bodyBytes, &bodyData)
	if err != nil {
		fmt.Println("响应 body Unmarshal 失败:", err)
	}

	result, _ := bodyData["result"].(map[string]interface{})
	errCode, _ := result["error"].(int)
	if errCode != 0 {
		fmt.Println("响应 body.result.error=", errCode)

		// TODO body.result.error=2 3 是会话问题，需重新登录
	}

	// TODO 存储结果
	fmt.Println("response OK ---->  ", bodyData)
	//
}

func DoTest() {
	//
	Login()

	//
	SendReq("POST", "/service/index.php?m=home&c=home&f=queryAccount", "")
}
