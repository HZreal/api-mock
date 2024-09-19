package vo

/**
 * @Author elastic·H
 * @Date 2024-08-08
 * @File: user.vo.go
 * @Description:
 */

type UserDetailInfo struct {
	Id       int                    `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Username string                 `json:"name" gorm:"column:username"`
	Phone    string                 `json:"phone" gorm:"column:phone"`
	Age      int                    `json:"age" gorm:"column:age"`
	Extra    map[string]interface{} `json:"extra" gorm:"column:extra;serializer:json"`
}
