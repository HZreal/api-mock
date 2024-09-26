package controller

import (
	"gin-init/common/response"
	"gin-init/model/dto"
	"gin-init/service"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	//
	TaskService *service.TaskService
}

func NewTaskController(taskService *service.TaskService) *TaskController {
	return &TaskController{TaskService: taskService}
}

// Create
func (ctl *TaskController) Create(c *gin.Context) {
	var body dto.TaskCreateDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response.Failed(c, response.ParamsError)
		return
	}

	//
	data := ctl.TaskService.Create(c, body)

	response.SuccessWithData(c, data)
}

// Start
func (ctl *TaskController) Start(c *gin.Context) {

	//
	service.Run()

	response.SuccessWithoutData(c)
}

// CreateStart
func (ctl *TaskController) CreateStart(c *gin.Context) {

	//
	response.SuccessWithoutData(c)
}

// GetList
func (ctl *TaskController) GetList(c *gin.Context) {

	var query dto.QueryPagination
	if err := c.ShouldBindQuery(&query); err != nil {
		response.Failed(c, response.ParamsError)
		return
	}

	// var body dto.UserListFilterDTO
	var body map[string]interface{}
	if err := c.ShouldBindJSON(&body); err != nil {
		response.Failed(c, response.ParamsError)
		return
	}

	//
	data := ctl.TaskService.GetList(c, query, body)

	response.SuccessWithData(c, data)
}

//	func (ctl *TaskController) GetAll(c *gin.Context) {
//		//
//		var body dto.UsersFilterDTO
//
//		if err := c.ShouldBindJSON(&body); err != nil {
//			response.Failed(c, response.ParamsError)
//			return
//		}
//
//		// 调用服务层
//		data := ctl.TaskService.GetAll(c, body)
//
//		//
//		response.SuccessWithData(c, data)
//	}

// func (ctl *TaskController) GetDetail(c *gin.Context) {
// 	//
// 	var body dto.QueryId
//
// 	if err := c.ShouldBindQuery(&body); err != nil {
// 		response.Failed(c, response.ParamsError)
// 		return
// 	}
//
// 	// 调用服务层
// 	data := ctl.TaskService.GetUserDetail(c, body.Id)
//
// 	//
// 	response.SuccessWithData(c, data)
// }
//
// func (ctl *TaskController) Update(c *gin.Context) {
// 	var body dto.UserUpdateDTO
//
// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		response.Failed(c, response.ParamsError)
// 		return
// 	}
//
// 	//
// 	user := ctl.TaskService.UpdateUser(c, body)
//
// 	response.SuccessWithData(c, user)
//
// }
//
// func (ctl *TaskController) Delete(c *gin.Context) {
// 	var body dto.BodyJsonId
//
// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		response.Failed(c, response.ParamsError)
// 		return
// 	}
//
// 	//
// 	ctl.TaskService.DeleteUser(c, body.Id)
//
// 	response.SuccessWithoutData(c)
// }
