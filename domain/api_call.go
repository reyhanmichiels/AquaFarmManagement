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

type ApiCallResponse struct {
	Endpoint        string `json:"endpoint"`
	Method          string `json:"method"`
	Count           int    `json:"count"`
	UniqueUserAgent int    `json:"unique_user_agent"`
}
