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
	"gin-init/utils"
	"io"
	"log"
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
	Method            string
	ReqUriArgs        string
	Uri               string
	ContentType       string
	Args              string
	RequestBodyParams []*entity.ParamStruct
}

// TODO 部分过滤的黑名单
var (
	prefixBlackList = []string{""}
	suffixBlackList = []string{""}
)

// TODO 后续改成访问 nginx 日志 目前无法直接获取虚拟机内的日志文件
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
		// return string(b)
		return strconv.FormatUint(b, 10)
	})
	return fixed, nil
}

// readAndParseLogFile
func readAndParseLogFile() ([]*record, error) {
	//
	filePath := getFilePath()

	var recordArr []*record

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
		fixedLine, err3 := unescapeHex(strLine)
		if err3 != nil {
			log.Printf("修复转义序列失败: %v\n内容: %s", err3, strLine)
			continue
		}

		//
		var entry logLine
		if err4 := json.Unmarshal([]byte(fixedLine), &entry); err4 != nil {
			fmt.Printf("日志行 Unmarshal 失败: %s, 错误: %v\n", line, err4)
			continue
		}

		// TODO 过滤掉静态资源请求, 后续可配置名单
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

		// 处理 RequestBody
		params := parseBody(entry.RequestBody)

		//
		recordItem := &record{
			Method:            entry.Method,
			ReqUriArgs:        entry.ReqUriArgs,
			Uri:               entry.Uri,
			ContentType:       entry.ContentType,
			Args:              entry.Args,
			RequestBodyParams: params,
		}
		recordArr = append(recordArr, recordItem)
	}
	return recordArr, nil
}

// 处理 body
func parseBody(lineBody string) (params []*entity.ParamStruct) {
	var body map[string]interface{}

	// TODO 改成策略的形式 类似 map 包含 条件、处理函数
	// if lineBody == "" {
	// 	//
	// 	params = []entity.ParamStruct{}
	// } else if regexp.MustCompile(`^p=`).MatchString(lineBody) {
	// 	// else if strings.HasPrefix(lineBody, "p=") {
	// 	body, _ = ParseBodyP(lineBody)
	// } else if regexp.MustCompile(`^([a-zA-Z0-9_%+-]+=[^&]*)+(&[a-zA-Z0-9_%+-]+=[^&]*)*$`).MatchString(lineBody) {
	// 	// else if strings.Contains(lineBody, "=") {
	// 	body, _ = ParseBodyFormUrlEncoded(lineBody)
	// } else if json.Valid([]byte(lineBody)) {
	// 	fmt.Println("lineBody 为 json 对象")
	// 	_ = json.Unmarshal([]byte(lineBody), &body)
	// } else {
	// 	//
	// 	fmt.Println("未知类型，请检查 ---->  ")
	// 	params = []entity.ParamStruct{}
	// }

	for _, handleWay := range handlers {
		condition := handleWay.Condition
		Handle := handleWay.Handle

		if condition(lineBody) {
			var err error
			body, err = Handle(lineBody)
			if err != nil {
				fmt.Println("error in Handle", lineBody)
			}
			break
		}
	}

	// TODO 扁平化操作
	params = bodyParamsToParamStruct(body)

	return
}

// TODO 扁平化操作
func bodyParamsToParamStruct(body map[string]interface{}) (params []*entity.ParamStruct) {
	if len(body) == 0 {
		return
	}

	//
	bodyFlatten := utils.Flatten(body)
	// 处理成参数对象
	// for k, v := range body {
	for k, v := range bodyFlatten {
		var itemType string
		var mock string
		switch i := v.(type) {
		case string:
			//
			itemType = "string"
			//
			isIntString := utils.IsIntegerString(i)
			mock = fmt.Sprintf("@string@(len=%d)(isIntString=%d)", len(i), isIntString)
		case int:
			//
			itemType = "int"
			//
			posiOrNega := utils.CheckPositiveOrNegative(i)
			lenCount := utils.DigitCount(i)
			mock = fmt.Sprintf("@int@(posiOrNega=%s)(len=%d)", posiOrNega, lenCount)
		case float64:
			itemType = "float64"
			//
			posiOrNega := utils.CheckPositiveOrNegative(i)
			integerDigits, decimalDigits := utils.CountDigits(i)
			mock = fmt.Sprintf("@float@(posiOrNega=%s)(integer=%d)(decimal=%d)", posiOrNega, integerDigits, decimalDigits)
		case bool:
			itemType = "bool"
			mock = fmt.Sprintf("@bool@")
		case map[string]interface{}:
			if len(i) != 0 {
				log.Printf("不是空 map, Flatten 可能出错；k=%s,v=%s", k, v)
				continue
			}
			itemType = "emptyObject"
			//
			mock = fmt.Sprintf("@emptyObject@")
		case []interface{}:
			if len(i) != 0 {
				log.Printf("不是空 [], Flatten 可能出错；k=%s,v=%s", k, v)
				continue
			}
			itemType = "emptyArray"
			//
			mock = fmt.Sprintf("@emptyArray@")
		default:
			itemType = "unknown"
		}
		item := &entity.ParamStruct{Name: k, Type: itemType, Mock: mock}
		params = append(params, item)
	}

	return
}
