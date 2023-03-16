package repository

import (
	"context"
	"gin-gorm-clean-template/entity"

	"gorm.io/gorm"
)

type RoleRepository interface {
	FindRoleByID(ctx context.Context, roleID uint64) (entity.Role, error)
	GetAllROle(ctx context.Context) ([]entity.Role, error)
	CreateRole(ctx context.Context, role entity.Role) (entity.Role, error)
	UpdateRole(ctx context.Context, role entity.User) (error)
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

func(db *roleConnection) GetAllROle(ctx context.Context) ([]entity.Role, error) {
	var roleList []entity.Role
	rx := db.connection.Find(&roleList)
	if rx.Error != nil {
		return roleList, rx.Error
	}
	return roleList, nil
}

func(db *roleConnection) CreateRole(ctx context.Context, role entity.Role) (entity.Role, error) {
	rx := db.connection.Create(&role)
	if rx.Error != nil {
		return entity.Role{}, rx.Error
	}
	return role, nil
}

func(db *roleConnection) UpdateRole(ctx context.Context, role entity.User) (error) {
	rx := db.connection.Updates(&role)
	if rx.Error != nil {
		return rx.Error
	}
	return nil
}