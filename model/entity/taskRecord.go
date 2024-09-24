package entity

type TaskRecordModel struct {
	Id       int                    `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	TaskId   string                 `gorm:"column:task_id"`
	ApiId    string                 `gorm:"column:api_id"`
	Request  map[string]interface{} `gorm:"column:Request;serializer:json"`
	Response map[string]interface{} `gorm:"column:response;serializer:json"`
	Status   string                 `gorm:"column:success_count"`

	// TODO
}

func NewTaskRecordModel() *TaskRecordModel {
	return &TaskRecordModel{}
}

func (t *TaskRecordModel) TableName() string {
	return "tb_task"
}
