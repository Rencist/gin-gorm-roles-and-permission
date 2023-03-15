package entity

import "github.com/google/uuid"

type RoleHasPermission struct {
	ID   			uint64  		`gorm:"primary_key;not_null" json:"id"`
	Name 			string 		`json:"name"`

	RoleID 			uuid.UUID 	`gorm:"foreignKey" json:"role_id"`
	Role   			*Role     	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role,omitempty"`
	PermissionID  	uuid.UUID 	`gorm:"foreignKey" json:"permission_id"`
	Permission    	*Permission `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"permission,omitempty"`

	Timestamp
}