package controller

import (
	"gin-init/common/response"
	"gin-init/service"
	"github.com/gin-gonic/gin"
)

type DemoController struct {
	//
	demoService *service.DemoService
}

func NewDemoController(demoService *service.DemoService) *DemoController {
	return &DemoController{demoService: demoService}
}

func (d *DemoController) SendMsgWithRabbitMQ(c *gin.Context) {
	d.demoService.SendMsgWithRabbitMQ(c)

	//
	response.SuccessWithoutData(c)
}
