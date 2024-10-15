package entity

type ParamStruct struct {
	Name    string      `json:"name"`
	Type    string      `json:"type"`
	Mock    string      `json:"mock"`
	Example interface{} `json:"example"`
}

type ApiModel struct {
	Id                  uint           `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Name                string         `gorm:"column:name"`
	Method              string         `gorm:"column:method"`
	UriArgs             string         `gorm:"column:uri_args"`
	Uri                 string         `gorm:"column:uri"`
	ContentType         string         `gorm:"column:content_type"`
	ResponseContentType string         `gorm:"column:response_content_type"`
	Args                string         `gorm:"column:args"`
	ExtraArgs           []*ParamStruct `gorm:"column:extra_args;serializer:json"`
	BodyType            uint           `gorm:"column:body_type"`
	RequestBody         string         `gorm:"column:request_body"`
	Params              []*ParamStruct `gorm:"column:params;serializer:json"`
}

func NewApiModel() *ApiModel {
	return &ApiModel{}
}

func (t *ApiModel) TableName() string {
	return "tb_api"
}
