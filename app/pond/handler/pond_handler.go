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

	util.SuccessResponse(c, http.StatusCreated, "successfully update pond", pond)
}

func (pondHandler *PondHandler) Get(c *gin.Context) {
	// get ponds
	ponds, errObject := pondHandler.pondUsecase.Get()
	if errObject != nil {
		errObject := errObject.(util.ErrorObject)
		util.FailResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return
	}

	util.SuccessResponse(c, http.StatusCreated, "successfully get all pond", ponds)
}

func (pondHandler *PondHandler) GetPondById(c *gin.Context) {
	// bind param
	pondId := c.Param("pondId")

	// get pond by id
	pond, errObject := pondHandler.pondUsecase.GetPondById(pondId)
	if errObject != nil {
		errObject := errObject.(util.ErrorObject)
		util.FailResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return
	}

	util.SuccessResponse(c, http.StatusOK, "successfully get pond by id", pond)
}
