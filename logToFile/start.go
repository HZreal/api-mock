package main

/**
 * @Author elastic·H
 * @Date 2024-09-29
 * @File: start.go
 * @Description:
 */

import (
	"encoding/json"
	"fmt"
	"gin-init/service"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

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

type PostmanCollection struct {
	Info Info   `json:"info"`
	Item []Item `json:"item"`
}

type Info struct {
	PostmanID string `json:"_postman_id"`
	Name      string `json:"name"`
	Schema    string `json:"schema"`
}

type Item struct {
	Name     string   `json:"name"`
	Request  Request  `json:"request"`
	Response []string `json:"response"`
	Event    []Event  `json:"event,omitempty"`
}

type Request struct {
	Method string   `json:"method"`
	URL    URL      `json:"url"`
	Header []KVItem `json:"header"`
	Body   Body     `json:"body"`
}

type KVItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Body struct {
	Mode       string       `json:"mode"`
	Urlencoded []KVItem     `json:"urlencoded,omitempty"`
	Raw        string       `json:"raw,omitempty"`
	Options    *BodyOptions `json:"options,omitempty"`
}

type BodyOptions struct {
	Raw RawLanguage `json:"raw"`
}

type RawLanguage struct {
	Language string `json:"language"`
}

type URL struct {
	Raw   string   `json:"raw"`
	Host  []string `json:"host"`
	Path  []string `json:"path"`
	Query []KVItem `json:"query"`
}

type Event struct {
	Listen string `json:"listen"`
	Script Script `json:"script"`
}

type Script struct {
	Exec []string `json:"exec"`
	Type string   `json:"type"`
}

func doTest() {
	// 创建结构体数组
	requests := []Item{
		{
			Name: "home&home&queryAccount",
			Request: Request{
				Method: "POST",
				Header: []KVItem{
					{
						Key:   "X-Requested-With",
						Value: "XMLHttpRequest",
					},
					{
						Key:   "Cookie",
						Value: "srv_session_id={{srv_session_id}}; PHPSESSID={{php_session_id}}; rememberMe=false; customerName=; userName=",
					},
					{
						Key:   "Content-Type",
						Value: "application/x-www-form-urlencoded",
					},
				},
				URL: URL{
					Raw: "{{url}}/service/index.php?m=home&c=home&f=queryAccount",
					Host: []string{
						"{{url}}",
					},
					Path: []string{
						"service",
						"index.php",
					},
					Query: []KVItem{
						{
							Key:   "m",
							Value: "home",
						},
						{
							Key:   "c",
							Value: "home",
						},
						{
							Key:   "f",
							Value: "queryAccount",
						},
					},
				},
				Body: Body{
					Mode: "none",
				},
			},
			Response: []string{},
			Event: []Event{
				{
					Listen: "test",
					Script: Script{
						Exec: []string{
							"//var usertest = postman.getResponseHeader(\"Authorization\");\r",
							"//var usertoken = JSON.parse(responseBody);\r",
							"//pm.environment.set(\"token\",usertoken.data.response.seq);\r",
							"//pm.globals.set(\"test\",usertest);\r",
							"\r",
							"pm.test(\"msg:成功，错误：0，响应码：200\", function () {\r",
							"    var jsonData = pm.response.json();\r",
							"    pm.expect(jsonData.result.msg).to.eql(\"成功\");\r",
							"    pm.expect(jsonData.result.error).to.eql(0);\r",
							"    pm.response.to.have.status(200);\r",
							"});\r",
							"\r",
							"",
						},
						Type: "text/javascript",
					},
				},
			},
		},
		{
			Name: "/api/api/info/list",
			Request: Request{
				Method: "POST",
				Header: []KVItem{
					{
						Key:   "Content-Type",
						Value: "application/json",
					},
				},
				Body: Body{
					Mode: "raw",
					Raw:  `{}`,
					Options: &BodyOptions{
						Raw: RawLanguage{Language: "json"},
					},
				},
				URL: URL{
					Raw: "{{url}}/api/api/info/list?page={{$randomInt}}&pageSize=10&sort=id:asc",
					Host: []string{
						"{{url}}",
					},
					Path: []string{
						"api",
						"api",
						"info",
						"list",
					},
					Query: []KVItem{
						{
							Key:   "page",
							Value: "{{$randomInt}}",
						},
						{
							Key:   "pageSize",
							Value: "10",
						},
						{
							Key:   "sort",
							Value: "id:asc",
						},
					},
				},
			},
			Response: []string{},
			Event:    []Event{},
		},
		{
			Name: "/api/api2/info/list",
			Request: Request{
				Method: "POST",
				Header: []KVItem{
					{
						Key:   "Content-Type",
						Value: "application/json",
					},
				},
				Body: Body{
					Mode: "urlencoded",
					Urlencoded: []KVItem{
						{Key: "p", Value: "{\"pagination\":{\"current\":1,\"pageSize\":10},\"sorter\":{},\"filter\":{}}"},
					},
				},
				URL: URL{
					Raw: "{{url}}/api/api/info/list?page={{$randomInt}}&pageSize=10&sort=id:asc",
					Host: []string{
						"{{url}}",
					},
					Path: []string{
						"api",
						"api",
						"info",
						"list",
					},
					Query: []KVItem{
						{
							Key:   "page",
							Value: "{{$randomInt}}",
						},
						{
							Key:   "pageSize",
							Value: "10",
						},
						{
							Key:   "sort",
							Value: "id:asc",
						},
					},
				},
			},
			Response: []string{},
			Event:    []Event{},
		},
	}

	// 创建 Postman Collection
	collection := PostmanCollection{
		Info: Info{
			PostmanID: "a24ee69f-7e18-4a96-8e6b-191951929321",
			Name:      "生成的 Collection",
			Schema:    "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		},
		Item: requests,
	}

	// 将结构体转换为 JSON
	file, err := os.Create("public/postman_collection_test.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false) // 禁用 HTML 转义

	err = encoder.Encode(collection)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return
	}

	fmt.Println("Postman collection saved to postman_collection_test.json")
}

func parseUrlencodeArgs(args string) []KVItem {
	values, err := url.ParseQuery(args)
	if err != nil {
		return []KVItem{}
	}

	var queries []KVItem
	for key, vals := range values {
		//
		for _, val := range vals {
			queries = append(queries, KVItem{Key: key, Value: val})
		}
	}

	return queries
}

func start(fileName string, collectionName string) {
	// 创建 Postman Collection
	collection := PostmanCollection{
		Info: Info{
			PostmanID: "a24ee69f-7e18-4a96-8e6b-191951929321",
			Name:      "OKCC-Collection-" + strconv.FormatInt(time.Now().Unix(), 10),
			Schema:    "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		},
	}

	//
	logEntries, err := service.ReadAndParseLogFile(fileName)
	if err != nil {
		return
	}

	for _, entry := range logEntries {
		//
		path := strings.Split(entry.Uri, "/")

		//
		query := parseUrlencodeArgs(entry.Args)

		//
		body := Body{}
		testEvent := Event{
			Listen: "test",
			Script: Script{
				Exec: responseAssertScript,
				Type: "text/javascript",
			},
		}
		var events []Event
		if strings.Contains(entry.Args, "m=login") {
			testEvent.Script.Exec = loginExecScript
		}
		events = append(events, testEvent)

		if entry.ContentType == "application/json" {
			// 几乎没有这种情况
			body.Mode = "raw"
			body.Raw = entry.RequestBody
			body.Options = &BodyOptions{
				Raw: RawLanguage{
					Language: "json",
				},
			}
		} else if entry.ContentType == "application/x-www-form-urlencoded" {
			// 主要是这种情况
			body.Mode = "urlencoded"

			if entry.BodyType == 2 {
				// p=
				body.Urlencoded = []KVItem{
					{Key: "p", Value: "{\"pagination\":{\"current\":1,\"pageSize\":10},\"sorter\":{},\"filter\":{}}"},
				}
				prerequestEvent := Event{
					Listen: "prerequest",
					Script: Script{
						Exec: pprerequestScript,
						Type: "text/javascript",
					},
				}
				events = append(events, prerequestEvent)
			} else if entry.BodyType == 3 {
				// a=1&b=2
				body.Urlencoded = parseUrlencodeArgs(entry.RequestBody)
			}
		} else {
			body.Mode = "none"
			body.Raw = ``
		}

		//
		requestItem := Item{
			Name: entry.ReqUriArgs,
			Request: Request{
				Method: entry.Method,
				Header: []KVItem{
					// {
					// 	Key:   "X-Requested-With",
					// 	Value: "XMLHttpRequest",
					// },
					{
						Key:   "Cookie",
						Value: "srv_session_id={{srv_session_id}}; PHPSESSID={{PHPSESSID}}; rememberMe={{rememberMe}}; customerName={{customerName}}; userName={{userName}}",
					},
					{
						Key:   "Content-Type",
						Value: entry.ContentType,
					},
				},
				URL: URL{
					Raw: "{{url}}" + entry.ReqUriArgs,
					Host: []string{
						"{{url}}",
					},
					Path:  path,
					Query: query,
				},
				Body: body,
			},
			Response: []string{},
			Event:    events,
		}

		//
		collection.Item = append(collection.Item, requestItem)
	}

	// 将结构体转换为 JSON
	file, err := os.Create("public/0930/postman_collection_" + collectionName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false) // 禁用 HTML 转义

	err = encoder.Encode(collection)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return
	}

	fmt.Println("Postman collection saved to postman_collection.json")
}

func main() {
	// doTest()

	// fileName := "D:/overall/project/api-mock/public/access.0930.log"
	fileName := "D:/overall/project/api-mock/public/0930/access_0930_ab.log"
	start(fileName, "access_0930_ab.json")
}
