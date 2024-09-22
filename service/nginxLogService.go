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
	"io"
	"log"
	"net/url"
	"os"
	"strings"
)

type logLine struct {
	TimeLocal       string `json:"time_local"`
	Method          string `json:"request_method"`
	UriArgs         string `json:"request_uri"`
	Uri             string `json:"uri"`
	ContentType     string `json:"content_type"`
	Args            string `json:"args"`
	RequestBody     string `json:"request_body"`
	Status          string `json:"status"`
	RequestTime     string `json:"request_time"`
	SentContentType string `json:"sent_content_type"`
}

// bodyP := "p=%7B%22pagination%22%3A%7B%22current%22%3A1%2C%22pageSize%22%3A10%7D%2C%22sorter%22%3A%7B%7D%2C%22filter%22%3A%7B%7D%7D"
func ParseBodyP(bodyP string) {

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
}

// sourceType=1&riskTypeStatus=-1
func ParseURLFormEncoded(encoded string) (map[string]string, error) {
	values, err := url.ParseQuery(encoded)
	if err != nil {
		return nil, fmt.Errorf("解析 URL 查询参数失败: %v", err)
	}

	result := make(map[string]string)
	for key, vals := range values {
		if len(vals) > 0 {
			result[key] = vals[0]
		}
	}

	return result, nil
}

func getFilePath() string {
	return "/Users/huang/work/xuntong/api-mock/public/access.log"
}

func unescapeRequestBody(input string) string {
	// 处理 JSON 中的转义字符，如 \x5C\x22 -> \"
	input = strings.ReplaceAll(input, `\x5C\x22`, `"`) // 转义的双引号
	input = strings.ReplaceAll(input, `\x5C`, `\`)     // 转义的反斜杠
	return input
}

func readAndParseLogFile() ([]logLine, error) {
	filePath := getFilePath()

	var logEntries []logLine

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

		// 将 line 转换为字符串
		strLine := string(line)

		// 全局替换非法转义字符，确保 JSON 可以解析
		strLine = unescapeRequestBody(strLine)

		//
		var entry logLine
		if err3 := json.Unmarshal([]byte(strLine), &entry); err3 != nil {
			fmt.Printf("Unmarshal 失败: %s, 错误: %v\n", line, err3)
			continue
		}

		logEntries = append(logEntries, entry)
	}
	return logEntries, nil
}
