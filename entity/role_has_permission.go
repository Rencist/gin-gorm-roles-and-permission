package entity

type RoleHasPermission struct {
	ID uint64 `gorm:"primary_key;not_null" json:"id"`

	RoleID       uint64      `gorm:"foreignKey" json:"role_id"`
	Role         *Role       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role,omitempty"`
	PermissionID uint64      `gorm:"foreignKey" json:"permission_id"`
	Permission   *Permission `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"permission,omitempty"`

	Timestamp
}