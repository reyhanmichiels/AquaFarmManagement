package mock

import (
	"github.com/reyhanmichiels/AquaFarmManagement/domain"
	"github.com/stretchr/testify/mock"
)

type PondRepositoryMock struct {
	Mock mock.Mock
}

func (pondRepositoryMock *PondRepositoryMock) FindPondByCondition(pond any, condition string, value any) error {
	args := pondRepositoryMock.Mock.Called(pond, condition, value)

	if args[0] != nil {
		return args[0].(error)
	}

	return nil
}

func (pondRepositoryMock *PondRepositoryMock) CreatePond(pond *domain.Pond) error {
	args := pondRepositoryMock.Mock.Called(pond)

	if args[0] != nil {
		return args[0].(error)
	}

	return nil
}

func (pondRepositoryMock *PondRepositoryMock) UpdatePond(pond *domain.Pond) error {
	args := pondRepositoryMock.Mock.Called(pond)

	if args[0] != nil {
		return args[0].(error)
	}

	return nil
}

func (pondRepositoryMock *PondRepositoryMock) GetPonds(ponds *[]domain.Pond) error {
	args := pondRepositoryMock.Mock.Called(ponds)

	if args[0] != nil {
		return args[0].(error)
	}

	return nil
}

func (pondRepositoryMock *PondRepositoryMock) GetPondById(pond *domain.PondApi, pondId string) error {
	args := pondRepositoryMock.Mock.Called(pond, pondId)

	if args[0] != nil {
		return args[0].(error)
	}

	return nil
}
