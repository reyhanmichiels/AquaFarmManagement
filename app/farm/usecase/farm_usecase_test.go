package usecase

import (
	"errors"
	"net/http"
	"testing"

	farm_mock "github.com/reyhanmichiels/AquaFarmManagement/app/farm/mock"

	"github.com/reyhanmichiels/AquaFarmManagement/domain"
	"github.com/reyhanmichiels/AquaFarmManagement/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var farmRepositoryMock = farm_mock.FarmRepositoryMock{
	Mock: mock.Mock{},
}

var farmUsecase = NewFarmUsecase(&farmRepositoryMock)

func TestCreate(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		//prepare data for func parameter
		parameter := domain.FarmBind{
			Name: "test_name",
		}

		//call mock
		farm := domain.Farm{
			Name: parameter.Name,
		}
		findFarmMock := farmRepositoryMock.Mock.On("FindFarmByCondition", &domain.Farm{}, "name = ?", parameter.Name).Return(errors.New("not found"))
		createFarmMock := farmRepositoryMock.Mock.On("CreateFarm", &farm).Return(nil)

		//call usecase
		successResponse, errorResponse := farmUsecase.Create(parameter)

		//test result
		assert.Nil(t, errorResponse, "err response should be nil")
		assert.Equal(t, parameter.Name, successResponse.Name, "name should be equal")

		findFarmMock.Unset()
		createFarmMock.Unset()
	})

	t.Run("should return error when duplicate entry", func(t *testing.T) {
		//prepare data for func parameter
		parameter := domain.FarmBind{
			Name: "test_name",
		}

		//call mock
		findFarmMock := farmRepositoryMock.Mock.On("FindFarmByCondition", &domain.Farm{}, "name = ?", parameter.Name).Return(nil)

		//call usecase
		_, errorResponse := farmUsecase.Create(parameter)

		//test result
		errObjectFromResponse := errorResponse.(util.ErrorObject)
		assert.Equal(t, errors.New("farm name is already used"), errObjectFromResponse.Err, "error should be equal")
		assert.Equal(t, http.StatusConflict, errObjectFromResponse.Code, "status code should be equal")
		assert.Equal(t, "failed to create farm", errObjectFromResponse.Message, "message should be equal")

		findFarmMock.Unset()
	})

	t.Run("should return error when failed to create farm", func(t *testing.T) {
		//prepare data for func parameter
		parameter := domain.FarmBind{
			Name: "test_name",
		}

		//call mock
		farm := domain.Farm{
			Name: parameter.Name,
		}
		findFarmMock := farmRepositoryMock.Mock.On("FindFarmByCondition", &domain.Farm{}, "name = ?", parameter.Name).Return(errors.New("not found"))
		createFarmMock := farmRepositoryMock.Mock.On("CreateFarm", &farm).Return(errors.New("sql failed"))

		//call usecase
		_, errorResponse := farmUsecase.Create(parameter)

		//test result
		errObjectFromResponse := errorResponse.(util.ErrorObject)
		assert.Equal(t, errors.New("sql failed"), errObjectFromResponse.Err, "error should be equal")
		assert.Equal(t, http.StatusInternalServerError, errObjectFromResponse.Code, "status code should be equal")
		assert.Equal(t, "failed to create farm", errObjectFromResponse.Message, "message should be equal")

		findFarmMock.Unset()
		createFarmMock.Unset()
	})
}
