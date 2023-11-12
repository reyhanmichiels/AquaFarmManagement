package usecase

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/reyhanmichiels/AquaFarmManagement/app/api_call/repository"
	"github.com/reyhanmichiels/AquaFarmManagement/domain"
	"github.com/reyhanmichiels/AquaFarmManagement/util"
)

type IApiCallUsecase interface {
	Get() (map[string]map[string]int, any)
}

type ApiCallUsecase struct {
	apiCallRepository repository.IApiCallRepository
}

func NewApiCallUsecase(apiCallRepository repository.IApiCallRepository) IApiCallUsecase {
	return &ApiCallUsecase{
		apiCallRepository: apiCallRepository,
	}
}

func (apiCallUsecase *ApiCallUsecase) Get() (map[string]map[string]int, any) {
	var apiCalls []domain.ApiCallResponse
	err := apiCallUsecase.apiCallRepository.GetApiCalls(&apiCalls)
	if err != nil {
		return map[string]map[string]int{}, util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Err:     err,
			Message: "failed to get api calls",
		}
	}

	if len(apiCalls) == 0 {
		return map[string]map[string]int{}, util.ErrorObject{
			Code:    http.StatusNotFound,
			Err:     errors.New("api call not found"),
			Message: "failed to get api calls",
		}
	}

	apiResponse := make(map[string]map[string]int, 0)
	for _, v := range apiCalls {
		endpointAndMethod := fmt.Sprintf("%s %s", v.Method, v.Endpoint)
		apiResponse[endpointAndMethod] = map[string]int{
			"count":             v.Count,
			"unique_user_agent": v.UniqueUserAgent,
		}
	}

	return apiResponse, nil
}
