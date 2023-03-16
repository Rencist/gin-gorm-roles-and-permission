package routes

import (
	"gin-gorm-clean-template/controller"
	"gin-gorm-clean-template/middleware"
	"gin-gorm-clean-template/service"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, UserController controller.UserController, jwtService service.JWTService, userService service.UserService, roleHasPermissionService service.RoleHasPermissionService) {
	userRoutes := router.Group("/api/user")
	{
		userRoutes.POST("", UserController.RegisterUser)
		userRoutes.GET("", middleware.Authenticate(jwtService, userService, roleHasPermissionService, "user.index"), UserController.GetAllUser)
		userRoutes.POST("/login", UserController.LoginUser)
		userRoutes.DELETE("/", middleware.Authenticate(jwtService, userService, roleHasPermissionService, "user.delete"), UserController.DeleteUser)
		userRoutes.PUT("/", middleware.Authenticate(jwtService, userService, roleHasPermissionService, "user.update"), UserController.UpdateUser)
		userRoutes.GET("/me", middleware.Authenticate(jwtService, userService, roleHasPermissionService, "user_me.index"), UserController.MeUser)
	}
}