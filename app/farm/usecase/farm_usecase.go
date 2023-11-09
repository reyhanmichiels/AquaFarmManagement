package usecase

import (
	"net/http"

	"github.com/reyhanmichiels/AquaFarmManagement/app/farm/repository"
	"github.com/reyhanmichiels/AquaFarmManagement/domain"
	"github.com/reyhanmichiels/AquaFarmManagement/util"
)

type IFarmUsecase interface {
	Create(request domain.CreateFarmBind) (domain.Farm, any)
}

type FarmUsecase struct {
	farmRepository repository.IFarmRepository
}

func NewFarmUsecase(farmRepository repository.IFarmRepository) IFarmUsecase {
	return &FarmUsecase{
		farmRepository: farmRepository,
	}
}

func (farmUsecase *FarmUsecase) Create(request domain.CreateFarmBind) (domain.Farm, any) {
	// check for duplicate entry
	isFarmExist := farmUsecase.farmRepository.FindFarmByCondition(domain.Farm{}, "name = ?", request.Name)
	if isFarmExist == nil {
		return domain.Farm{}, util.ErrorObject{
			Code:    http.StatusConflict,
			Err:     isFarmExist,
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
