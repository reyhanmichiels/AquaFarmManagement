package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reyhanmichiels/AquaFarmManagement/app/pond/usecase"
	"github.com/reyhanmichiels/AquaFarmManagement/domain"
	"github.com/reyhanmichiels/AquaFarmManagement/util"
)

type PondHandler struct {
	pondUsecase usecase.IPondUsecase
}

func NewPondHandler(pondUsecase usecase.IPondUsecase) *PondHandler {
	return &PondHandler{
		pondUsecase: pondUsecase,
	}
}

func (pondHandler *PondHandler) Create(c *gin.Context) {
	//bind request
	var request domain.PondBind
	err := c.ShouldBindJSON(&request)
	if err != nil {
		util.FailResponse(c, http.StatusBadRequest, "failed to bind request", err)
		return
	}

	//create pond
	pond, errObject := pondHandler.pondUsecase.Create(request)
	if errObject != nil {
		errObject := errObject.(util.ErrorObject)
		util.FailResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return
	}

	util.SuccessResponse(c, http.StatusCreated, "successfully create pond", pond)
}

func (pondHandler *PondHandler) Update(c *gin.Context) {
	//bind request
	var request domain.PondBind
	err := c.ShouldBindJSON(&request)
	if err != nil {
		util.FailResponse(c, http.StatusBadRequest, "failed to bind request", err)
		return
	}

	// bind param
	pondId := c.Param("pondId")

	//create pond
	pond, errObject := pondHandler.pondUsecase.Update(request, pondId)
	if errObject != nil {
		errObject := errObject.(util.ErrorObject)
		util.FailResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return
	}

	util.SuccessResponse(c, http.StatusCreated, "successfully create pond", pond)
}
