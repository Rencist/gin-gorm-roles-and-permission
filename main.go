package main

import (
	"gin-gorm-clean-template/common"
	"gin-gorm-clean-template/config"
	"gin-gorm-clean-template/controller"
	"gin-gorm-clean-template/repository"
	"gin-gorm-clean-template/routes"
	"gin-gorm-clean-template/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		res := common.BuildErrorResponse("Gagal Terhubung ke Server", err.Error(), common.EmptyObj{})
		(*gin.Context).JSON((&gin.Context{}), http.StatusBadGateway, res)
		return
	}

	var (
		db *gorm.DB = config.SetupDatabaseConnection()
		
		jwtService service.JWTService = service.NewJWTService()

		roleRepository repository.RoleRepository = repository.NewRoleRepository(db)
		roleService service.RoleService = service.NewRoleService(roleRepository)

		permissionRepository repository.PermissionRepository = repository.NewPermissionRepository(db)
		permissionService service.PermissionService = service.NewPermissionService(permissionRepository)

		roleHasPermissionRepository repository.RoleHasPermissionRepository = repository.NewRoleHasPermissionRepository(db)
		roleHasPermissionService service.RoleHasPermissionService = service.NewRoleHasPermissionService(roleHasPermissionRepository)

		userRepository repository.UserRepository = repository.NewUserRepository(db)
		userService service.UserService = service.NewUserService(userRepository)
		userController controller.UserController = controller.NewUserController(userService, roleService, jwtService)
	)

	server := gin.Default()
	routes.UserRoutes(server, userController, jwtService, userService, roleHasPermissionService, permissionService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	server.Run("127.0.0.1:" + port)
}