package utils

/**
 * @Author elastic·H
 * @Date 2024-10-12
 * @File: file.go
 * @Description:
 */

import (
	"os"
)

// IsExist 检查文件或者文件夹是否存在
func IsExist(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}
