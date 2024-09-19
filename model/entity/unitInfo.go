package entity

import (
	"time"
)

// UnitInfo 生成 gorm 结构体
type UnitInfo struct {
	Id          uint                   `json:"id" gorm:"column:id;primary_key;comment:'主键id'"`
	Name        string                 `json:"name" gorm:"column:name;comment:'名称'"`
	Code        string                 `json:"code" gorm:"column:code;NOT NULL;comment:'编号'"`
	Alias       string                 `json:"alias" gorm:"column:alias;comment:'别名'"`
	ExternalNo  string                 `json:"externalNo" gorm:"column:external_no;comment:'外部编号'"`
	Type        string                 `json:"type" gorm:"column:type;NOT NULL;comment:'类型'"`
	Tag         string                 `json:"tag" gorm:"column:tag;comment:'标签'"`
	BasicInfo   map[string]interface{} `json:"basicInfo" gorm:"column:basic_info;type:jsonb;comment:'基本信息';serializer:json"`
	DynamicInfo map[string]interface{} `json:"dynamicInfo" gorm:"column:dynamic_info;type:jsonb;comment:'其他动态属性';serializer:json"`
	Extra       map[string]interface{} `json:"extra" gorm:"column:extra;type:jsonb;comment:'扩展信息（预留存储其他信息）';serializer:json"`
	ParentId    string                 `json:"parentId" gorm:"column:parent_id;comment:'父级编号'"`
	BirthTime   time.Time              `json:"birthTime" gorm:"column:birth_time;comment:'外平台设备创建时间'"`
	CreateTime  time.Time              `json:"createTime" gorm:"column:create_time;default:now();NOT NULL;comment:'创建时间'"`
	UpdateTime  time.Time              `json:"updateTime" gorm:"column:update_time;default:now();NOT NULL;comment:'更新时间'"`
}

func (t *UnitInfo) TableName() string {
	return "tb_unit"
}
