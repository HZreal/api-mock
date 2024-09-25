package entity

type TaskRecordModel struct {
	Id       uint64                 `gorm:"column:id;type:int(11);AUTO_INCREMENT;primary_key;NOT NULL" json:"id"`
	TaskId   uint64                 `gorm:"column:task_id;type:int(11);comment:任务 id;NOT NULL" json:"task_id"`
	ApiId    uint64                 `gorm:"column:api_id;type:int(11);comment:接口 id;NOT NULL" json:"api_id"`
	Params   map[string]interface{} `gorm:"column:params;type:json;comment:参数;serializer:json" json:"params"`
	Response map[string]interface{} `gorm:"column:response;type:json;comment:结果;serializer:json" json:"response"`
	Status   uint8                  `gorm:"column:status;type:tinyint(4) unsigned;default:0;comment:状态：0 - 待处理；1 - 进行中；2 - 已完成；3 - 失败;NOT NULL" json:"status"`
	Cost     uint16                 `gorm:"column:cost;type:mediumint(9) unsigned;comment:时间消耗 ms" json:"cost"`
	// TODO
}

func NewTaskRecordModel() *TaskRecordModel {
	return &TaskRecordModel{}
}

func (t *TaskRecordModel) TableName() string {
	return "tb_task_record"
}
