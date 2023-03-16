package service

import (
	"context"
	"gin-gorm-clean-template/entity"
	"gin-gorm-clean-template/repository"
)

type PermissionService interface {
	FindByPermissionID(ctx context.Context, permissionID uint64) (entity.Permission, error)
}

type permissionService struct {
	permissionRepository repository.PermissionRepository
}

func NewPermissionService(pr repository.PermissionRepository) PermissionService {
	return &permissionService{
		permissionRepository: pr,
	}
}

func(pr *permissionService) FindByPermissionID(ctx context.Context, permissionID uint64) (entity.Permission, error) {
	return pr.permissionRepository.FindByPermissionID(ctx, permissionID)
}