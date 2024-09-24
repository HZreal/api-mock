package service

/**
 * @Author elastic·H
 * @Date 2024-09-19
 * @File: nginxLogService.go
 * @Description:
 */

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gin-init/model/entity"
	"io"
	"log"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type logLine struct {
	TimeLocal       string `json:"time_local"`
	Method          string `json:"request_method"`
	ReqUriArgs      string `json:"request_uri"`
	Uri             string `json:"uri"`
	ContentType     string `json:"content_type"`
	Args            string `json:"args"`
	RequestBody     string `json:"request_body"`
	Status          string `json:"status"`
	RequestTime     string `json:"request_time"`
	SentContentType string `json:"sent_content_type"`
}

type record struct {
	Method      string
	ReqUriArgs  string
	Uri         string
	ContentType string
	Args        string
	RequestBody string
}

// TODO 部分过滤的黑名单
var (
	prefixBlackList = []string{""}
	suffixBlackList = []string{""}
)

// ParseBodyP
// bodyP := "p=%7B%22pagination%22%3A%7B%22current%22%3A1%2C%22pageSize%22%3A10%7D%2C%22sorter%22%3A%7B%7D%2C%22filter%22%3A%7B%7D%7D"
func ParseBodyP(bodyP string) map[string]interface{} {

	// 第一步：解析 URL-编码的表单数据
	values, err := url.ParseQuery(bodyP)
	if err != nil {
		log.Fatalf("解析 URL 查询参数失败: %v", err)
	}

	// 提取 'p' 参数的值
	pValues, exists := values["p"]
	if !exists || len(pValues) == 0 {
		log.Fatalf("'p' 参数不存在或为空")
	}
	pEncoded := pValues[0]

	// 第二步：URL 解码
	pDecoded, err := url.QueryUnescape(pEncoded)
	if err != nil {
		log.Fatalf("URL 解码失败: %v", err)
	}

	fmt.Println("解码后的 JSON 字符串:", pDecoded)

	// 第三步：JSON 反序列化
	var parsedData map[string]interface{}
	err = json.Unmarshal([]byte(pDecoded), &parsedData)
	if err != nil {
		log.Fatalf("JSON 反序列化失败: %v", err)
	}

	// 输出解析结果
	fmt.Printf("解析后的 JSON 对象: %+v\n", parsedData)
	return parsedData
}

// ParseURLFormEncoded
// sourceType=1&riskTypeStatus=-1
func ParseURLFormEncoded(encoded string) (map[string]interface{}, error) {
	values, err := url.ParseQuery(encoded)
	if err != nil {
		return nil, fmt.Errorf("解析 URL 查询参数失败: %v", err)
	}

	result := make(map[string]interface{})
	for key, vals := range values {
		if len(vals) > 0 {
			result[key] = vals[0]
		}
	}

	return result, nil
}

func getFilePath() string {
	return "D:/overall/project/api-mock/public/access.log"
}

func unescapeRequestBody(input string) string {
	// 处理 JSON 中的转义字符，如 \x5C\x22 -> \"
	input = strings.ReplaceAll(input, `\x5C\x22`, `"`) // 转义的双引号
	input = strings.ReplaceAll(input, `\x5C`, `\`)     // 转义的反斜杠
	return input
}

// unescapeHex
// 处理转义
func unescapeHex(s string) (string, error) {
	re := regexp.MustCompile(`\\x([0-9A-Fa-f]{2})`)
	fixed := re.ReplaceAllStringFunc(s, func(match string) string {
		hex := match[2:]
		b, err := strconv.ParseUint(hex, 16, 8)
		if err != nil {
			return match // 保留原始
		}
		return string(b)
	})
	return fixed, nil
}

func readAndParseLogFile() ([]record, error) {
	//
	filePath := getFilePath()

	var recordArr []record

	//
	openFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer openFile.Close()

	//
	reader := bufio.NewReader(openFile)
	for {
		line, _, err2 := reader.ReadLine()
		if err2 == io.EOF {
			break
		}
		if err2 != nil {
			continue
		}

		strLine := string(line)

		// 全局替换非法转义字符，确保 JSON 可以解析
		fixedLine, errr := unescapeHex(strLine)
		if errr != nil {
			log.Printf("修复转义序列失败: %v\n内容: %s", errr, strLine)
			panic(errr)
		}

		//
		var entry logLine
		if err3 := json.Unmarshal([]byte(fixedLine), &entry); err3 != nil {
			fmt.Printf("日志行 Unmarshal 失败: %s, 错误: %v\n", line, err3)
			continue
		}

		// 过滤掉静态资源请求, 后续可配置名单
		if strings.HasSuffix(entry.Uri, ".js") ||
			strings.HasSuffix(entry.Uri, ".png") ||
			strings.HasPrefix(entry.Uri, "/skin") {
			continue
		}

		// TODO 重复问题，更新

		// 处理空值
		if entry.ContentType == "-" {
			entry.ContentType = ""
		}
		if entry.Args == "-" {
			entry.Args = ""
		}
		if entry.RequestBody == "-" {
			entry.RequestBody = ""
		}

		//
		recordItem := record{
			Method:      entry.Method,
			ReqUriArgs:  entry.ReqUriArgs,
			Uri:         entry.Uri,
			ContentType: entry.ContentType,
			Args:        entry.Args,
			RequestBody: entry.RequestBody,
		}
		recordArr = append(recordArr, recordItem)
	}
	return recordArr, nil
}

// 处理 body
func parseBody(lineBody string) (params []entity.ParamStruct) {
	var body map[string]interface{}

	if lineBody == "" {
		//
		params = []entity.ParamStruct{}
	} else if strings.HasPrefix(lineBody, "p=") {
		//
		body = ParseBodyP(lineBody)
	} else if strings.Contains(lineBody, "=") {
		//
		body, _ = ParseURLFormEncoded(lineBody)
	} else {
		err2 := json.Unmarshal([]byte(lineBody), &body)
		if err2 != nil {
			fmt.Println("lineBody Unmarshal 失败, params 为空", err2)
			params = []entity.ParamStruct{}
		} else {
			//
			fmt.Println("lineBody Unmarshal成功 body  ---->  ", body)
		}
	}

	// TODO 处理 body 成参数
	for k, v := range body {
		var itemType string
		switch v.(type) {
		case string:
			itemType = "string"
		case int:
			itemType = "int"
		case float32:
			itemType = "float32"
		case float64:
			itemType = "float64"
		case bool:
			itemType = "bool"
		case map[string]interface{}:
			itemType = "object"
		default:
			itemType = "unknown"
		}
		item := entity.ParamStruct{Name: k, Type: itemType}
		params = append(params, item)
	}

	return
}
