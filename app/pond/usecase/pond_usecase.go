package usecase

import (
	"errors"
	"net/http"

	farm_repository "github.com/reyhanmichiels/AquaFarmManagement/app/farm/repository"
	pond_repository "github.com/reyhanmichiels/AquaFarmManagement/app/pond/repository"
	"github.com/reyhanmichiels/AquaFarmManagement/domain"
	"github.com/reyhanmichiels/AquaFarmManagement/util"
)

type IPondUsecase interface {
	Create(request domain.PondBind) (domain.Pond, any)
	Update(request domain.PondBind, pondId string) (domain.Pond, any)
	Get() ([]domain.Pond, any)
	GetPondById(pondId string) (domain.PondApi, any)
	Delete(pondId string) any
}

type PondUsecase struct {
	pondRepository pond_repository.IPondRepository
	farmRepository farm_repository.IFarmRepository
}

func NewPondUsecase(pondRepository pond_repository.IPondRepository, farmRepository farm_repository.IFarmRepository) IPondUsecase {
	return &PondUsecase{
		pondRepository: pondRepository,
		farmRepository: farmRepository,
	}
}

func (pondUsecase *PondUsecase) Create(request domain.PondBind) (domain.Pond, any) {
	// check for duplicate entry
	isPondExist := pondUsecase.pondRepository.FindPondByCondition(&domain.Pond{}, "name = ?", request.Name)
	if isPondExist == nil {
		return domain.Pond{}, util.ErrorObject{
			Code:    http.StatusConflict,
			Err:     errors.New("pond name is already used"),
			Message: "failed to create pond",
		}
	}

	//check if farm exist
	isFarmExist := pondUsecase.farmRepository.FindFarmByCondition(&domain.Farm{}, "id = ?", request.FarmID)
	if isFarmExist != nil {
		return domain.Pond{}, util.ErrorObject{
			Code:    http.StatusBadRequest,
			Err:     errors.New("farm is not found"),
			Message: "failed to create pond",
		}
	}

	// create pond
	pond := domain.Pond{
		Name:   request.Name,
		FarmID: request.FarmID,
	}
	err := pondUsecase.pondRepository.CreatePond(&pond)
	if err != nil {
		return domain.Pond{}, util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Err:     err,
			Message: "failed to create pond",
		}
	}

	return pond, nil
}

func (pondUsecase *PondUsecase) Update(request domain.PondBind, pondId string) (domain.Pond, any) {
	// check for duplicate entry
	isPondExist := pondUsecase.pondRepository.FindPondByCondition(&domain.Pond{}, "name = ?", request.Name)
	if isPondExist == nil {
		return domain.Pond{}, util.ErrorObject{
			Code:    http.StatusConflict,
			Err:     errors.New("pond name is already used"),
			Message: "failed to update pond",
		}
	}

	// check if farm exist
	isFarmExist := pondUsecase.farmRepository.FindFarmByCondition(&domain.Farm{}, "id = ?", request.FarmID)
	if isFarmExist != nil {
		return domain.Pond{}, util.ErrorObject{
			Code:    http.StatusBadRequest,
			Err:     errors.New("farm is not found"),
			Message: "failed to update pond",
		}
	}

	// check if pond exist
	var pond domain.Pond
	pondUsecase.pondRepository.FindPondByCondition(&pond, "id = ?", pondId)

	pond.ID = pondId
	pond.Name = request.Name
	pond.FarmID = request.FarmID

	// update pond
	err := pondUsecase.pondRepository.UpdatePond(&pond)
	if err != nil {
		return domain.Pond{}, util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Err:     err,
			Message: "failed to update pond",
		}
	}
	return pond, nil
}

func (pondUsecase *PondUsecase) Get() ([]domain.Pond, any) {
	// get ponds
	var ponds []domain.Pond
	err := pondUsecase.pondRepository.GetPonds(&ponds)
	if err != nil {
		return []domain.Pond{}, util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Err:     err,
			Message: "failed to get all pond",
		}
	}

	// check if pond exist
	if len(ponds) == 0 {
		return []domain.Pond{}, util.ErrorObject{
			Code:    http.StatusNotFound,
			Err:     errors.New("pond not found"),
			Message: "failed to get all pond",
		}
	}

	return ponds, nil
}

func (pondUsecase *PondUsecase) GetPondById(pondId string) (domain.PondApi, any) {
	// get ponds
	var pond domain.PondApi
	isPondExist := pondUsecase.pondRepository.GetPondById(&pond, pondId)

	// check if pond exist
	if isPondExist != nil {
		return domain.PondApi{}, util.ErrorObject{
			Code:    http.StatusNotFound,
			Err:     errors.New("pond not found"),
			Message: "failed to get pond by id",
		}
	}

	return pond, nil
}

func (pondUsecase *PondUsecase) Delete(pondId string) any {
	var pond domain.Pond
	// check if pond exist
	isPondExist := pondUsecase.pondRepository.FindPondByCondition(&pond, "id = ?", pondId)
	if isPondExist != nil {
		return util.ErrorObject{
			Code:    http.StatusNotFound,
			Err:     errors.New("pond not found"),
			Message: "failed to delete pond",
		}
	}

	//delete pond
	err := pondUsecase.pondRepository.DeletePond(&pond)
	if err != nil {
		return util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Err:     err,
			Message: "failed to delete pond",
		}
	}

	return nil
}
