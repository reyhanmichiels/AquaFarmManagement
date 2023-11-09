package domain

import "time"

// Model for Farm entity
type Farm struct {
	ID        string    `json:"id" gorm:"type:uuid; not null; primary key"`
	Ponds     []Pond    `json:"ponds"`
	Name      string    `json:"name" gorm:"type:varchar(100); not null; unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
