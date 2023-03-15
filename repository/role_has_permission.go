package repository

import (
	"context"
	"gin-gorm-clean-template/entity"

	"gorm.io/gorm"
)

type RoleHasPermissionRepository interface {
	FindRoleHasPermissionByRoleID(ctx context.Context, roleID uint64) (entity.RoleHasPermission, error)
}

type roleHasPermissionConnection struct {
	connection *gorm.DB
}

func NewRoleHasPermissionRepository(db *gorm.DB) RoleHasPermissionRepository {
	return &roleHasPermissionConnection{
		connection: db,
	}
}

func(db *roleHasPermissionConnection) FindRoleHasPermissionByRoleID(ctx context.Context, roleID uint64) (entity.RoleHasPermission, error){
	var roleHasPermission entity.RoleHasPermission
	rx := db.connection.Where("role_id = ?").Find(&roleHasPermission)
	if rx.Error != nil {
		return roleHasPermission, rx.Error
	}
	return roleHasPermission, nil
}
