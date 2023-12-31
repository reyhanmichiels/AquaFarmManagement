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
	var request domain.FarmBind
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

func (farmHandler *FarmHandler) Update(c *gin.Context) {
	//bind and validate data
	var request domain.FarmBind
	err := c.ShouldBindJSON(&request)
	if err != nil {
		util.FailResponse(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	farmId := c.Param("farmId")

	//update farm
	farm, errObject := farmHandler.farmUsecase.Update(request, farmId)
	if errObject != nil {
		errObject := errObject.(util.ErrorObject)
		util.FailResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return
	}

	util.SuccessResponse(c, http.StatusOK, "successfully update farm", farm)
}

func (farmHandler *FarmHandler) Get(c *gin.Context) {
	//get farms
	farms, errObject := farmHandler.farmUsecase.Get()
	if errObject != nil {
		errObject := errObject.(util.ErrorObject)
		util.FailResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return
	}

	util.SuccessResponse(c, http.StatusOK, "successfully get all farm", farms)
}

func (farmHandler *FarmHandler) GetFarmById(c *gin.Context) {
	//bind param
	farmId := c.Param("farmId")

	//get farm by id
	farm, errObject := farmHandler.farmUsecase.GetFarmById(farmId)
	if errObject != nil {
		errObject := errObject.(util.ErrorObject)
		util.FailResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return
	}

	util.SuccessResponse(c, http.StatusOK, "successfully get farm by id", farm)
}

func (farmHandler *FarmHandler) Delete(c *gin.Context) {
	//bind param
	farmId := c.Param("farmId")

	//delete farm
	errObject := farmHandler.farmUsecase.Delete(farmId)
	if errObject != nil {
		errObject := errObject.(util.ErrorObject)
		util.FailResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return
	}

	util.SuccessResponse(c, http.StatusOK, "successfully delete farm", nil)
}
