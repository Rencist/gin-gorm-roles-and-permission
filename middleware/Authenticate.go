package middleware

import (
	"gin-gorm-clean-template/common"
	"gin-gorm-clean-template/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(jwtService service.JWTService, userService service.UserService, roleHasPermissionService service.RoleHasPermissionService, permissionService service.PermissionService, permissionRoutes string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Ditemukan", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if !strings.Contains(authHeader, "Bearer ") {
			response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if !token.Valid {
			response := common.BuildErrorResponse("Gagal Memproses Request", "Akses Ditolak", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID, err := jwtService.GetUserIDByToken(authHeader)
		if err != nil {
			response := common.BuildErrorResponse("Gagal Memproses Request", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		user, err := userService.FindUserByID(ctx, userID)
		if err != nil {
			response := common.BuildErrorResponse("Gagal Memproses Request", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		roleHasPermission, err := roleHasPermissionService.FindRoleHasPermissionByRoleID(ctx, user.RoleID)
		if err != nil {
			response := common.BuildErrorResponse("Gagal Memproses Request", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		for _, v := range roleHasPermission {
			permission, err := permissionService.FindByPermissionID(ctx, v.PermissionID)
			if err != nil {
				response := common.BuildErrorResponse("Gagal Memproses Request", err.Error(), nil)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
			if permission.Routes == permissionRoutes {
				ctx.Set("token", authHeader)
				ctx.Set("userID", userID)
				ctx.Next()	
			}
		}
		response := common.BuildErrorResponse("User Tidak Memiliki Akses ke Endpoint Ini", "Unauthorized", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
		
	}
}