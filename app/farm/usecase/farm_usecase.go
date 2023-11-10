package usecase

import (
	"errors"
	"net/http"

	"github.com/reyhanmichiels/AquaFarmManagement/app/farm/repository"
	"github.com/reyhanmichiels/AquaFarmManagement/domain"
	"github.com/reyhanmichiels/AquaFarmManagement/util"
)

type IFarmUsecase interface {
	Create(request domain.FarmBind) (domain.Farm, any)
	Update(request domain.FarmBind, farmId string) (domain.Farm, any)
	Get() ([]domain.Farm, any)
}

type FarmUsecase struct {
	farmRepository repository.IFarmRepository
}

func NewFarmUsecase(farmRepository repository.IFarmRepository) IFarmUsecase {
	return &FarmUsecase{
		farmRepository: farmRepository,
	}
}

func (farmUsecase *FarmUsecase) Create(request domain.FarmBind) (domain.Farm, any) {
	// check for duplicate entry
	isFarmExist := farmUsecase.farmRepository.FindFarmByCondition(&domain.Farm{}, "name = ?", request.Name)
	if isFarmExist == nil {
		return domain.Farm{}, util.ErrorObject{
			Code:    http.StatusConflict,
			Err:     errors.New("farm name is already used"),
			Message: "failed to create farm",
		}
	}

	// create new farm
	farm := domain.Farm{
		Name: request.Name,
	}
	err := farmUsecase.farmRepository.CreateFarm(&farm)
	if err != nil {
		return domain.Farm{}, util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Err:     err,
			Message: "failed to create farm",
		}
	}

	return farm, nil
}

func (farmUsecase *FarmUsecase) Update(request domain.FarmBind, farmId string) (domain.Farm, any) {
	// check for duplicate entry
	isFarmExist := farmUsecase.farmRepository.FindFarmByCondition(&domain.Farm{}, "name = ?", request.Name)
	if isFarmExist == nil {
		return domain.Farm{}, util.ErrorObject{
			Code:    http.StatusConflict,
			Err:     errors.New("farm name is already used"),
			Message: "failed to update farm",
		}
	}

	// update farm
	farm := domain.Farm{
		ID:   farmId,
		Name: request.Name,
	}
	err := farmUsecase.farmRepository.UpdateFarm(&farm)
	if err != nil {
		return domain.Farm{}, util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Err:     err,
			Message: "failed to update farm",
		}
	}

	return farm, nil
}

func (farmUsecase *FarmUsecase) Get() ([]domain.Farm, any) {
	var farms []domain.Farm
	err := farmUsecase.farmRepository.GetFarms(&farms)
	if err != nil {
		return nil, util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Err:     err,
			Message: "failed to get all farm",
		}
	}

	if len(farms) == 0 {
		return nil, util.ErrorObject{
			Code:    http.StatusNotFound,
			Err:     errors.New("farm not found"),
			Message: "failed to get all farm",
		}
	}

	return farms, nil
}
