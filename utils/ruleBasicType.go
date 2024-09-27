package utils

/**
 * @Author elastic·H
 * @Date 2024-09-27
 * @File: ruleBasicType.go
 * @Description:
 */

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Number interface {
	int | float64
}

// CheckPositiveOrNegative 判断整数是正数还是负数
func CheckPositiveOrNegative[T Number](num T) string {
	if num > 0 {
		return "+"
	} else if num < 0 {
		return "-"
	}
	return "0"
}

// DigitCount 判断整数是几位数
func DigitCount(num int) int {
	if num == 0 {
		return 1
	}
	return int(math.Log10(math.Abs(float64(num)))) + 1
}

// IsIntegerString 判断整数是否是整型字符串
func IsIntegerString(s string) int {
	_, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return 1
}

// FloatIsInteger 浮点数是否可整数
func FloatIsInteger(val float64) bool {
	_, frac := math.Modf(val)
	return frac == 0
}

// CountDigits 获取浮点数小数点前后的位数
func CountDigits(num float64) (int, int) {
	// 取绝对值
	absNum := math.Abs(num)

	// 小数点前的位数
	integerPart := int(absNum)
	integerDigits := len(strconv.Itoa(integerPart))

	// 小数点后的位数
	numStr := fmt.Sprintf("%f", absNum)
	decimalPart := strings.Split(numStr, ".")[1]
	decimalPart = strings.TrimRight(decimalPart, "0") // 去掉末尾的0
	decimalDigits := len(decimalPart)

	return integerDigits, decimalDigits
}
