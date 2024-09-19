package vo

/**
 * @Author elastic·H
 * @Date 2024-08-08
 * @File: sys.vo.go
 * @Description:
 */

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type SysLoginVo struct {
	Token string `json:"token" gorm:"serializer:json"`
}

// 可以自定义 JsonObject 类型，但必须实现 Value 和 Scan 方法
type JsonObject map[string]interface{}

// MarshalJSON 实现 MarshalJSON 接口，为了将自定义类型序列化为 JSON 字符串。这是为了让 gin 框架能够正确地将 JSONB 类型数据转换为 JSON 格式返回给客户端
func (j JsonObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}(j))
}

// UnmarshalJSON 实现 UnmarshalJSON 接口，为了将 JSON 字符串反序列化为自定义类型。这是为了能够从 JSON 数据中正确地解析出 JSONB 类型的字段。
func (j *JsonObject) UnmarshalJSON(data []byte) error {
	var tmp map[string]interface{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*j = tmp
	return nil
}

// Value 将自定义类型转换为数据库可接受的格式。对于 jsonb 类型，需要将 JSONB 类型转换为字节数组，以便 gorm 可以将其存储在数据库中
func (j JsonObject) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan 实现这个方法是为了将数据库中的数据转换为自定义类型。当从数据库中读取 jsonb 字段时，需要将其转换回 JSONB 类型
func (j *JsonObject) Scan(value interface{}) error {
	if value == nil {
		*j = make(map[string]interface{})
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return j.UnmarshalJSON(v)
	case string:
		return j.UnmarshalJSON([]byte(v))
	default:
		return errors.New("unsupported type for JsonObject")
	}
}
