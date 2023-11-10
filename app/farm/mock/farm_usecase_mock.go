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
