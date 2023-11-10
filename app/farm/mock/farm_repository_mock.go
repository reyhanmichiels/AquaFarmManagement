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

func (farmRepoMock *FarmRepositoryMock) UpdateFarm(farm *domain.Farm) error {
	args := farmRepoMock.Mock.Called(farm)

	if args[0] != nil {
		return args[0].(error)
	}

	return nil
}

func (farmRepoMock *FarmRepositoryMock) GetFarms(farms *[]domain.Farm) error {
	args := farmRepoMock.Mock.Called(farms)

	if args[0] != nil {
		return args[0].(error)
	}

	return nil
}

func (farmRepoMock *FarmRepositoryMock) GetFarmById(farm *domain.FarmApi, farmId string) error {
	args := farmRepoMock.Mock.Called(farm, farmId)

	if args[0] != nil {
		return args[0].(error)
	}

	return nil
}
