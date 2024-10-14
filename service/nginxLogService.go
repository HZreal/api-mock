package service

/**
 * @Author elastic·H
 * @Date 2024-09-19
 * @File: nginxLogService.go
 * @Description:
 */

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"gin-init/model/entity"
	"gin-init/utils"
	"io"
	"log"
	"net/url"
	"os"
	"regexp"
	"sort"
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

type Record struct {
	Name              string
	Method            string
	ReqUriArgs        string
	Uri               string
	ContentType       string
	Args              string
	BodyType          uint
	RequestBody       string
	RequestBodyParams []*entity.ParamStruct
	SentContentType   string `json:"sent_content_type"`
}

// TODO 部分过滤的黑名单
var (
	prefixBlackList = []string{""}
	suffixBlackList = []string{""}
)

// TODO 后续改成访问 nginx 日志 目前无法直接获取虚拟机内的日志文件
func getFilePath() string {
	return "D:/overall/project/api-mock/public/access.0930.log"
}

func unescapeRequestBody(input string) string {
	// 处理 JSON 中的转义字符，如 \x5C\x22 -> \"
	input = strings.ReplaceAll(input, `\x5C\x22`, `"`) // 转义的双引号
	input = strings.ReplaceAll(input, `\x5C`, `\`)     // 转义的反斜杠
	return input
}

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

func unescapeHex2(s string) (string, error) {
	// 1. 处理双引号的转义
	s = strings.ReplaceAll(s, `\"`, `"`)

	// 2. 处理十六进制转义字符
	re := regexp.MustCompile(`\\x([0-9A-Fa-f]{2})`)
	fixed := re.ReplaceAllStringFunc(s, func(match string) string {
		hex := match[2:]
		b, err := strconv.ParseUint(hex, 16, 8)
		if err != nil {
			return match // 保留原始字符串
		}
		return string(rune(b)) // 直接返回对应的字符
	})

	// 3. 处理其他常见转义字符
	replacer := strings.NewReplacer(
		`\\`, `\`,
		`\r`, "\r",
		`\n`, "\n",
		`\t`, "\t",
	)
	fixed = replacer.Replace(fixed)

	return fixed, nil
}

func preprocessLine(line string) (string, error) {
	// 匹配 \x 后跟两个十六进制字符的模式
	re := regexp.MustCompile(`\\x[0-9A-Fa-f]{2}`)

	// 将匹配到的 \xHH 转换为对应的 Unicode 转义字符 \u00HH
	processedLine := re.ReplaceAllStringFunc(line, func(match string) string {
		// 提取十六进制部分
		hexValue := match[2:]
		return fmt.Sprintf("\\u00%s", hexValue)
	})

	return processedLine, nil
}

func unescapeLog(input string) (string, error) {
	// 将 \xXX 转换为对应字符
	re := regexp.MustCompile(`\\x([0-9A-Fa-f]{2})`)
	result := re.ReplaceAllStringFunc(input, func(s string) string {
		var b byte
		fmt.Sscanf(s[2:], "%X", &b)
		return string(b)
	})

	// 处理其他转义字符
	result = string(bytes.ReplaceAll([]byte(result), []byte(`\"`), []byte(`"`)))
	result = string(bytes.ReplaceAll([]byte(result), []byte(`\\`), []byte(`\`)))

	return result, nil
}

func unescapeJSONString(s string) (string, error) {
	// Handle \xXX escapes
	re := regexp.MustCompile(`\\x([0-9A-Fa-f]{2})`)
	s = re.ReplaceAllStringFunc(s, func(match string) string {
		code, _ := strconv.ParseUint(match[2:], 16, 8)
		return string(rune(code))
	})

	// Handle standard JSON escapes
	s = strings.NewReplacer(
		`\"`, `"`,
		`\\`, `\`,
		`\n`, "\n",
		`\r`, "\r",
		`\t`, "\t",
		`\b`, "\b",
		`\f`, "\f",
	).Replace(s)

	return s, nil
}

// ReadAndParseLogFile
func ReadAndParseLogFile(filePath string) ([]*Record, error) {
	//
	var recordArr []*Record

	//
	openFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer openFile.Close()

	//
	reader := bufio.NewReaderSize(openFile, 16*1024)
	for {
		line, _, err2 := reader.ReadLine()
		if err2 == io.EOF {
			break
		}
		if err2 != nil {
			continue
		}

		strLine := string(line)

		// 全局替换转义字符，确保 JSON 可以解析
		fixedLine, err3 := preprocessLine(strLine)
		// fixedLine, err3 := unescapeLog(strLine)
		// fixedLine, err3 := strconv.Unquote(`"` + strLine + `"`)
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
		if strings.HasSuffix(entry.ReqUriArgs, ".js") ||
			strings.HasSuffix(entry.Uri, ".png") ||
			strings.HasSuffix(entry.Uri, ".txt") ||
			strings.HasSuffix(entry.Uri, ".xlsx") ||
			strings.HasSuffix(entry.Uri, ".tar") ||
			strings.Contains(entry.Uri, "favicon.ico") ||
			strings.HasPrefix(entry.Uri, "/pub") ||
			strings.Contains(entry.Args, "f=logout") ||
			strings.HasPrefix(entry.Uri, "/skin") {
			continue
		}
		if !strings.HasSuffix(entry.Uri, ".php") {
			continue
		}

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

		// 提取 args
		queryKeys := entry.Args
		queryMaps, err := url.ParseQuery(entry.Args)
		if err == nil {
			mfc := make([]string, 0, len(queryMaps))
			for k, v := range queryMaps {
				if k == "m" || k == "f" || k == "c" {
					sort.Strings(v)
					mfc = append(mfc, k+"="+strings.Join(v, ","))
				}
			}
			sort.Strings(mfc)
			queryKeys = strings.Join(mfc, "|")
		}
		Name := entry.Method + "-" + entry.Uri + "-" + queryKeys

		// 处理 RequestBody
		params, bodyType := parseBody(entry.RequestBody, entry.ContentType)

		//
		recordItem := &Record{
			Name:              Name,
			Method:            entry.Method,
			ReqUriArgs:        entry.ReqUriArgs,
			Uri:               entry.Uri,
			ContentType:       entry.ContentType,
			Args:              entry.Args,
			RequestBody:       entry.RequestBody,
			BodyType:          bodyType,
			RequestBodyParams: params,
			SentContentType:   entry.SentContentType,
		}
		recordArr = append(recordArr, recordItem)
	}
	return recordArr, nil
}

// 处理 body
func parseBody(lineBody string, contentType string) (params []*entity.ParamStruct, bodyType uint) {
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

	// for _, handleWay := range handlers {
	// 	condition := handleWay.Condition
	// 	Handle := handleWay.Handle
	//
	// 	if condition(lineBody) {
	// 		var err error
	// 		body, bodyType, err = Handle(lineBody)
	// 		if err != nil {
	// 			fmt.Println("error in Handle", lineBody)
	// 		}
	// 		break
	// 	}
	// }

	// body 是否匹配的标记
	var matchFlag = false
	for _, handlerItem := range Handlers {
		if handlerItem.Condition(lineBody, contentType) {
			var err error
			body, err = handlerItem.BodyHandle(lineBody, contentType)
			if err != nil {
				fmt.Println("error in BodyHandle", lineBody)
				break
			}
			bodyType = handlerItem.GetBodyType()
			matchFlag = true
			break
		}
	}
	if !matchFlag {
		log.Printf("Not match Flag Handlers ---->  %s", lineBody)
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
		case float64:
			if utils.FloatIsInteger(i) {
				// int
				itemType = "int"
				//
				posiOrNega := utils.CheckPositiveOrNegative(i)
				lenCount := utils.DigitCount(int(i))
				mock = fmt.Sprintf("@int@(posiOrNega=%s)(len=%d)", posiOrNega, lenCount)
			} else {
				// float
				itemType = "float64"
				//
				posiOrNega := utils.CheckPositiveOrNegative(i)
				integerDigits, decimalDigits := utils.CountDigits(i)
				mock = fmt.Sprintf("@float@(posiOrNega=%s)(integer=%d)(decimal=%d)", posiOrNega, integerDigits, decimalDigits)
			}
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

func Import() ([]*Record, error) {
	filePath := getFilePath()
	return ReadAndParseLogFile(filePath)
}
