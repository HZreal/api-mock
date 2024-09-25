package routes

import (
	"gin-init/core/wire"
	"gin-init/middleware"
	"github.com/gin-gonic/gin"
	"time"
)

func AddUserRouter(r *gin.RouterGroup) {
	// userController := controller.UserController{}
	appController, _ := wire.InitializeApp()

	userGroup := r.Group("user")
	{
		userGroup.POST("info/all", appController.UserController.GetAllUser)
		// userGroup.GET("info/detail", userController.GetUserDetail)
		userGroup.GET("info/detail", middleware.RateLimitMiddleware(5, time.Minute), appController.UserController.GetUserDetail)
		// userGroup.POST("info/list", middleware.JWTMiddleware(), userController.GetUserList)
		userGroup.POST("info/list", middleware.JWTMiddleware(), appController.UserController.GetUserList)
		// userGroup.POST("info/create", userController.CreateUser)
		userGroup.POST("info/create", appController.UserController.CreateUser)
		userGroup.POST("info/update", appController.UserController.UpdateUser)
		userGroup.POST("info/update/passwd", appController.UserController.UpdateUserPassword)
		userGroup.POST("info/reset/passwd", appController.UserController.ResetUserPassword)
		userGroup.POST("info/delete", appController.UserController.DeleteUser)
	}
}
