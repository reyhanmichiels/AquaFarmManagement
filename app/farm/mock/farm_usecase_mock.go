package mock

import (
	"github.com/reyhanmichiels/AquaFarmManagement/domain"
	"github.com/reyhanmichiels/AquaFarmManagement/util"
	"github.com/stretchr/testify/mock"
)

type FarmUsecaseMock struct {
	Mock mock.Mock
}

func (farmUsecaseMock *FarmUsecaseMock) Create(request domain.FarmBind) (domain.Farm, any) {
	args := farmUsecaseMock.Mock.Called(request)

	if args[1] != nil {
		return domain.Farm{}, args[1].(util.ErrorObject)
	}

	return args[0].(domain.Farm), nil
}

func (farmUsecaseMock *FarmUsecaseMock) Update(request domain.FarmBind, farmId string) (domain.Farm, any) {
	args := farmUsecaseMock.Mock.Called(request)

	if args[1] != nil {
		return domain.Farm{}, args[1].(util.ErrorObject)
	}

	return args[0].(domain.Farm), nil
}

func (farmUsecaseMock *FarmUsecaseMock) Get() ([]domain.Farm, any) {
	args := farmUsecaseMock.Mock.Called()

	if args[1] != nil {
		return nil, args[1].(util.ErrorObject)
	}

	return args[0].([]domain.Farm), nil
}

func (farmUsecaseMock *FarmUsecaseMock) GetFarmById(farmId string) (domain.FarmApi, any) {
	args := farmUsecaseMock.Mock.Called(farmId)

	if args[1] != nil {
		return domain.FarmApi{}, args[1].(util.ErrorObject)
	}

	return args[0].(domain.FarmApi), nil
}
