package dto

import (
	"github.com/google/uuid"
)

type UserCreateDto struct {
	Name 		string 		`json:"name" binding:"required"`
	Email 		string 		`json:"email" binding:"email"`
	NoTelp 		string 		`json:"no_telp" binding:"required"`
	Password 	string  	`json:"password" binding:"required"`

	RoleID      uint64 		`gorm:"foreignKey" json:"role_id" binding:"required"`
}

type UserUpdateDto struct {
	ID        	uuid.UUID   `gorm:"primary_key" json:"id" form:"id"`
	Name 		string 		`json:"name" form:"name"`
	Email 		string 		`json:"email" form:"email"`
	NoTelp 		string 		`json:"no_telp" form:"no_telp"`
	Password 	string  	`json:"password" form:"password"`
}

type UserLoginDTO struct {
	Email 		string 		`json:"email" form:"email" binding:"email"`
	Password 	string  	`json:"password" form:"password" binding:"required"`
}