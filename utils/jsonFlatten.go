package utils

/**
 * @Author elastic·H
 * @Date 2024-09-24
 * @File: jsonFlatten.go
 * @Description:
 */

import "encoding/json"

// Flatten 将嵌套的 JSON 对象扁平化为一个层级的 map
func Flatten(input map[string]interface{}, prefix string, result map[string]interface{}) {
	for key, value := range input {
		newKey := key
		if prefix != "" {
			newKey = prefix + "." + key
		}

		switch v := value.(type) {
		case map[string]interface{}:
			Flatten(v, newKey, result)
		default:
			result[newKey] = v
		}
	}
}

// FlattenJSON 将 JSON 字符串扁平化为 map
func FlattenJSON(jsonStr string) (map[string]interface{}, error) {
	var input map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &input); err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	Flatten(input, "", result)
	return result, nil
}

// Unflatten 将扁平化的 map 转换为嵌套的 JSON 对象
func Unflatten(input map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for key, value := range input {
		parts := splitKey(key)
		current := result

		for i, part := range parts {
			if i == len(parts)-1 {
				current[part] = value
			} else {
				if _, exists := current[part]; !exists {
					current[part] = make(map[string]interface{})
				}
				current = current[part].(map[string]interface{})
			}
		}
	}
	return result
}

// splitKey 将扁平化的键拆分为切片
func splitKey(key string) []string {
	return []string{}
}
