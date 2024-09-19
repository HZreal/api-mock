package controller

import (
	"gin-init/common/response"
	"gin-init/model/dto"
	"gin-init/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SysController struct {
	//
	sysService *service.SysService
}

func NewSysController(sysService *service.SysService) *SysController {
	return &SysController{sysService: sysService}
}

func (uC *SysController) Login(c *gin.Context) {
	//
	var loginData dto.LoginData
	if err := c.ShouldBindJSON(&loginData); err != nil {
		response.Failed(c, response.ParamsError)
		return
	}

	//
	data, err := uC.sysService.Login(c, loginData)
	if err != nil {
		response.Failed(c, response.UnKnownError)
		return
	}

	//
	response.SuccessWithData(c, data)

}

func (uC *SysController) Logout(c *gin.Context) {
	//
	c.JSON(http.StatusOK, gin.H{})

}
