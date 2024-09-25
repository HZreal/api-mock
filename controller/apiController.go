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

func (ctl *ApiController) Create(c *gin.Context) {
	var body dto.ApiCreateDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response.Failed(c, response.ParamsError)
		return
	}

	//
	data := ctl.ApiService.Create(c, body)

	response.SuccessWithData(c, data)
}

func (ctl *ApiController) Import(c *gin.Context) {
	//
	ctl.ApiService.Import()

	response.SuccessWithoutData(c)
}

//	func (ctl *ApiController) GetAll(c *gin.Context) {
//		//
//		var body dto.UsersFilterDTO
//
//		if err := c.ShouldBindJSON(&body); err != nil {
//			response.Failed(c, response.ParamsError)
//			return
//		}
//
//		// 调用服务层
//		data := ctl.ApiService.GetAll(c, body)
//
//		//
//		response.SuccessWithData(c, data)
//	}

// GetList
func (ctl *ApiController) GetList(c *gin.Context) {

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
	data := ctl.ApiService.GetList(c, query, body)

	response.SuccessWithData(c, data)
}

// func (ctl *ApiController) GetDetail(c *gin.Context) {
// 	//
// 	var body dto.QueryId
//
// 	if err := c.ShouldBindQuery(&body); err != nil {
// 		response.Failed(c, response.ParamsError)
// 		return
// 	}
//
// 	// 调用服务层
// 	data := ctl.ApiService.GetUserDetail(c, body.Id)
//
// 	//
// 	response.SuccessWithData(c, data)
// }
//
// func (ctl *ApiController) Update(c *gin.Context) {
// 	var body dto.UserUpdateDTO
//
// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		response.Failed(c, response.ParamsError)
// 		return
// 	}
//
// 	//
// 	user := ctl.ApiService.UpdateUser(c, body)
//
// 	response.SuccessWithData(c, user)
//
// }
//
// func (ctl *ApiController) Delete(c *gin.Context) {
// 	var body dto.BodyJsonId
//
// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		response.Failed(c, response.ParamsError)
// 		return
// 	}
//
// 	//
// 	ctl.ApiService.DeleteUser(c, body.Id)
//
// 	response.SuccessWithoutData(c)
// }
