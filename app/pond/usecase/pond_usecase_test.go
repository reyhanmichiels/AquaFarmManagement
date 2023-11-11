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

func TestUpdate(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		// prepare usecase parameter
		request := domain.PondBind{
			Name:   "pondName",
			FarmID: "farmID",
		}

		pondId := "pondID"

		// call mock
		findPondMock := pondRepository.Mock.On("FindPondByCondition", &domain.Pond{}, "name = ?", request.Name).Return(errors.New("pond is not found"))
		findFarmMock := farmRepository.Mock.On("FindFarmByCondition", &domain.Farm{}, "id = ?", request.FarmID).Return(nil)

		var pond domain.Pond

		findPondByIdMock := pondRepository.Mock.On("FindPondByCondition", &domain.Pond{}, "id = ?", pondId).Return(nil)

		pond.ID = pondId
		pond.Name = request.Name
		pond.FarmID = request.FarmID

		updatePondMock := pondRepository.Mock.On("UpdatePond", &pond).Return(nil)

		// call usecase
		successResponse, errorResponse := pondUsecase.Update(request, pondId)

		//test response
		assert.Nil(t, errorResponse, "error response should be nil")
		assert.Equal(t, pondId, successResponse.ID, "pond id should be equal")
		assert.Equal(t, request.Name, successResponse.Name, "pond name should be equal")
		assert.Equal(t, request.FarmID, successResponse.FarmID, "farm id should be equal")

		findPondMock.Unset()
		findFarmMock.Unset()
		findPondByIdMock.Unset()
		updatePondMock.Unset()
	})

	t.Run("should return error when duplicate entry", func(t *testing.T) {
		// prepare usecase parameter
		request := domain.PondBind{
			Name:   "pondName",
			FarmID: "farmID",
		}

		pondId := "pondID"

		// call mock
		findPondMock := pondRepository.Mock.On("FindPondByCondition", &domain.Pond{}, "name = ?", request.Name).Return(nil)

		// call usecase
		_, errorResponse := pondUsecase.Update(request, pondId)
		errObject := errorResponse.(util.ErrorObject)

		//test response
		assert.Equal(t, http.StatusConflict, errObject.Code, "status code should be equal")
		assert.Equal(t, "failed to update pond", errObject.Message, "message should be equal")
		assert.Equal(t, errors.New("pond name is already used"), errObject.Err, "error should be equal")

		findPondMock.Unset()
	})

	t.Run("should return error when farm is not found", func(t *testing.T) {
		// prepare usecase parameter
		request := domain.PondBind{
			Name:   "pondName",
			FarmID: "farmID",
		}

		pondId := "pondID"

		// call mock
		findPondMock := pondRepository.Mock.On("FindPondByCondition", &domain.Pond{}, "name = ?", request.Name).Return(errors.New("pond is not found"))
		findFarmMock := farmRepository.Mock.On("FindFarmByCondition", &domain.Farm{}, "id = ?", request.FarmID).Return(errors.New("farm is not found"))

		// call usecase
		_, errorResponse := pondUsecase.Update(request, pondId)

		//test response
		errObject := errorResponse.(util.ErrorObject)

		assert.Equal(t, http.StatusBadRequest, errObject.Code, "status code should be equal")
		assert.Equal(t, "failed to update pond", errObject.Message, "message should be equal")
		assert.Equal(t, errors.New("farm is not found"), errObject.Err, "error should be equal")

		findPondMock.Unset()
		findFarmMock.Unset()
	})

	t.Run("should return error when failed to update pond", func(t *testing.T) {
		// prepare usecase parameter
		request := domain.PondBind{
			Name:   "pondName",
			FarmID: "farmID",
		}

		pondId := "pondID"

		// call mock
		findPondMock := pondRepository.Mock.On("FindPondByCondition", &domain.Pond{}, "name = ?", request.Name).Return(errors.New("pond is not found"))
		findFarmMock := farmRepository.Mock.On("FindFarmByCondition", &domain.Farm{}, "id = ?", request.FarmID).Return(nil)

		var pond domain.Pond

		findPondByIdMock := pondRepository.Mock.On("FindPondByCondition", &domain.Pond{}, "id = ?", pondId).Return(nil)

		pond.ID = pondId
		pond.Name = request.Name
		pond.FarmID = request.FarmID

		updatePondMock := pondRepository.Mock.On("UpdatePond", &pond).Return(errors.New("testError"))

		// call usecase
		_, errorResponse := pondUsecase.Update(request, pondId)

		//test response
		errObject := errorResponse.(util.ErrorObject)

		assert.Equal(t, http.StatusInternalServerError, errObject.Code, "status code should be equal")
		assert.Equal(t, "failed to update pond", errObject.Message, "message should be equal")
		assert.Equal(t, errors.New("testError"), errObject.Err, "error should be equal")

		findPondMock.Unset()
		findFarmMock.Unset()
		findPondByIdMock.Unset()
		updatePondMock.Unset()
	})
}

func TestGet(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		// call mock
		pondsResponse := []domain.Pond{
			{
				ID:     "pondID1",
				Name:   "pondName1",
				FarmID: "farmID1",
			},
			{
				ID:     "pondID2",
				Name:   "pondName2",
				FarmID: "farmID2",
			},
			{
				ID:     "pondID3",
				Name:   "pondName3",
				FarmID: "farmID3",
			},
		}
		var ponds []domain.Pond
		getPondsMock := pondRepository.Mock.On("GetPonds", &ponds).Return(nil).Run(func(args mock.Arguments) {
			arg := args[0].(*[]domain.Pond)
			*arg = append(*arg, pondsResponse...)
		})

		// call usecase
		successResponse, errorResponse := pondUsecase.Get()

		//test response
		assert.Nil(t, errorResponse, "error response should be nil")
		for i, v := range successResponse {
			assert.Equal(t, pondsResponse[i].ID, v.ID, "pond id should be equal")
			assert.Equal(t, pondsResponse[i].Name, v.Name, "pond name should be equal")
			assert.Equal(t, pondsResponse[i].FarmID, v.FarmID, "farm id should be equal")
		}

		getPondsMock.Unset()
	})

	t.Run("should return error when pond not found", func(t *testing.T) {
		// call mock
		var ponds []domain.Pond
		getPondsMock := pondRepository.Mock.On("GetPonds", &ponds).Return(nil)

		// call usecase
		_, errorResponse := pondUsecase.Get()

		//test response
		errObject := errorResponse.(util.ErrorObject)

		assert.Equal(t, http.StatusNotFound, errObject.Code, "status code should be equal")
		assert.Equal(t, "failed to get all pond", errObject.Message, "message should be equal")
		assert.Equal(t, errors.New("pond not found"), errObject.Err, "error should be equal")

		getPondsMock.Unset()
	})

	t.Run("should return error when fail get ponds", func(t *testing.T) {
		// call mock
		var ponds []domain.Pond
		getPondsMock := pondRepository.Mock.On("GetPonds", &ponds).Return(errors.New("testError"))

		// call usecase
		_, errorResponse := pondUsecase.Get()

		//test response
		errObject := errorResponse.(util.ErrorObject)

		assert.Equal(t, http.StatusInternalServerError, errObject.Code, "status code should be equal")
		assert.Equal(t, "failed to get all pond", errObject.Message, "message should be equal")
		assert.Equal(t, errors.New("testError"), errObject.Err, "error should be equal")
		getPondsMock.Unset()
	})
}
