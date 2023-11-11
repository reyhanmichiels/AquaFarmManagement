package repository

import (
	"github.com/reyhanmichiels/AquaFarmManagement/domain"
	"gorm.io/gorm"
)

type IPondRepository interface {
	FindPondByCondition(pond any, condition string, value any) error
	CreatePond(pond *domain.Pond) error
	UpdatePond(pond *domain.Pond) error
}

type PondRepository struct {
	db *gorm.DB
}

func NewPondRepository(db *gorm.DB) IPondRepository {
	return &PondRepository{
		db: db,
	}
}

func (pondRepository *PondRepository) FindPondByCondition(pond any, condition string, value any) error {
	err := pondRepository.db.Model(&domain.Pond{}).First(pond, condition, value).Error
	return err
}

func (pondRepository *PondRepository) CreatePond(pond *domain.Pond) error {
	tx := pondRepository.db.Begin()

	err := tx.Create(pond).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (pondRepository *PondRepository) UpdatePond(pond *domain.Pond) error {
	tx := pondRepository.db.Begin()

	err := tx.Save(pond).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
