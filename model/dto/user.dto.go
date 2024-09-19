package dto

type UserCreateDTO struct {
	Username string `json:"username" binding:"required,min=5,max=20,alphanum"`
	Password string `json:"password" binding:"required,min=6,max=30"`
	Phone    string `json:"phone" binding:"required,len=11,numeric"`
	Age      int    `json:"age" binding:"required,min=0,max=120"`
}

type UserUpdateDTO struct {
	Id       int    `json:"id" binding:"required"`
	Username string `json:"username" binding:"omitempty,min=5,max=20,alphanum"`
	Phone    string `json:"phone" binding:"omitempty,len=11,numeric"`
	Age      int    `json:"age" binding:"omitempty,min=0,max=120"`
}

type UsersFilterDTO struct {
	Username string `json:"username" binding:"omitempty,min=5,max=20,alphanum"`
	Phone    string `json:"phone" binding:"omitempty,len=11,numeric"`
	Age      int    `json:"age" binding:"omitempty,min=0,max=120"`
}

type UserListFilterDTO struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Age      int    `json:"age"`
}
