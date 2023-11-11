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
