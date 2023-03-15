package dto

type RoleHasPermissionCreateDto struct {
	ID int64 `gorm:"primary_key" json:"id" form:"id"`

	RoleID       string `gorm:"foreignKey" json:"role_id" form:"role_id" binding:"required"`
	PermissionID string `gorm:"foreignKey" json:"permission_id" form:"permission_id" binding:"required"`
}

type RoleHasPermissionUpdateDto struct {
	ID int64 `gorm:"primary_key" json:"id" form:"id"`

	RoleID       string `gorm:"foreignKey" json:"role_id" form:"role_id"`
	PermissionID string `gorm:"foreignKey" json:"permission_id" form:"permission_id"`
}