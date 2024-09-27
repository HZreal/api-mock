package service

/**
 * @Author elastic·H
 * @Date 2024-09-27
 * @File: ruleMockService.go
 * @Description:
 */

import (
	"encoding/json"
	"fmt"
	"gin-init/model/entity"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// 模拟生成数据
// TODO ！！！ 待优化，改成 策略枚举的形式，未来可以 自定义规则
func generate(params []*entity.ParamStruct) map[string]interface{} {
	rand.Seed(time.Now().UnixNano())
	result := make(map[string]interface{})

	for _, param := range params {
		mock := param.Mock
		switch param.Type {
		case "string":
			length := extractValueFromMock(mock, "len")
			isIntString := extractValueFromMock(mock, "isIntString")
			if isIntString == 1 {
				result[param.Name] = generateIntString(length)
			} else {
				result[param.Name] = generateRandomString(length)
			}
		case "int":
			length := extractValueFromMock(mock, "len")
			posiOrNega := extractSignFromMock(mock, "posiOrNega")
			result[param.Name] = generateInt(length, posiOrNega)
		case "float64":
			integer := extractValueFromMock(mock, "integer")
			decimal := extractValueFromMock(mock, "decimal")
			posiOrNega := extractSignFromMock(mock, "posiOrNega")
			result[param.Name] = generateFloat(integer, decimal, posiOrNega)
		case "bool":
			result[param.Name] = rand.Intn(2) == 1
		case "emptyObject":
			result[param.Name] = make(map[string]interface{})
		case "emptyArray":
			result[param.Name] = []interface{}{}
		default:
			result[param.Name] = nil
		}
	}

	return result
}

// 从 mock 字符串中提取整数值
func extractValueFromMock(mock, key string) int {
	start := strings.Index(mock, fmt.Sprintf("%s=", key)) + len(key) + 1
	end := strings.Index(mock[start:], ")") + start
	value, _ := strconv.Atoi(mock[start:end])
	return value
}

// 从 mock 字符串中提取正负号
func extractSignFromMock(mock, key string) string {
	start := strings.Index(mock, fmt.Sprintf("%s=", key)) + len(key) + 1
	end := strings.Index(mock[start:], ")") + start
	return mock[start:end]
}

// 生成指定长度的随机字符串
func generateRandomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]rune, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// 生成指定长度的整数字符串
func generateIntString(length int) string {
	result := make([]rune, length)
	for i := range result {
		result[i] = rune('0' + rand.Intn(10))
	}
	return string(result)
}

// 生成指定长度的整数
func generateInt(length int, sign string) int {
	// 特殊处理整数 0
	if sign == "0" {
		return 0
	}

	_min := int(math.Pow10(length - 1))
	_max := int(math.Pow10(length)) - 1

	if sign == "-" {
		return rand.Intn(_max-_min+1) - _min
	}
	return rand.Intn(_max-_min+1) + _min
}

// 生成浮点数
func generateFloat(integer, decimal int, sign string) float64 {
	// 特殊处理浮点数 0.0
	if sign == "0" {
		return 0.0
	}

	integerPart := generateInt(integer, sign)
	decimalPart := rand.Intn(int(math.Pow10(decimal)))

	return float64(integerPart) + float64(decimalPart)/math.Pow10(decimal)
}

func TestMock() {
	params := []*entity.ParamStruct{
		{Name: "key11", Type: "string", Mock: "@string@(len=8)(isIntString=0)"},
		{Name: "key12", Type: "string", Mock: "@string@(len=5)(isIntString=1)"},

		{Name: "key21", Type: "int", Mock: "@int@(posiOrNega=+)(len=3)"},
		{Name: "key22", Type: "int", Mock: "@int@(posiOrNega=-)(len=3)"},
		{Name: "key23", Type: "int", Mock: "@int@(posiOrNega=0)(len=0)"},
		{Name: "key24", Type: "int", Mock: "@int@(posiOrNega=0)(len=1)"},

		{Name: "key31", Type: "float64", Mock: "@float@(posiOrNega=+)(integer=2)(decimal=3)"},
		{Name: "key32", Type: "float64", Mock: "@float@(posiOrNega=-)(integer=1)(decimal=2)"},
		{Name: "key33", Type: "float64", Mock: "@float@(posiOrNega=0)(integer=0)(decimal=3)"},
		{Name: "key34", Type: "float64", Mock: "@float@(posiOrNega=0)(integer=0)(decimal=0)"},

		{Name: "key4", Type: "bool", Mock: "@bool@"},
		{Name: "key5", Type: "emptyObject", Mock: "@emptyObject@"},
		{Name: "key6", Type: "emptyArray", Mock: "@emptyArray@"},
	}

	_map := generate(params)
	fmt.Println("_map  ---->  ", _map)

	marshal, err := json.Marshal(_map)
	if err != nil {
		return
	}
	fmt.Println("marshal  ---->  ", string(marshal))

}
