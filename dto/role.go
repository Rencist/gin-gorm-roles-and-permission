package dto

type RoleCreateDto struct {
	ID   int64  `gorm:"primary_key" json:"id" form:"id"`
	Name string `json:"name" form:"name" binding:"required"`
}

type RoleUpdateDto struct {
	ID   int64  `gorm:"primary_key" json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}