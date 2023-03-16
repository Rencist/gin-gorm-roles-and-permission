package repository

import (
	"context"
	"gin-gorm-clean-template/entity"

	"gorm.io/gorm"
)

type PermissionRepository interface {
	FindByPermissionID(ctx context.Context, permissionID uint64) (entity.Permission, error)
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