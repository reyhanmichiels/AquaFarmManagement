package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Model for Farm entity
type Farm struct {
	ID        string    `json:"id" gorm:"type:uuid; not null; primary key"`
	Ponds     []Pond    `json:"-"`
	Name      string    `json:"name" gorm:"type:varchar(100); not null; unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Automate generate uuid when create farm
func (farm *Farm) BeforeCreate(tx *gorm.DB) error {
	farm.ID = uuid.NewString()
	return nil
}

type CreateFarmBind struct {
	Name string `json:"name" binding:"required,max=100,min=4"`
}
