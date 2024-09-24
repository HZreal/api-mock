package entity

type TaskModel struct {
	Id           int    `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Name         string `gorm:"column:name"`
	TotalCount   string `gorm:"column:total_count"`
	SuccessCount string `gorm:"column:success_count"`

	// TODO
	Uri         string `gorm:"column:uri"`
	ContentType string `gorm:"column:content_type"`
	Args        string `gorm:"column:args"`
}

func NewTaskModel() *TaskModel {
	return &TaskModel{}
}

func (t *TaskModel) TableName() string {
	return "tb_task"
}
