package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reyhanmichiels/AquaFarmManagement/app/api_call/usecase"
	"github.com/reyhanmichiels/AquaFarmManagement/util"
)

type ApiCallHandler struct {
	apiCallUsecase usecase.IApiCallUsecase
}

func NewApiCallHandler(apiCallUsecase usecase.IApiCallUsecase) *ApiCallHandler {
	return &ApiCallHandler{
		apiCallUsecase: apiCallUsecase,
	}
}

func (apiCallHandler *ApiCallHandler) Get(c *gin.Context) {
	apiCalls, errObject := apiCallHandler.apiCallUsecase.Get()
	if errObject != nil {
		errObject := errObject.(util.ErrorObject)
		util.FailResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return
	}

	util.SuccessResponse(c, http.StatusOK, "successfully get all api call", apiCalls)
}
