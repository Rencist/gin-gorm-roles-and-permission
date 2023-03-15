package entity

type Role struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`

	Timestamp
}