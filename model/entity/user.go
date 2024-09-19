package entity

type UserModel struct {
	Id       int                    `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Username string                 `gorm:"column:username"`
	Password string                 `gorm:"column:password"`
	Phone    string                 `gorm:"column:phone"`
	Age      int                    `gorm:"column:age"`
	Status   bool                   `gorm:"column:status"`
	Extra    map[string]interface{} `gorm:"column:extra;serializer:json"`
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (t *UserModel) TableName() string {
	return "tb_user"
}
