package repository

import (
	"github.com/reyhanmichiels/AquaFarmManagement/domain"
	"gorm.io/gorm"
)

type IFarmRepository interface {
	FindFarmByCondition(farm any, condition string, value any) error
	CreateFarm(farm *domain.Farm) error
}

type FarmRepository struct {
	db *gorm.DB
}

func NewFarmRepository(db *gorm.DB) IFarmRepository {
	return &FarmRepository{
		db: db,
	}
}

func (farmRepo *FarmRepository) FindFarmByCondition(farm any, condition string, value any) error {
	err := farmRepo.db.Model(domain.Farm{}).First(farm, condition, value).Error

	if err != nil {
		return err
	}

	return nil
}

func (farmRepo *FarmRepository) CreateFarm(farm *domain.Farm) error {
	tx := farmRepo.db.Begin()

	err := tx.Create(farm).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
