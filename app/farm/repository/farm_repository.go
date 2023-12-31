package repository

import (
	"github.com/reyhanmichiels/AquaFarmManagement/domain"
	"gorm.io/gorm"
)

type IFarmRepository interface {
	FindFarmByCondition(farm any, condition string, value any) error
	CreateFarm(farm *domain.Farm) error
	UpdateFarm(farm *domain.Farm) error
	GetFarms(farms *[]domain.Farm) error
	GetFarmById(farm *domain.FarmApi, farmId string) error
	DeleteFarm(farm *domain.Farm) error
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
	err := farmRepo.db.Model(&domain.Farm{}).First(farm, condition, value).Error
	return err
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

func (farmRepo *FarmRepository) UpdateFarm(farm *domain.Farm) error {
	tx := farmRepo.db.Begin()

	err := tx.Save(farm).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (farmRepo *FarmRepository) GetFarms(farms *[]domain.Farm) error {
	err := farmRepo.db.Find(farms).Error
	return err
}

func (farmRepo *FarmRepository) GetFarmById(farm *domain.FarmApi, farmId string) error {
	err := farmRepo.db.Model(&domain.Farm{}).Preload("Ponds").First(farm, "id = ?", farmId).Error
	return err
}

func (farmRepo *FarmRepository) DeleteFarm(farm *domain.Farm) error {
	tx := farmRepo.db.Begin()

	var ponds []domain.Pond
	tx.Model(&domain.Pond{}).Find(&ponds, "farm_id = ?", farm.ID)
	if len(ponds) != 0 {
		for _, pond := range ponds {
			err := tx.Delete(&pond).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	err := tx.Delete(farm).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
