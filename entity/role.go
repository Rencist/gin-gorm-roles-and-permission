package entity

type Role struct {
	ID   int64  `gorm:"primary_key;not_null" json:"id"`
	Name string `json:"name"`

	Timestamp
}