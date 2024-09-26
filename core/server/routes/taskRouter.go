package routes

import (
	"gin-init/core/wire"
	"github.com/gin-gonic/gin"
)

func AddTaskRouter(r *gin.RouterGroup) {
	appController, _ := wire.InitializeApp()

	taskGroup := r.Group("task")
	{
		taskGroup.POST("info/create", appController.TaskController.Create)
		// taskGroup.POST("info/all", appController.TaskController.GetAll)
		// taskGroup.GET("info/detail", appController.TaskController.GetDetail)
		taskGroup.POST("info/list", appController.TaskController.GetList)
		// taskGroup.POST("info/update", appController.TaskController.Updater)
		// taskGroup.POST("info/delete", appController.TaskController.Deleter)
		taskGroup.POST("action/start", appController.TaskController.Start)
		taskGroup.POST("action/createStart", appController.TaskController.CreateStart)
	}
}
