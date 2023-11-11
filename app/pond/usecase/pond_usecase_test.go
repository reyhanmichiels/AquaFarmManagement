package usecase

import (
	"errors"
	"net/http"
	"testing"

	farm_mock "github.com/reyhanmichiels/AquaFarmManagement/app/farm/mock"
	pond_mock "github.com/reyhanmichiels/AquaFarmManagement/app/pond/mock"
	"github.com/reyhanmichiels/AquaFarmManagement/domain"
	"github.com/reyhanmichiels/AquaFarmManagement/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var pondRepository = pond_mock.PondRepositoryMock{
	Mock: mock.Mock{},
}

var farmRepository = farm_mock.FarmRepositoryMock{
	Mock: mock.Mock{},
}

var pondUsecase = NewPondUsecase(&pondRepository, &farmRepository)

func TestCreate(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		// prepare usecase parameter
		request := domain.PondBind{
			Name:   "pondName",
			FarmID: "farmID",
		}

		// call mock
		pond := domain.Pond{
			Name:   request.Name,
			FarmID: request.FarmID,
		}

		findPondMock := pondRepository.Mock.On("FindPondByCondition", &domain.Pond{}, "name = ?", request.Name).Return(errors.New("pond is not found"))
		findFarmMock := farmRepository.Mock.On("FindFarmByCondition", &domain.Farm{}, "id = ?", request.FarmID).Return(nil)
		createPondMock := pondRepository.Mock.On("CreatePond", &pond).Return(nil).Run(func(args mock.Arguments) {
			arg := args[0].(*domain.Pond)
			arg.ID = "pondID"
		})

		// call usecase
		successResponse, errorResponse := pondUsecase.Create(request)

		//test response
		assert.Nil(t, errorResponse, "error response should be nil")
		assert.Equal(t, "pondID", successResponse.ID, "pond id should be equal")
		assert.Equal(t, request.Name, successResponse.Name, "pond name should be equal")
		assert.Equal(t, request.FarmID, successResponse.FarmID, "farm id should be equal")

		findPondMock.Unset()
		findFarmMock.Unset()
		createPondMock.Unset()
	})

	t.Run("should return error when duplicate entry", func(t *testing.T) {
		// prepare usecase parameter
		request := domain.PondBind{
			Name:   "pondName",
			FarmID: "farmID",
		}

		// call mock
		findPondMock := pondRepository.Mock.On("FindPondByCondition", &domain.Pond{}, "name = ?", request.Name).Return(nil)

		// call usecase
		_, errorResponse := pondUsecase.Create(request)

		//test response
		errObject := errorResponse.(util.ErrorObject)

		assert.Equal(t, http.StatusConflict, errObject.Code, "status code should be equal")
		assert.Equal(t, "failed to create pond", errObject.Message, "message should be equal")
		assert.Equal(t, errors.New("pond name is already used"), errObject.Err, "error should be equal")

		findPondMock.Unset()
	})

	t.Run("should return error when farm is not found", func(t *testing.T) {
		// prepare usecase parameter
		request := domain.PondBind{
			Name:   "pondName",
			FarmID: "farmID",
		}

		// call mock
		findPondMock := pondRepository.Mock.On("FindPondByCondition", &domain.Pond{}, "name = ?", request.Name).Return(errors.New("pond is not found"))
		findFarmMock := farmRepository.Mock.On("FindFarmByCondition", &domain.Farm{}, "id = ?", request.FarmID).Return(errors.New("farm is not found"))

		// call usecase
		_, errorResponse := pondUsecase.Create(request)

		//test response
		errObject := errorResponse.(util.ErrorObject)

		assert.Equal(t, http.StatusBadRequest, errObject.Code, "status code should be equal")
		assert.Equal(t, "failed to create pond", errObject.Message, "message should be equal")
		assert.Equal(t, errors.New("farm is not found"), errObject.Err, "error should be equal")

		findPondMock.Unset()
		findFarmMock.Unset()
	})

	t.Run("should return error when failed to create pond", func(t *testing.T) {
		// prepare usecase parameter
		request := domain.PondBind{
			Name:   "pondName",
			FarmID: "farmID",
		}

		// call mock
		pond := domain.Pond{
			Name:   request.Name,
			FarmID: request.FarmID,
		}

		findPondMock := pondRepository.Mock.On("FindPondByCondition", &domain.Pond{}, "name = ?", request.Name).Return(errors.New("pond is not found"))
		findFarmMock := farmRepository.Mock.On("FindFarmByCondition", &domain.Farm{}, "id = ?", request.FarmID).Return(nil)
		createPondMock := pondRepository.Mock.On("CreatePond", &pond).Return(errors.New("testError"))

		// call usecase
		_, errorResponse := pondUsecase.Create(request)

		//test response
		errObject := errorResponse.(util.ErrorObject)

		assert.Equal(t, http.StatusInternalServerError, errObject.Code, "status code should be equal")
		assert.Equal(t, "failed to create pond", errObject.Message, "message should be equal")
		assert.Equal(t, errors.New("testError"), errObject.Err, "error should be equal")

		findPondMock.Unset()
		findFarmMock.Unset()
		createPondMock.Unset()
	})
}
