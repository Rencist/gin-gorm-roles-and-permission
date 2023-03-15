package service

import (
	"context"
	"gin-gorm-clean-template/entity"
	"gin-gorm-clean-template/repository"
)

type RoleService interface {
	FindRoleByID(ctx context.Context, roleID uint64) (entity.Role, error)
}

type roleService struct {
	roleRepository repository.RoleRepository
}

func NewRoleService(rr repository.RoleRepository) RoleService {
	return &roleService{
		roleRepository: rr,
	}
}

func(rs *roleService) FindRoleByID(ctx context.Context, roleID uint64) (entity.Role, error) {
	return rs.roleRepository.FindRoleByID(ctx, roleID)
}