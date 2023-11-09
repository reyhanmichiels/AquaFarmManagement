package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reyhanmichiels/AquaFarmManagement/app/farm/usecase"
	"github.com/reyhanmichiels/AquaFarmManagement/domain"
	"github.com/reyhanmichiels/AquaFarmManagement/util"
)

type FarmHandler struct {
	farmUsecase usecase.IFarmUsecase
}

func NewFarmHandler(farmUsecase usecase.IFarmUsecase) *FarmHandler {
	return &FarmHandler{
		farmUsecase: farmUsecase,
	}
}

func (farmHandler *FarmHandler) Create(c *gin.Context) {
	//bind and validate data
	var request domain.CreateFarmBind
	err := c.ShouldBindJSON(&request)
	if err != nil {
		util.FailResponse(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	//create new farm
	farm, errObject := farmHandler.farmUsecase.Create(request)
	if errObject != nil {
		errObject := errObject.(util.ErrorObject)
		util.FailResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return
	}

	util.SuccessResponse(c, http.StatusCreated, "successfully create farm", farm)
}
