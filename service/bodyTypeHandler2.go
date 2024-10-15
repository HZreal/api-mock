package service

/**
 * @Author elastic·H
 * @Date 2024-09-30
 * @File: bodyTypeHandler2.go
 * @Description:
 */

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-init/constant"
	"io"
	"mime"
	"mime/multipart"
	"net/url"
	"regexp"
	"strings"
)

const (
	Nothing uint = iota
	BodyTypeUrlencodedBodyEmpty
	BodyTypeUrlencodedBodyPString
	BodyTypeUrlencodedBodyFormUrlEncoded
	BodyTypeUrlencodedBodyJson
	BodyTypeFormDataBody
	BodyTypeEmptyContentType
	BodyTypeFormDataEmptyBody
)

type HandleWayInterface interface {
	Condition(bodyString string, contentType string) bool
	BodyHandle(bodyString string, contentType string) (map[string]interface{}, error)
	GetBodyType() uint
}

type BaseBody struct {
	BodyType uint
}

func (b *BaseBody) GetBodyType() uint {
	return b.BodyType
}

// /////////////////////////////////
type UrlencodedBodyEmpty struct {
	*BaseBody
}

func (p *UrlencodedBodyEmpty) Condition(bodyString string, contentType string) bool {
	return contentType == constant.APPLICATION_FORM_URLENCODED && bodyString == ""
}

func (p *UrlencodedBodyEmpty) BodyHandle(bodyString string, contentType string) (map[string]interface{}, error) {
	return make(map[string]interface{}), nil
}

// /////////////////////////////////
type UrlencodedBodyPString struct {
	*BaseBody
}

func (p *UrlencodedBodyPString) Condition(bodyString string, contentType string) bool {
	return contentType == constant.APPLICATION_FORM_URLENCODED && regexp.MustCompile(`^p=`).MatchString(bodyString)
}

func (p *UrlencodedBodyPString) BodyHandle(bodyString string, contentType string) (map[string]interface{}, error) {
	// 第一步：解析 URL-编码的表单数据
	values, err := url.ParseQuery(bodyString)
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

// /////////////////////////////////
type UrlencodedBodyFormUrlEncoded struct {
	*BaseBody
}

func (p *UrlencodedBodyFormUrlEncoded) Condition(bodyString string, contentType string) bool {
	return contentType == constant.APPLICATION_FORM_URLENCODED && regexp.MustCompile(`^([a-zA-Z0-9_%+-]+=[^&]*)+(&[a-zA-Z0-9_%+-]+=[^&]*)*$`).MatchString(bodyString)
}

func (p *UrlencodedBodyFormUrlEncoded) BodyHandle(bodyString string, contentType string) (map[string]interface{}, error) {
	values, err := url.ParseQuery(bodyString)
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

// /////////////////////////////////
type UrlencodedBodyJson struct {
	*BaseBody
}

func (p *UrlencodedBodyJson) Condition(bodyString string, contentType string) bool {
	return contentType == constant.APPLICATION_FORM_URLENCODED && json.Valid([]byte(bodyString))
}

func (p *UrlencodedBodyJson) BodyHandle(bodyString string, contentType string) (map[string]interface{}, error) {
	var body map[string]interface{}
	_ = json.Unmarshal([]byte(bodyString), &body)
	return body, nil
}

// /////////////////////////////////
type FormDataEmptyBody struct {
	*BaseBody
}

func (p *FormDataEmptyBody) Condition(bodyString string, contentType string) bool {
	return strings.Contains(contentType, constant.MULTIPART_FORM_DATA) && bodyString == ""
}

func (p *FormDataEmptyBody) BodyHandle(bodyString string, contentType string) (map[string]interface{}, error) {
	return make(map[string]interface{}), nil
}

// /////////////////////////////////
type FormDataBody struct {
	*BaseBody
}

func (p *FormDataBody) Condition(bodyString string, contentType string) bool {
	return strings.Contains(contentType, constant.MULTIPART_FORM_DATA) && bodyString != ""
}

func (p *FormDataBody) BodyHandle(bodyString string, contentType string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	// 解析Content-Type获取boundary
	_, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return nil, fmt.Errorf("解析Content-Type失败: %v", err)
	}
	boundary := params["boundary"]

	// 创建multipart reader
	reader := multipart.NewReader(strings.NewReader(bodyString), boundary)

	// 读取每个部分
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("读取part失败: %v", err)
		}

		// 读取part的内容
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(part)
		if err != nil {
			return nil, fmt.Errorf("读取part内容失败: %v", err)
		}

		content := buf.String()

		// 如果有文件名，则这是一个文件上传字段
		if part.FileName() != "" {
			contentArr, _ := json.Marshal(map[string]string{
				"contentType": part.Header.Get("Content-Type"),
				"filename":    part.FileName(),
				"content":     "file content",
			})
			content = string(contentArr)
		}

		// 将字段信息添加到结果map
		result[part.FormName()] = content
	}

	return result, nil
}

// /////////////////////////////////
type EmptyContentTypeBody struct {
	*BaseBody
}

func (p *EmptyContentTypeBody) Condition(bodyString string, contentType string) bool {
	return contentType == "" && bodyString == ""
}

func (p *EmptyContentTypeBody) BodyHandle(bodyString string, contentType string) (map[string]interface{}, error) {
	return make(map[string]interface{}), nil
}

var Handlers = []HandleWayInterface{
	&EmptyContentTypeBody{BaseBody: &BaseBody{BodyType: BodyTypeEmptyContentType}},
	&UrlencodedBodyEmpty{BaseBody: &BaseBody{BodyType: BodyTypeUrlencodedBodyEmpty}},
	&UrlencodedBodyPString{BaseBody: &BaseBody{BodyType: BodyTypeUrlencodedBodyPString}},
	&UrlencodedBodyFormUrlEncoded{BaseBody: &BaseBody{BodyType: BodyTypeUrlencodedBodyFormUrlEncoded}},
	&UrlencodedBodyJson{BaseBody: &BaseBody{BodyType: BodyTypeUrlencodedBodyJson}},
	&FormDataEmptyBody{BaseBody: &BaseBody{BodyType: BodyTypeFormDataEmptyBody}},
	&FormDataBody{BaseBody: &BaseBody{BodyType: BodyTypeFormDataBody}},
}

func GetHandlers() []HandleWayInterface {
	return Handlers
}
