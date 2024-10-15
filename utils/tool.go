package utils

/**
 * @Author elastic·H
 * @Date 2024-10-14
 * @File: tool.go
 * @Description:
 */

import (
	"golang.org/x/exp/constraints"
)

// MergeArrays 泛型函数，合并两个数组，避免重复元素
func MergeArrays[T any, K constraints.Ordered](
	target, source []T,
	keyFunc func(T) K, // 用于获取每个元素的唯一标识符
) []T {
	exists := make(map[K]bool)
	for _, item := range target {
		exists[keyFunc(item)] = true
	}

	for _, item := range source {
		if !exists[keyFunc(item)] {
			target = append(target, item)
			exists[keyFunc(item)] = true
		}
	}

	return target
}
