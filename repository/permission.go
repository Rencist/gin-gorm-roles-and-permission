package repository

import (
	"context"
	"gin-gorm-clean-template/entity"

	"gorm.io/gorm"
)

type PermissionRepository interface {
	FindByPermissionID(ctx context.Context, permissionID uint64) (entity.Permission, error)
	GetAllPermission(ctx context.Context) ([]entity.Permission, error)
	CreatePermission(ctx context.Context, permission entity.Permission) (entity.Permission, error)
	UpdatePermission(ctx context.Context, permission entity.Permission) (error)
}

type permissionConnection struct {
	connection *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionConnection{
		connection: db,
	}
}

func(db *permissionConnection) FindByPermissionID(ctx context.Context, permissionID uint64) (entity.Permission, error) {
	var permission entity.Permission
	px := db.connection.Where("id = ?", permissionID).Take(&permission)
	if px.Error != nil {
		return permission, px.Error
	}
	return permission, nil
}

func(db *permissionConnection) GetAllPermission(ctx context.Context) ([]entity.Permission, error) {
	var permissionList []entity.Permission
	px := db.connection.Find(&permissionList)
	if px.Error != nil {
		return permissionList, px.Error
	}
	return permissionList, nil
}

func(db *permissionConnection) CreatePermission(ctx context.Context, permission entity.Permission) (entity.Permission, error) {
	px := db.connection.Create(&permission)
	if px.Error != nil {
		return permission, px.Error
	}
	return permission, nil
}

func(db *permissionConnection) UpdatePermission(ctx context.Context, permission entity.Permission) (error) {
	px := db.connection.Updates(&permission)
	if px.Error != nil {
		return px.Error
	}
	return nil
}