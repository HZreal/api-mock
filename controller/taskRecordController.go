package controller

import (
	"gin-init/common/response"
	"gin-init/model/dto"
	"gin-init/service"
	"github.com/gin-gonic/gin"
)

type TaskRecordController struct {
	//
	TaskRecordService *service.TaskRecordService
}

func NewTaskRecordController(taskRecordService *service.TaskRecordService) *TaskRecordController {
	return &TaskRecordController{TaskRecordService: taskRecordService}
}

// Create
func (ctl *TaskRecordController) Create(c *gin.Context) {
	var body dto.TaskRecordCreateDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response.Failed(c, response.ParamsError)
		return
	}

	//
	data := ctl.TaskRecordService.Create(c, body)

	response.SuccessWithData(c, data)
}

// GetList
func (ctl *TaskRecordController) GetList(c *gin.Context) {

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
	data := ctl.TaskRecordService.GetList(c, query, body)

	response.SuccessWithData(c, data)
}

//	func (ctl *TaskRecordController) GetAll(c *gin.Context) {
//		//
//		var body dto.UsersFilterDTO
//
//		if err := c.ShouldBindJSON(&body); err != nil {
//			response.Failed(c, response.ParamsError)
//			return
//		}
//
//		// 调用服务层
//		data := ctl.TaskRecordService.GetAll(c, body)
//
//		//
//		response.SuccessWithData(c, data)
//	}

// func (ctl *TaskRecordController) GetDetail(c *gin.Context) {
// 	//
// 	var body dto.QueryId
//
// 	if err := c.ShouldBindQuery(&body); err != nil {
// 		response.Failed(c, response.ParamsError)
// 		return
// 	}
//
// 	// 调用服务层
// 	data := ctl.TaskRecordService.GetUserDetail(c, body.Id)
//
// 	//
// 	response.SuccessWithData(c, data)
// }
//
// func (ctl *TaskRecordController) Update(c *gin.Context) {
// 	var body dto.UserUpdateDTO
//
// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		response.Failed(c, response.ParamsError)
// 		return
// 	}
//
// 	//
// 	user := ctl.TaskRecordService.UpdateUser(c, body)
//
// 	response.SuccessWithData(c, user)
//
// }
//
// func (ctl *TaskRecordController) Delete(c *gin.Context) {
// 	var body dto.BodyJsonId
//
// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		response.Failed(c, response.ParamsError)
// 		return
// 	}
//
// 	//
// 	ctl.TaskRecordService.DeleteUser(c, body.Id)
//
// 	response.SuccessWithoutData(c)
// }
