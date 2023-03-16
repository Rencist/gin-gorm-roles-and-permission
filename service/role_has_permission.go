package service

import (
	"context"
	"gin-gorm-clean-template/entity"
	"gin-gorm-clean-template/repository"
)

type RoleHasPermissionService interface {
	FindRoleHasPermissionByRoleID(ctx context.Context, roleID uint64) ([]entity.RoleHasPermission, error)
}

type roleHasPermissionService struct {
	roleHasPermissionRepository repository.RoleHasPermissionRepository
}

func NewRoleHasPermissionService(rr repository.RoleHasPermissionRepository) RoleHasPermissionService {
	return &roleHasPermissionService{
		roleHasPermissionRepository: rr,
	}
}

func(rs *roleHasPermissionService) FindRoleHasPermissionByRoleID(ctx context.Context, roleID uint64) ([]entity.RoleHasPermission, error) {
	return rs.roleHasPermissionRepository.FindRoleHasPermissionByRoleID(ctx, roleID)
}