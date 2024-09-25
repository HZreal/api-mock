package entity

type TaskModel struct {
	Id           int    `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Name         string `gorm:"column:name"`
	TotalCount   uint32 `gorm:"column:total_count"`
	SuccessCount uint32 `gorm:"column:success_count"`
}

func NewTaskModel() *TaskModel {
	return &TaskModel{}
}

func (t *TaskModel) TableName() string {
	return "tb_task"
}
