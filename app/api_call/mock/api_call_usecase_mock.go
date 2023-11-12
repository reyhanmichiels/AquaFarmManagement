package mock

import (
	"github.com/reyhanmichiels/AquaFarmManagement/util"
	"github.com/stretchr/testify/mock"
)

type ApiCallUsecaseMock struct {
	Mock mock.Mock
}

func (apiCallUsecaseMock *ApiCallUsecaseMock) Get() (map[string]map[string]int, any) {
	args := apiCallUsecaseMock.Mock.Called()

	if args[1] != nil {
		return nil, args[1].(util.ErrorObject)
	}

	return args[0].(map[string]map[string]int), nil
}
