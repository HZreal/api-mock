package routes

import (
	"gin-init/core/wire"
	"github.com/gin-gonic/gin"
)

func AddTaskRecordRouter(r *gin.RouterGroup) {
	appController, _ := wire.InitializeApp()

	taskGroup := r.Group("taskRecord")
	{
		taskGroup.POST("info/create", appController.TaskRecordController.Create)
		// taskGroup.POST("info/all", appController.TaskRecordController.GetAllUser)
		// taskGroup.GET("info/detail", appController.TaskRecordController.GetUserDetail)
		taskGroup.POST("info/list", appController.TaskRecordController.GetList)
		// taskGroup.POST("info/update", appController.TaskRecordController.UpdateUser)
		// taskGroup.POST("info/delete", appController.TaskRecordController.DeleteUser)
	}
}
