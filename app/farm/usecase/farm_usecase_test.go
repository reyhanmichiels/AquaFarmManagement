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
		createFarmMock := farmRepositoryMock.Mock.On("CreateFarm", &farm).Return(nil).Run(func(args mock.Arguments) {
			arg := args[0].(*domain.Farm)
			arg.ID = "testId"
			arg.Name = parameter.Name
		})

		//call usecase
		successResponse, errorResponse := farmUsecase.Create(parameter)

		//test result
		assert.Nil(t, errorResponse, "err response should be nil")
		assert.Equal(t, parameter.Name, successResponse.Name, "name should be equal")
		assert.Equal(t, "testId", successResponse.ID, "name should be equal")

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

func TestUpdate(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		//prepare data for func parameter
		parameter := domain.FarmBind{
			Name: "testUpdateName",
		}
		farmId := "testId"

		//call mock
		farm := domain.Farm{
			ID:   farmId,
			Name: parameter.Name,
		}
		findFarmMock := farmRepositoryMock.Mock.On("FindFarmByCondition", &domain.Farm{}, "name = ?", parameter.Name).Return(errors.New("not found"))
		updateFarmMock := farmRepositoryMock.Mock.On("UpdateFarm", &farm).Return(nil).Run(func(args mock.Arguments) {
			arg := args[0].(*domain.Farm)
			arg.ID = farmId
			arg.Name = parameter.Name
		})

		successResponse, errorResponse := farmUsecase.Update(parameter, farmId)

		//test result
		assert.Nil(t, errorResponse, "err response should be nil")
		assert.Equal(t, parameter.Name, successResponse.Name, "name should be equal")
		assert.Equal(t, farmId, successResponse.ID, "name should be equal")

		findFarmMock.Unset()
		updateFarmMock.Unset()
	})

	t.Run("should return error when duplicate entry", func(t *testing.T) {
		//prepare data for func parameter
		parameter := domain.FarmBind{
			Name: "testUpdateName",
		}
		farmId := "testId"

		//call mock
		findFarmMock := farmRepositoryMock.Mock.On("FindFarmByCondition", &domain.Farm{}, "name = ?", parameter.Name).Return(nil)

		_, errorResponse := farmUsecase.Update(parameter, farmId)

		//test result
		errObjectFromResponse := errorResponse.(util.ErrorObject)
		assert.Equal(t, errors.New("farm name is already used"), errObjectFromResponse.Err, "error should be equal")
		assert.Equal(t, http.StatusConflict, errObjectFromResponse.Code, "status code should be equal")
		assert.Equal(t, "failed to update farm", errObjectFromResponse.Message, "message should be equal")

		findFarmMock.Unset()
	})

	t.Run("should return error when failed update farm", func(t *testing.T) {
		//prepare data for func parameter
		parameter := domain.FarmBind{
			Name: "testUpdateName",
		}
		farmId := "testId"

		//call mock
		farm := domain.Farm{
			ID:   farmId,
			Name: parameter.Name,
		}
		findFarmMock := farmRepositoryMock.Mock.On("FindFarmByCondition", &domain.Farm{}, "name = ?", parameter.Name).Return(errors.New("not found"))
		updateFarmMock := farmRepositoryMock.Mock.On("UpdateFarm", &farm).Return(errors.New("sql failed"))

		_, errorResponse := farmUsecase.Update(parameter, farmId)

		//test result
		errObjectFromResponse := errorResponse.(util.ErrorObject)
		assert.Equal(t, errors.New("sql failed"), errObjectFromResponse.Err, "error should be equal")
		assert.Equal(t, http.StatusInternalServerError, errObjectFromResponse.Code, "status code should be equal")
		assert.Equal(t, "failed to update farm", errObjectFromResponse.Message, "message should be equal")

		findFarmMock.Unset()
		updateFarmMock.Unset()
	})
}

func TestGet(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		//call mock
		farmsResponse := []domain.Farm{
			{
				ID:   "testID1",
				Name: "testName1",
			},
			{
				ID:   "testID2",
				Name: "testName2",
			},
			{
				ID:   "testID3",
				Name: "testName3",
			},
		}
		var farms []domain.Farm
		getFarmsMock := farmRepositoryMock.Mock.On("GetFarms", &farms).Return(nil).Run(func(args mock.Arguments) {
			arg := args[0].(*[]domain.Farm)
			*arg = append(*arg, farmsResponse...)
		})

		successResponse, errorResponse := farmUsecase.Get()

		//test result
		assert.Nil(t, errorResponse, "err response should be nil")
		for i := range successResponse {
			assert.Equal(t, farmsResponse[i].Name, successResponse[i].Name, "name should be equal")
			assert.Equal(t, farmsResponse[i].ID, successResponse[i].ID, "id should be equal")
		}

		getFarmsMock.Unset()
	})

	t.Run("should return error when sql failed", func(t *testing.T) {
		//call mock
		var farms []domain.Farm
		getFarmsMock := farmRepositoryMock.Mock.On("GetFarms", &farms).Return(errors.New("sql failed"))

		_, errorResponse := farmUsecase.Get()

		//test result
		errObjectFromResponse := errorResponse.(util.ErrorObject)
		assert.Equal(t, errors.New("sql failed"), errObjectFromResponse.Err, "error should be equal")
		assert.Equal(t, http.StatusInternalServerError, errObjectFromResponse.Code, "status code should be equal")
		assert.Equal(t, "failed to get all farm", errObjectFromResponse.Message, "message should be equal")

		getFarmsMock.Unset()
	})

	t.Run("should return error when farm is not exist", func(t *testing.T) {
		//call mock
		var farms []domain.Farm
		getFarmsMock := farmRepositoryMock.Mock.On("GetFarms", &farms).Return(nil)

		_, errorResponse := farmUsecase.Get()

		//test result
		errObjectFromResponse := errorResponse.(util.ErrorObject)
		assert.Equal(t, errors.New("farm not found"), errObjectFromResponse.Err, "error should be equal")
		assert.Equal(t, http.StatusNotFound, errObjectFromResponse.Code, "status code should be equal")
		assert.Equal(t, "failed to get all farm", errObjectFromResponse.Message, "message should be equal")

		getFarmsMock.Unset()
	})
}

func TestGetFarmById(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		//call mock
		farmResponse := domain.FarmApi{
			ID:   "testID",
			Name: "farm1",
			Ponds: []domain.Pond{
				{
					ID:     "pondID1",
					Name:   "pond1",
					FarmID: "testID",
				},
				{
					ID:     "pondID2",
					Name:   "pond2",
					FarmID: "testID",
				},
			},
		}
		var farm domain.FarmApi
		getFarmByIdMock := farmRepositoryMock.Mock.On("GetFarmById", &farm, farmResponse.ID).Return(nil).Run(func(args mock.Arguments) {
			arg := args[0].(*domain.FarmApi)
			arg.ID = farmResponse.ID
			arg.Name = farmResponse.Name
			arg.Ponds = farmResponse.Ponds
		})

		successResponse, errorResponse := farmUsecase.GetFarmById(farmResponse.ID)

		//test result
		assert.Nil(t, errorResponse, "err response should be nil")

		assert.Equal(t, farmResponse.ID, successResponse.ID, "id should be equal")
		assert.Equal(t, farmResponse.Name, successResponse.Name, "name should be equal")

		for i, v := range successResponse.Ponds {
			assert.Equal(t, v.Name, successResponse.Ponds[i].Name, "pond name should be equal")
			assert.Equal(t, v.ID, successResponse.Ponds[i].ID, "pond id should be equal")
			assert.Equal(t, v.FarmID, successResponse.Ponds[i].FarmID, "farm id should be equal")
		}

		getFarmByIdMock.Unset()
	})

	t.Run("should return error when farm is not exist", func(t *testing.T) {
		//call mock
		var farm domain.FarmApi
		getFarmByIdMock := farmRepositoryMock.Mock.On("GetFarmById", &farm, "testID").Return(errors.New("record not found"))

		_, errorResponse := farmUsecase.GetFarmById("testID")

		//test result
		errObjectFromResponse := errorResponse.(util.ErrorObject)
		assert.Equal(t, errors.New("record not found"), errObjectFromResponse.Err, "error should be equal")
		assert.Equal(t, http.StatusNotFound, errObjectFromResponse.Code, "status code should be equal")
		assert.Equal(t, "failed to get farm by id", errObjectFromResponse.Message, "message should be equal")

		getFarmByIdMock.Unset()
	})
}
