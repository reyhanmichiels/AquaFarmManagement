package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Model for Farm entity
type Farm struct {
	ID        string         `json:"id" gorm:"type:uuid; not null; primary key"`
	Ponds     []Pond         `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name      string         `json:"name" gorm:"type:varchar(100); not null; unique"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// Automate generate uuid when create farm
func (farm *Farm) BeforeCreate(tx *gorm.DB) error {
	farm.ID = uuid.NewString()
	return nil
}

type FarmBind struct {
	Name string `json:"name" binding:"required,max=100,min=4"`
}

type FarmApi struct {
	ID        string    `json:"id"`
	Ponds     []Pond    `json:"ponds" gorm:"foreignKey:FarmID"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
