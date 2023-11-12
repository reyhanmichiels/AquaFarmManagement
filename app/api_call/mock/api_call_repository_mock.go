package mock

import (
	"github.com/reyhanmichiels/AquaFarmManagement/domain"
	"github.com/stretchr/testify/mock"
)

type ApiCallRepositoryMock struct {
	Mock mock.Mock
}

func (apiCallRepositoryMock *ApiCallRepositoryMock) GetApiCalls(apiCalls *[]domain.ApiCallResponse) error {
	args := apiCallRepositoryMock.Mock.Called(apiCalls)

	if args[0] != nil {
		return args[0].(error)
	}

	return nil
}
