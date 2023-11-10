package mock

import (
	"github.com/reyhanmichiels/AquaFarmManagement/domain"
	"github.com/stretchr/testify/mock"
)

type FarmRepositoryMock struct {
	Mock mock.Mock
}

func (farmRepoMock *FarmRepositoryMock) FindFarmByCondition(farm any, condition string, value any) error {
	args := farmRepoMock.Mock.Called(farm, condition, value)

	if args[0] != nil {
		return args[0].(error)
	}

	return nil
}

func (farmRepoMock *FarmRepositoryMock) CreateFarm(farm *domain.Farm) error {
	args := farmRepoMock.Mock.Called(farm)

	if args[0] != nil {
		return args[0].(error)
	}

	return nil
}
