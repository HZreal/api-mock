package entity

type TaskModel struct {
	Id           int                    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;NOT NULL" json:"id"`
	Name         string                 `gorm:"column:name;type:varchar(16);comment:名称" json:"name"`
	TotalCount   uint32                 `gorm:"column:total_count;type:int(11) unsigned;default:0;comment:总数量;NOT NULL" json:"total_count"`
	SuccessCount uint32                 `gorm:"column:success_count;type:int(11) unsigned;default:0;comment:成功数量;NOT NULL" json:"success_count"`
	Current      uint64                 `gorm:"column:current;type:int(11) unsigned;comment:当前" json:"current"`
	Settings     map[string]interface{} `gorm:"column:settings;type:json;comment:设置" json:"settings"`
	Progress     uint8                  `gorm:"column:progress;type:tinyint(4) unsigned;comment:进度 0 - 100" json:"progress"`
	Status       uint8                  `gorm:"column:status;type:tinyint(4) unsigned;default:0;comment:状态：0 - 待处理；1 - 进行中；2 - 已完成；3 - 失败;NOT NULL" json:"status"`
}

func NewTaskModel() *TaskModel {
	return &TaskModel{}
}

func (t *TaskModel) TableName() string {
	return "tb_task"
}
