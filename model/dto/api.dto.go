package dto

type ApiCreateDTO struct {
	Name   string                 `json:"name" binding:"required,min=1,max=20"`
	Url    string                 `json:"url" binding:"required,max=64"`
	Method string                 `json:"method" binding:"required"`
	Header map[string]interface{} `json:"header" binding:"required"`
}

// TODO

type ApiUpdateDTO struct {
	Id       int    `json:"id" binding:"required"`
	Username string `json:"username" binding:"omitempty,min=5,max=20,alphanum"`
	Phone    string `json:"phone" binding:"omitempty,len=11,numeric"`
	Age      int    `json:"age" binding:"omitempty,min=0,max=120"`
}

type ApisFilterDTO struct {
	Username string `json:"username" binding:"omitempty,min=5,max=20,alphanum"`
	Phone    string `json:"phone" binding:"omitempty,len=11,numeric"`
	Age      int    `json:"age" binding:"omitempty,min=0,max=120"`
}

type ApiListFilterDTO struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Age      int    `json:"age"`
}
