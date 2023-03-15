package entity

type Permission struct {
	ID     int64  `gorm:"primary_key;not_null" json:"id"`
	Routes string `json:"routes"`

	Timestamp
}