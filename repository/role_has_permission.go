package repository

import (
	"context"
	"gin-gorm-clean-template/entity"

	"gorm.io/gorm"
)

type RoleHasPermissionRepository interface {
	GetAllRoleHasPermission(ctx context.Context) ([]entity.RoleHasPermission, error)
	FindRoleHasPermissionByRoleID(ctx context.Context, roleID uint64) ([]entity.RoleHasPermission, error)
	CreateRoleHasPermission(ctx context.Context, roleID uint64, permissionID uint64) (error)
	DeleteRoleHasPermission(ctx context.Context, roleHasPermissionID uint64) (error)
}

type roleHasPermissionConnection struct {
	connection *gorm.DB
}

func NewRoleHasPermissionRepository(db *gorm.DB) RoleHasPermissionRepository {
	return &roleHasPermissionConnection{
		connection: db,
	}
}

func(db *roleHasPermissionConnection) FindRoleHasPermissionByRoleID(ctx context.Context, roleID uint64) ([]entity.RoleHasPermission, error){
	var roleHasPermission []entity.RoleHasPermission
	rx := db.connection.Where("role_id = ?", roleID).Find(&roleHasPermission)
	if rx.Error != nil {
		return roleHasPermission, rx.Error
	}
	return roleHasPermission, nil
}

func(db *roleHasPermissionConnection) GetAllRoleHasPermission(ctx context.Context) ([]entity.RoleHasPermission, error) {
	var roleHasPermission []entity.RoleHasPermission
	rx := db.connection.Find(&roleHasPermission)
	if rx.Error != nil {
		return roleHasPermission, rx.Error
	}
	return roleHasPermission, nil
}

func(db *roleHasPermissionConnection) CreateRoleHasPermission(ctx context.Context, roleID uint64, permissionID uint64) (error) {
	var roleHasPermission = entity.RoleHasPermission{
		ID: 1,
		RoleID: roleID,
		PermissionID: permissionID,
	}
	px := db.connection.Create(&roleHasPermission)
	if px.Error != nil {
		return px.Error
	}
	return nil
}

func(db *roleHasPermissionConnection) DeleteRoleHasPermission(ctx context.Context, roleHasPermissionID uint64) (error) {
	px := db.connection.Delete(&entity.RoleHasPermission{}, &roleHasPermissionID)
	if px.Error != nil {
		return px.Error
	}
	return nil
}