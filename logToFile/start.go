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
	"gin-init/config"
	"gin-init/constant"
	"gin-init/logToFile/relat"
	"gin-init/model/entity"
	"gin-init/service"
	"gin-init/utils"
	"github.com/google/uuid"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	CollectionSchema string = "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	loginCondition   string = "m=login"

	Mode_Urlencoded string = "urlencoded"
	Mode_Raw        string = "raw"
	Mode_Form_Data  string = "formdata"
	Mode_None       string = "none"
)

// PostmanCollection 不包含目录的，简单枚举
type PostmanCollection struct {
	Info Info   `json:"info"`
	Item []Item `json:"item"`
}

// PostmanCollection2 包含目录的
type PostmanCollection2 struct {
	Info Info      `json:"info"`
	Item []ItemDir `json:"item"`
}

type Info struct {
	PostmanID string `json:"_postman_id"`
	Name      string `json:"name"`
	Schema    string `json:"schema"`
}

type ItemDir struct {
	Name string `json:"name"`
	Item []Item `json:"item"`
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
	Value string `json:"value,omitempty"`
	Type  string `json:"type,omitempty"`
	Src   string `json:"src,omitempty"`
}

type Body struct {
	Mode       string       `json:"mode"`
	Urlencoded []KVItem     `json:"urlencoded,omitempty"`
	FormData   []KVItem     `json:"formdata,omitempty"`
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

func start1(fileName string, collectionName string) {
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

	for _, line := range logEntries {
		//
		path := strings.Split(line.Uri, "/")

		//
		query := parseUrlencodedArgs(line.Args)

		//
		body := Body{}
		testEvent := Event{
			Listen: "test",
			Script: Script{
				Exec: relat.ResponseAssertScript,
				Type: "text/javascript",
			},
		}
		var events []Event
		if strings.Contains(line.Args, loginCondition) {
			testEvent.Script.Exec = relat.LoginExecScript
		}
		events = append(events, testEvent)

		if line.ContentType == constant.APPLICATION_JSON {
			// 几乎没有这种情况
			body.Mode = "raw"
			body.Raw = line.RequestBody
			body.Options = &BodyOptions{
				Raw: RawLanguage{
					Language: "json",
				},
			}
		} else if line.ContentType == constant.APPLICATION_FORM_URLENCODED {
			// 主要是这种情况
			body.Mode = "urlencoded"

			if line.BodyType == service.BodyTypeUrlencodedBodyPString {
				// p=
				body.Urlencoded = []KVItem{
					{Key: "p", Value: "{\"pagination\":{\"current\":1,\"pageSize\":10},\"sorter\":{},\"filter\":{}}"},
				}
				prerequestEvent := Event{
					Listen: "prerequest",
					Script: Script{
						Exec: relat.PprerequestScript,
						Type: "text/javascript",
					},
				}
				events = append(events, prerequestEvent)
			} else if line.BodyType == service.BodyTypeUrlencodedBodyFormUrlEncoded {
				// a=1&b=2
				body.Urlencoded = parseUrlencodedArgs(line.RequestBody)
			}
		} else {
			body.Mode = "none"
			body.Raw = ``
		}

		//
		requestItem := Item{
			Name: line.ReqUriArgs,
			Request: Request{
				Method: line.Method,
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
						Value: line.ContentType,
					},
				},
				URL: URL{
					Raw: "{{url}}" + line.ReqUriArgs,
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

func parseUrlencodedArgs(args string) []KVItem {
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

// 构建 Collection
func exportCollection(timestamp string) *PostmanCollection2 {
	// 创建 Collection
	collection := &PostmanCollection2{
		Info: Info{
			PostmanID: uuid.New().String(),
			Name:      config.Conf.LogToFile.CollectionNamePrefix + timestamp,
			Schema:    CollectionSchema,
		},
	}

	var lastID uint = 0
	for {
		var apiRecords []entity.ApiModel
		// 使用选择特定字段
		result := relat.DB.Where("id > ?", lastID).Order("id").Limit(config.Conf.LogToFile.BatchSize).Find(&apiRecords)
		if result.Error != nil {
			log.Fatalf("failed to retrieve apiRecords: %v", result.Error)
		}

		if len(apiRecords) == 0 {
			break // 没有更多数据
		}

		var itemDir = ItemDir{
			Name: config.Conf.LogToFile.CollectionDirNamePrefix + strconv.Itoa(int(lastID)),
		}

		// 处理用户数据
		for _, line := range apiRecords {

			//
			path := strings.Split(line.Uri, "/")

			//
			query := parseUrlencodedArgs(line.Args)

			//
			body := Body{}

			//
			testEvent := Event{
				Listen: "test",
				Script: Script{
					Exec: relat.ResponseAssertScript,
					Type: "text/javascript",
				},
			}
			var events []Event
			if strings.Contains(line.Args, loginCondition) {
				testEvent.Script.Exec = relat.LoginExecScript
			}
			events = append(events, testEvent)

			if line.ContentType == constant.APPLICATION_JSON {
				//
				body.Mode = Mode_Raw
				body.Raw = line.RequestBody
				body.Options = &BodyOptions{
					Raw: RawLanguage{
						Language: "json",
					},
				}
			} else if line.ContentType == constant.APPLICATION_FORM_URLENCODED {
				// 主要是这种情况
				body.Mode = Mode_Urlencoded

				if line.BodyType == service.BodyTypeUrlencodedBodyPString {
					// p=
					body.Urlencoded = []KVItem{
						{Key: "p", Value: "{\"pagination\":{\"current\":1,\"pageSize\":10},\"sorter\":{},\"filter\":{}}"},
					}
					prerequestEvent := Event{
						Listen: "prerequest",
						Script: Script{
							Exec: relat.PprerequestScript,
							Type: "text/javascript",
						},
					}
					events = append(events, prerequestEvent)
				} else if line.BodyType == service.BodyTypeUrlencodedBodyFormUrlEncoded {
					// a=1&b=2
					bodyKVArr := parseUrlencodedArgs(line.RequestBody)
					if strings.Contains(line.Args, loginCondition) {
						// todo userName 改成变量
					}
					body.Urlencoded = bodyKVArr
				}
			} else if strings.HasPrefix(line.ContentType, constant.MULTIPART_FORM_DATA) {
				// TODO
				body.Mode = Mode_Form_Data
				body.FormData = []KVItem{}
				for _, item := range line.Params {
					example := (item.Example).(string)
					var examplePara struct {
						ContentType string `json:"contentType"`
						Filename    string `json:"filename"`
						Content     string `json:"content"`
					}
					ii := KVItem{Key: item.Name}
					err := json.Unmarshal([]byte(example), &examplePara)
					if err != nil {
						ii.Type = "text"
						ii.Value = example

					} else {
						ii.Type = "file"
						ii.Src = ""
					}
					body.FormData = append(body.FormData, ii)

				}
			} else {
				body.Mode = Mode_None
				body.Raw = ``
			}

			//
			requestItem := Item{
				Name: line.UriArgs,
				Request: Request{
					Method: line.Method,
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
							Value: line.ContentType,
						},
					},
					URL: URL{
						Raw: "{{url}}" + line.UriArgs,
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
			itemDir.Item = append(itemDir.Item, requestItem)
		}

		// 下一个批次的迭代
		lastID = apiRecords[len(apiRecords)-1].Id

		//
		collection.Item = append(collection.Item, itemDir)
	}

	return collection
}

// 提取数据库接口
func dbToJsonFile() {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	collection := exportCollection(timestamp)

	// 检查目录
	CollectionStoreDir := config.Conf.LogToFile.CollectionStoreDir
	if !utils.IsExist(CollectionStoreDir) {
		err := os.MkdirAll(CollectionStoreDir, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}
	storePath := filepath.Join(
		CollectionStoreDir,
		fmt.Sprintf("postman_collection_%s.json", timestamp),
	)
	file, err := os.Create(storePath)
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

	fmt.Printf("Postman collection saved to %s", storePath)
}

// 日志导入到库
func nginxLogToDb() {
	// filePath := "D:/overall/project/api-mock/public/access.1010.log"
	// filePath := flag.String("file", "", "Path to the log file")
	// flag.Parse()
	//
	// if *filePath == "" {
	// 	fmt.Println("Please provide a log file path using the -file flag.")
	// 	os.Exit(1)
	// }
	// fmt.Println("filePath  ---->  ", *filePath)
	// relat.ParseAndImport(filePath)

	//
	relat.ParseAndImport(config.Conf.LogToFile.LogSourcePath)
}

func start2() {
	//
	nginxLogToDb()

	//
	dbToJsonFile()
}

func main() {
	//
	// doTest()

	// fileName := "D:/overall/project/api-mock/public/access.1010.log"
	// fileName := "D:/overall/project/api-mock/public/0930/access_0930_ab.log"
	// start1(fileName, "access_0930_ab.json")

	//
	start2()
}
