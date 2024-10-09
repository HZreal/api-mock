package main

/**
 * @Author elastic·H
 * @Date 2024-10-09
 * @File: execScript.go
 * @Description:
 */

import (
	"fmt"
)

func t() {
	fmt.Println("t")
}

var (
	loginExecScript = []string{
		"// 提取响应中的 Cookies\r",
		"let cookies = pm.cookies;\r",
		"\r",
		"// 提取 'srv_session_id' Cookie 并存储到环境变量\r",
		"if (cookies.has('srv_session_id')) {\r",
		"    let srvSessionId = cookies.get('srv_session_id');\r",
		"    pm.environment.set('srv_session_id', srvSessionId);\r",
		"    console.log('srv_session_id:', srvSessionId);\r",
		"} else {\r",
		"    console.warn('srv_session_id not found in response cookies.');\r",
		"}\r",
		"\r",
		"// 提取 'PHPSESSID' Cookie 并存储到环境变量\r",
		"if (cookies.has('PHPSESSID')) {\r",
		"    let phpSessionId = cookies.get('PHPSESSID');\r",
		"    pm.environment.set('PHPSESSID', phpSessionId);\r",
		"    console.log('PHPSESSID:', phpSessionId);\r",
		"} else {\r",
		"    console.warn('PHPSESSID not found in response cookies.');\r",
		"}\r",
	}

	responseAssertScript = []string{
		"pm.test(\"msg:成功，错误：0，响应码：200\", function () {\r",
		"    var jsonData = pm.response.json();\r",
		// "    pm.expect(jsonData.result.msg).to.eql(\"成功\");\r",
		"    pm.expect(jsonData.result.error).to.eql(0);\r",
		"    pm.response.to.have.status(200);\r",
		"});\r",
	}

	pprerequestScript = []string{
		"// 生成 1 到 10000 的随机页码\r",
		"let currentPage = Math.floor(Math.random() * 10000) + 1; // 生成 1 到 10000 之间的正整数\r",
		"\r",
		"// 允许的 pageSize 枚举值\r",
		"let pageSizeOptions = [10, 20, 30, 50, 100];\r",
		"let pageSize = pageSizeOptions[Math.floor(Math.random() * pageSizeOptions.length)];\r",
		"\r",
		"// 生成 JSON 字符串，包含分页、排序和过滤信息\r",
		"let requestPayload = {\r",
		"    pagination: {\r",
		"        current: currentPage,\r",
		"        pageSize: pageSize\r",
		"    },\r",
		"    sorter: {},\r",
		"    filter: {}\r",
		"};\r",
		"\r",
		"// 转换为 JSON 字符串\r",
		"let requestPayloadString = JSON.stringify(requestPayload);\r",
		"\r",
		"// 将生成的字符串设置为键 `p` 的值\r",
		"pm.request.body.update({\r",
		"    mode: 'urlencoded',\r",
		"    urlencoded: [\r",
		"        { key: 'p', value: requestPayloadString, type: 'text' }\r",
		"    ]\r",
		"});\r",
		"\r",
		"// 输出到控制台，便于调试\r",
		"console.log('Updated request body:', requestPayloadString);\r",
	}
)
