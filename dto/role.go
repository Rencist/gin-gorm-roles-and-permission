package dto

type RoleCreateDto struct {
	ID   uint64 `gorm:"primary_key" json:"id"`
	Name string `json:"name" binding:"required"`
}

type RoleUpdateDto struct {
	ID   uint64 `gorm:"primary_key" json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}