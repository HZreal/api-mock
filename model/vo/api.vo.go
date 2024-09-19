package vo

/**
 * @Author elastic·H
 * @Date 2024-08-08
 * @File: user.vo.go
 * @Description:
 */

type ApiDetailInfo struct {
	Id     int                    `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Name   string                 `json:"name" gorm:"column:username"`
	Url    string                 `json:"url" gorm:"column:phone"`
	Method string                 `json:"method" gorm:"column:age"`
	Header map[string]interface{} `json:"header" gorm:"column:extra;serializer:json"`
}
