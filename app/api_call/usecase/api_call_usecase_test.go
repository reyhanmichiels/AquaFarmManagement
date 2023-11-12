package usecase

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	api_call_repo_mock "github.com/reyhanmichiels/AquaFarmManagement/app/api_call/mock"
	"github.com/reyhanmichiels/AquaFarmManagement/domain"
	"github.com/reyhanmichiels/AquaFarmManagement/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var apiCallRepositoryMock = api_call_repo_mock.ApiCallRepositoryMock{
	Mock: mock.Mock{},
}

var apiCallUsecase = NewApiCallUsecase(&apiCallRepositoryMock)

func TestGet(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		// call mock
		apiCallResponse := []domain.ApiCallResponse{
			{
				Endpoint:        "endpoint1",
				Method:          "method1",
				Count:           1,
				UniqueUserAgent: 1,
			},
			{
				Endpoint:        "endpoint2",
				Method:          "method2",
				Count:           1,
				UniqueUserAgent: 2,
			},
		}

		var apiCalls []domain.ApiCallResponse
		getApiCallsMock := apiCallRepositoryMock.Mock.On("GetApiCalls", &apiCalls).Return(nil).Run(func(args mock.Arguments) {
			arg := args[0].(*[]domain.ApiCallResponse)
			*arg = append(*arg, apiCallResponse...)
		})

		// call usecase
		successResponse, errorResponse := apiCallUsecase.Get()

		//test response
		assert.Nil(t, errorResponse, "error response should be nil")

		expectedResponse := make(map[string]map[string]int, 0)
		for _, v := range apiCallResponse {
			endpointAndMethod := fmt.Sprintf("%s %s", v.Method, v.Endpoint)
			expectedResponse[endpointAndMethod] = map[string]int{
				"count":             v.Count,
				"unique_user_agent": v.UniqueUserAgent,
			}
		}

		assert.Equal(t, expectedResponse["endpoint1 method1"], successResponse["endpoint1 method1"], "endpoint1 method1 should be equal")
		assert.Equal(t, expectedResponse["endpoint2 method2"], successResponse["endpoint2 method2"], "endpoint2 method2 should be equal")

		getApiCallsMock.Unset()
	})

	t.Run("should return error when api call is not found", func(t *testing.T) {
		// call mock
		var apiCalls []domain.ApiCallResponse
		getApiCallsMock := apiCallRepositoryMock.Mock.On("GetApiCalls", &apiCalls).Return(nil)

		// call usecase
		_, errorResponse := apiCallUsecase.Get()

		//test response
		errObject := errorResponse.(util.ErrorObject)

		assert.Equal(t, http.StatusNotFound, errObject.Code, "status code should be equal")
		assert.Equal(t, errors.New("api call not found"), errObject.Err, "error should be equal")
		assert.Equal(t, "failed to get api calls", errObject.Message, "message should be equal")

		getApiCallsMock.Unset()
	})

	t.Run("should return error when failed to get api calls", func(t *testing.T) {
		// call mock
		var apiCalls []domain.ApiCallResponse
		getApiCallsMock := apiCallRepositoryMock.Mock.On("GetApiCalls", &apiCalls).Return(errors.New("testError"))

		// call usecase
		_, errorResponse := apiCallUsecase.Get()

		//test response
		errObject := errorResponse.(util.ErrorObject)

		assert.Equal(t, http.StatusInternalServerError, errObject.Code, "status code should be equal")
		assert.Equal(t, errors.New("testError"), errObject.Err, "error should be equal")
		assert.Equal(t, "failed to get api calls", errObject.Message, "message should be equal")

		getApiCallsMock.Unset()
	})
}
