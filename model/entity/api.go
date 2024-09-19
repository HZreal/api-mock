package entity

type ApiModel struct {
	Id     int                    `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Name   string                 `gorm:"column:name"`
	Url    string                 `gorm:"column:url"`
	Method string                 `gorm:"column:method"`
	Header map[string]interface{} `gorm:"column:header;serializer:json"`
}

func NewApiModel() *ApiModel {
	return &ApiModel{}
}

func (t *ApiModel) TableName() string {
	return "tb_api"
}
