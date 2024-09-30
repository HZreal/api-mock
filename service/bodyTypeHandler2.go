package service

/**
 * @Author elastic·H
 * @Date 2024-09-30
 * @File: bodyTypeHandler2.go
 * @Description:
 */

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
)

type HandleWayInterface interface {
	Condition(s string) bool
	BodyHandle(s string) (map[string]interface{}, error)
	GetBodyType() uint
}

type BaseBody struct {
	BodyType uint
}

func (b *BaseBody) GetBodyType() uint {
	return b.BodyType
}

type BodyEmpty struct {
	*BaseBody
}

func (p *BodyEmpty) Condition(s string) bool {
	return s == ""
}

func (p *BodyEmpty) BodyHandle(s string) (map[string]interface{}, error) {
	return make(map[string]interface{}), nil
}

type BodyPString struct {
	*BaseBody
}

func (p *BodyPString) Condition(s string) bool {
	return regexp.MustCompile(`^p=`).MatchString(s)
}

func (p *BodyPString) BodyHandle(s string) (map[string]interface{}, error) {
	// 第一步：解析 URL-编码的表单数据
	values, err := url.ParseQuery(s)
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

type BodyFormUrlEncoded struct {
	*BaseBody
}

func (p *BodyFormUrlEncoded) Condition(s string) bool {
	return regexp.MustCompile(`^([a-zA-Z0-9_%+-]+=[^&]*)+(&[a-zA-Z0-9_%+-]+=[^&]*)*$`).MatchString(s)
}

func (p *BodyFormUrlEncoded) BodyHandle(s string) (map[string]interface{}, error) {
	values, err := url.ParseQuery(s)
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

type BodyJson struct {
	*BaseBody
}

func (p *BodyJson) Condition(s string) bool {
	return json.Valid([]byte(s))
}

func (p *BodyJson) BodyHandle(s string) (map[string]interface{}, error) {
	var body map[string]interface{}
	_ = json.Unmarshal([]byte(s), &body)
	return body, nil
}

var Handlers = []HandleWayInterface{
	&BodyEmpty{BaseBody: &BaseBody{BodyType: 1}},
	&BodyPString{BaseBody: &BaseBody{BodyType: 2}},
	&BodyFormUrlEncoded{BaseBody: &BaseBody{BodyType: 3}},
	&BodyJson{BaseBody: &BaseBody{BodyType: 4}},
}
