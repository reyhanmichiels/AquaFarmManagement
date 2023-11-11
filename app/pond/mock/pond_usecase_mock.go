package mock

import (
	"github.com/reyhanmichiels/AquaFarmManagement/domain"
	"github.com/reyhanmichiels/AquaFarmManagement/util"
	"github.com/stretchr/testify/mock"
)

type PondUsecaseMock struct {
	Mock mock.Mock
}

func (pondUsecaseMock *PondUsecaseMock) Create(request domain.PondBind) (domain.Pond, any) {
	args := pondUsecaseMock.Mock.Called(request)

	if args[1] != nil {
		return domain.Pond{}, args[1].(util.ErrorObject)
	}

	return args[0].(domain.Pond), nil
}

func (pondUsecaseMock *PondUsecaseMock) Update(request domain.PondBind, pondId string) (domain.Pond, any) {
	args := pondUsecaseMock.Mock.Called(request, pondId)

	if args[1] != nil {
		return domain.Pond{}, args[1].(util.ErrorObject)
	}

	return args[0].(domain.Pond), nil
}

func (pondUsecaseMock *PondUsecaseMock) Get() ([]domain.Pond, any) {
	args := pondUsecaseMock.Mock.Called()

	if args[1] != nil {
		return []domain.Pond{}, args[1].(util.ErrorObject)
	}

	return args[0].([]domain.Pond), nil
}

func (pondUsecaseMock *PondUsecaseMock) GetPondById(pondId string) (domain.PondApi, any) {
	args := pondUsecaseMock.Mock.Called(pondId)

	if args[1] != nil {
		return domain.PondApi{}, args[1].(util.ErrorObject)
	}

	return args[0].(domain.PondApi), nil
}
