package repository

import (
	"context"
	"gin-gorm-clean-template/entity"

	"gorm.io/gorm"
)

type RoleRepository interface {
	FindRoleByID(ctx context.Context, roleID uint64) (entity.Role, error)
}

type roleConnection struct {
	connection *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleConnection{
		connection: db,
	}
}

func(db *roleConnection) FindRoleByID(ctx context.Context, roleID uint64) (entity.Role, error) {
	var role entity.Role
	rx := db.connection.Where("id = ?", roleID).Take(&role)
	if rx.Error != nil {
		return role, rx.Error
	}
	return role, nil
}