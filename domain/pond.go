package domain

import "time"

// Model for Pond entity
type Pond struct {
	ID        string    `json:"id" gorm:"type:uuid;not null; primary key"`
	FarmID    string    `json:"farm_id"`
	Farm      Farm      `json:"-"`
	Name      string    `json:"name" gorm:"type:varchar(100); not null; unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
