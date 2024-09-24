package controller

import (
	"gin-init/common/response"
	"gin-init/model/dto"
	"gin-init/service"
	"github.com/gin-gonic/gin"
)

type ApiController struct {
	//
	ApiService *service.ApiService
}

func NewApiController(apiService *service.ApiService) *ApiController {
	return &ApiController{ApiService: apiService}
}

func (uC *ApiController) Create(c *gin.Context) {
	var body dto.ApiCreateDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response.Failed(c, response.ParamsError)
		return
	}

	//
	data := uC.ApiService.Create(c, body)

	response.SuccessWithData(c, data)
}

func (uC *ApiController) Import(c *gin.Context) {
	//
	uC.ApiService.Import()

	response.SuccessWithoutData(c)
}

// func (uC *ApiController) GetAll(c *gin.Context) {
// 	//
// 	var body dto.UsersFilterDTO
//
// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		response.Failed(c, response.ParamsError)
// 		return
// 	}
//
// 	// 调用服务层
// 	data := uC.ApiService.GetAll(c, body)
//
// 	//
// 	response.SuccessWithData(c, data)
// }
//
// func (uC *ApiController) GetList(c *gin.Context) {
//
// 	var query dto.QueryPagination
// 	if err := c.ShouldBindQuery(&query); err != nil {
// 		response.Failed(c, response.ParamsError)
// 		return
// 	}
//
// 	// var body dto.UserListFilterDTO
// 	var body map[string]interface{}
// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		response.Failed(c, response.ParamsError)
// 		return
// 	}
//
// 	//
// 	data := uC.ApiService.GetUserList(c, query, body)
//
// 	response.SuccessWithData(c, data)
// }
//
// func (uC *ApiController) GetDetail(c *gin.Context) {
// 	//
// 	var body dto.QueryId
//
// 	if err := c.ShouldBindQuery(&body); err != nil {
// 		response.Failed(c, response.ParamsError)
// 		return
// 	}
//
// 	// 调用服务层
// 	data := uC.ApiService.GetUserDetail(c, body.Id)
//
// 	//
// 	response.SuccessWithData(c, data)
// }
//
// func (uC *ApiController) Update(c *gin.Context) {
// 	var body dto.UserUpdateDTO
//
// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		response.Failed(c, response.ParamsError)
// 		return
// 	}
//
// 	//
// 	user := uC.ApiService.UpdateUser(c, body)
//
// 	response.SuccessWithData(c, user)
//
// }
//
// func (uC *ApiController) Delete(c *gin.Context) {
// 	var body dto.BodyJsonId
//
// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		response.Failed(c, response.ParamsError)
// 		return
// 	}
//
// 	//
// 	uC.ApiService.DeleteUser(c, body.Id)
//
// 	response.SuccessWithoutData(c)
// }
