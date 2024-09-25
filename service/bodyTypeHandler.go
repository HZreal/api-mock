package service

/**
 * @Author elastic·H
 * @Date 2024-09-24
 * @File: bodyTypeHandler.go
 * @Description: 不同 body 类型不同的处理器
 */

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
)

type HandleWay struct {
	Condition func(s string) bool
	Handle    func(s string) (map[string]interface{}, error)
}

var handlers = []HandleWay{
	{
		Condition: ConditionEmpty,
		Handle:    ParseBodyEmpty,
	},
	{
		Condition: ConditionBodyP,
		Handle:    ParseBodyP,
	},
	{
		Condition: ConditionBodyFormUrlEncoded,
		Handle:    ParseBodyFormUrlEncoded,
	},
	{
		Condition: ConditionBodyJson,
		Handle:    ParseBodyJson,
	},
}

// ConditionEmpty -
func ConditionEmpty(s string) bool {
	return s == ""
}

// ParseBodyEmpty -
func ParseBodyEmpty(bodyEmpty string) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}

// ConditionBodyP
func ConditionBodyP(s string) bool {
	return regexp.MustCompile(`^p=`).MatchString(s)
}

// ParseBodyP
// bodyP := "p=%7B%22pagination%22%3A%7B%22current%22%3A1%2C%22pageSize%22%3A10%7D%2C%22sorter%22%3A%7B%7D%2C%22filter%22%3A%7B%7D%7D"
func ParseBodyP(bodyP string) (map[string]interface{}, error) {
	// 第一步：解析 URL-编码的表单数据
	values, err := url.ParseQuery(bodyP)
	if err != nil {
		return nil, fmt.Errorf("解析 URL 查询参数失败: %w", err)
	}

	// 提取 'p' 参数的值
	pValues, exists := values["p"]
	if !exists || len(pValues) == 0 {
		return nil, fmt.Errorf("'p' 参数不存在或为空")
	}
	pEncoded := pValues[0]

	// 第二步：URL 解码
	pDecoded, err := url.QueryUnescape(pEncoded)
	if err != nil {
		return nil, fmt.Errorf("URL 解码失败: %w", err)
	}

	fmt.Println("解码后的 JSON 字符串:", pDecoded)

	// 第三步：JSON 反序列化
	var parsedData map[string]interface{}
	err = json.Unmarshal([]byte(pDecoded), &parsedData)
	if err != nil {
		return nil, fmt.Errorf("JSON 反序列化失败: %w", err)
	}

	// 输出解析结果
	fmt.Printf("解析后的 JSON 对象: %+v\n", parsedData)
	return parsedData, nil
}

// ConditionBodyFormUrlEncoded
func ConditionBodyFormUrlEncoded(s string) bool {
	return regexp.MustCompile(`^([a-zA-Z0-9_%+-]+=[^&]*)+(&[a-zA-Z0-9_%+-]+=[^&]*)*$`).MatchString(s)
}

// ParseBodyFormUrlEncoded
// sourceType=1&riskTypeStatus=-1
func ParseBodyFormUrlEncoded(encoded string) (map[string]interface{}, error) {
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

func ConditionBodyJson(s string) bool {
	return json.Valid([]byte(s))
}

// ParseBodyJson
// 解析 JSON
func ParseBodyJson(jsonStr string) (map[string]interface{}, error) {
	var body map[string]interface{}
	_ = json.Unmarshal([]byte(jsonStr), &body)
	return body, nil
}
