package repository

import (
	"github.com/reyhanmichiels/AquaFarmManagement/domain"
	"gorm.io/gorm"
)

type IApiCallRepository interface {
	GetApiCalls(apiCalls *[]domain.ApiCallResponse) error
}

type ApiCallRepository struct {
	db *gorm.DB
}

func NewApiCallRepository(db *gorm.DB) IApiCallRepository {
	return &ApiCallRepository{
		db: db,
	}
}

func (apiCallRepository *ApiCallRepository) GetApiCalls(apiCalls *[]domain.ApiCallResponse) error {
	query := `
	select endpoint, method, count(ip_adress) as count, count(distinct ip_adress) as unique_user_agent 
	from api_calls 
	GROUP BY endpoint, method
	`

	err := apiCallRepository.db.Model(&domain.ApiCall{}).Raw(query).Scan(apiCalls).Error
	return err
}
