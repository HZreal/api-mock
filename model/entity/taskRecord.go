package entity

type TaskRecordModel struct {
	Id       uint64                 `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	TaskId   uint64                 `gorm:"column:task_id"`
	ApiId    uint64                 `gorm:"column:api_id"`
	params   map[string]interface{} `gorm:"column:params;serializer:json"`
	Response map[string]interface{} `gorm:"column:response;serializer:json"`
	Status   uint8                  `gorm:"column:status"`
	Cost     uint16                 `gorm:"column:cost"`

	// TODO
}

func NewTaskRecordModel() *TaskRecordModel {
	return &TaskRecordModel{}
}

func (t *TaskRecordModel) TableName() string {
	return "tb_task_record"
}
