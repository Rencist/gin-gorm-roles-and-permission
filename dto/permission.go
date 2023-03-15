package dto

type PermissionCreateDto struct {
	ID   int64  `gorm:"primary_key" json:"id" form:"id"`
	Routes string `json:"routes" form:"routes" binding:"required"`
}

type PermissionUpdateDto struct {
	ID   int64  `gorm:"primary_key" json:"id" form:"id"`
	Routes string `json:"routes" form:"routes"`
}