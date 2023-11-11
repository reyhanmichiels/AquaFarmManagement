package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Model for Pond entity
type Pond struct {
	ID        string         `json:"id" gorm:"type:uuid;not null; primary key"`
	FarmID    string         `json:"farm_id" gorm:"type:uuid;not null"`
	Farm      Farm           `json:"-"`
	Name      string         `json:"name" gorm:"type:varchar(100); not null; unique"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// Automate generate uuid when create farm
func (pond *Pond) BeforeCreate(tx *gorm.DB) error {
	pond.ID = uuid.NewString()
	return nil
}

type PondBind struct {
	Name   string `json:"name" binding:"required,max=100,min=4"`
	FarmID string `json:"farm_id" binding:"required"`
}

type PondApi struct {
	ID        string    `json:"id"`
	FarmID    string    `json:"farm_id"`
	Farm      Farm      `json:"farm"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
