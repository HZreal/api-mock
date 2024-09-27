package utils

/**
 * @Author elastic·H
 * @Date 2024-09-24
 * @File: jsonFlatten.go
 * @Description:
 */

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

func Flatten(input map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	flattenHelper(input, "", result)
	return result
}

func flattenHelper(input interface{}, prefix string, result map[string]interface{}) {
	switch v := input.(type) {
	case map[string]interface{}:
		if len(v) == 0 {
			result[prefix] = v
		} else {
			for key, value := range v {
				newKey := key
				if prefix != "" {
					newKey = prefix + "." + key
				}
				flattenHelper(value, newKey, result)
			}
		}
	case []interface{}:
		if len(v) == 0 {
			result[prefix] = v
		} else {
			for i, item := range v {
				newKey := fmt.Sprintf("%s[%d]", prefix, i)
				flattenHelper(item, newKey, result)
			}
		}
	default:
		result[prefix] = v
	}
}

func Unflatten(flatMap map[string]interface{}) map[string]interface{} {
	// The resulting nested structure.
	result := make(map[string]interface{})

	// Regular expression to match key structure, e.g. `a.b[0].c`.
	regex := regexp.MustCompile(`\.?([^.\[\]]+)|\[(\d+)\]`)

	for key, value := range flatMap {
		cur := result
		var prop string

		// Find all matches for the current key.
		matches := regex.FindAllStringSubmatch(key, -1)

		for i, match := range matches {
			// If the second match group (index) is not empty, it's an array index.
			if match[2] != "" {
				// Convert the index from string to int.
				index, _ := strconv.Atoi(match[2])

				// Initialize the current position as a slice if it doesn't exist.
				if _, ok := cur[prop]; !ok {
					cur[prop] = make([]interface{}, index+1)
				}

				// Convert current prop to a slice.
				curSlice := cur[prop].([]interface{})

				// Expand the slice if needed.
				if len(curSlice) <= index {
					newSlice := make([]interface{}, index+1)
					copy(newSlice, curSlice)
					cur[prop] = newSlice
					curSlice = newSlice
				}

				// If we're at the last key part, assign the value.
				if i == len(matches)-1 {
					curSlice[index] = value
				} else {
					// Move to the next level, initializing as necessary.
					if curSlice[index] == nil {
						curSlice[index] = make(map[string]interface{})
					}
					cur = curSlice[index].(map[string]interface{})
				}
				prop = ""
			} else {
				// Non-array key part.
				if prop != "" {
					if _, ok := cur[prop]; !ok {
						cur[prop] = make(map[string]interface{})
					}
					cur = cur[prop].(map[string]interface{})
				}
				prop = match[1]

				// If we're at the last key part, assign the value.
				if i == len(matches)-1 {
					cur[prop] = value
				}
			}
		}
	}

	return result
}

func DoTest() {
	// 示例 JSON 对象
	// jsonStr := `{"z":{},"a":{"c":1,"d":"xxx"},"b":[1,"bbb",true],"l":[{"l1":1,"l2":"xxx"}],"e":{"f":{"f":1,"g":true,"h":null,"j":{"k":"kkk"}}}}`
	jsonStr := `{"a":1,"b":null,"c":true,"d":{},"e":{"e1":1,"e2":[],"e3":["1"],"e4":{},"e5":{"e51":1,"e52":[{"x":1},"xx",3]}},"f":[],"g":[1,"xxx",true,null,[],{}],"h":[{"x":1,"y":"xxx1","z":true},{"x":2,"y":"xxx2","z":false}]}`
	// jsonStr := `{"a":1,"b":null,"c":true,"d":{},"e":{"e1":1,"e2":[],"e3":["1"],"e4":{},"e5":{"e51":1}},"f":[],"g":[1,"xxx",true,null,[],{}],"h":[{"x":1,"y":"xxx1","z":true},{"x":2,"y":"xxx2","z":false}]}`
	// jsonStr := `{"a":"a1","b":{"b1":"b11"},"c":["c1","c2","c3"],"d":[{"d1":"d2"}]}`
	fmt.Println("Json:", jsonStr)

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	fmt.Println("Origin:", data)

	// 扁平化
	flattened := Flatten(data)
	fmt.Println("Flattened:", flattened)

	// 反扁平化
	unflattened := Unflatten(flattened)
	fmt.Println("Origin:", unflattened)

	//
	byteArr, _ := json.Marshal(unflattened)
	fmt.Println("Json:", string(byteArr))

	if jsonStr == string(byteArr) {
		fmt.Println("!!!!!!!!!!!!!!!!!!!!")
	}

}
