package domain

import (
	"time"
)

// Model for Api Call entity
type ApiCall struct {
	Endpoint  string    `json:"endpoint" gorm:"type:varchar(100); not null;"`
	Method    string    `json:"method" gorm:"type:varchar(20); not null"`
	IpAdress  string    `json:"ip_adress" gorm:"type:varchar(100); not null;"`
	CreatedAt time.Time `json:"created_at"`
}
