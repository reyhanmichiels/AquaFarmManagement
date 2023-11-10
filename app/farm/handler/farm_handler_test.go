package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	farm_mock "github.com/reyhanmichiels/AquaFarmManagement/app/farm/mock"
	"github.com/reyhanmichiels/AquaFarmManagement/domain"
	"github.com/reyhanmichiels/AquaFarmManagement/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var farmUsecaseMock = farm_mock.FarmUsecaseMock{
	Mock: mock.Mock{},
}

var farmHandler = NewFarmHandler(&farmUsecaseMock)

func TestCreateFarm(t *testing.T) {
	t.Run("should create farm", func(t *testing.T) {
		//prepare data for call handler
		data := domain.CreateFarmBind{
			Name: "testName",
		}

		// call mock
		mockCallResponse := domain.Farm{
			ID:    "testID",
			Name:  data.Name,
			Ponds: nil,
		}

		mockCall := farmUsecaseMock.Mock.On("Create", data).Return(mockCallResponse, nil)

		// call handler
		engine := gin.Default()
		engine.POST("/api/farms", farmHandler.Create)

		dataAsJson, err := json.Marshal(data)
		if err != nil {
			t.Fatal(err)
		}

		response := httptest.NewRecorder()
		requestCall, err := http.NewRequest("POST", "/api/farms", bytes.NewBuffer(dataAsJson))
		if err != nil {
			t.Fatal(err.Error())
		}

		engine.ServeHTTP(response, requestCall)

		//test response
		var responseData map[string]any
		err = json.Unmarshal(response.Body.Bytes(), &responseData)
		if err != nil {
			t.Fatal(err.Error())
		}

		farmData := responseData["data"].(map[string]any)
		assert.Equal(t, http.StatusCreated, response.Code, "status code should be equal")
		assert.Equal(t, "success", responseData["status"], "status should be equal")
		assert.Equal(t, "successfully create farm", responseData["message"], "message should be equal")
		assert.Equal(t, mockCallResponse.ID, farmData["id"], "farm id should be equal")
		assert.Equal(t, mockCallResponse.Name, farmData["name"], "farm name should be equal")
		assert.Nil(t, farmData["ponds"])

		mockCall.Unset()
	})

	t.Run("should reject when request invalid", func(t *testing.T) {
		// call handler
		engine := gin.Default()
		engine.POST("/api/farms", farmHandler.Create)

		response := httptest.NewRecorder()
		requestCall, err := http.NewRequest("POST", "/api/farms", nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		engine.ServeHTTP(response, requestCall)

		//test response
		var responseData map[string]any
		err = json.Unmarshal(response.Body.Bytes(), &responseData)
		if err != nil {
			t.Fatal(err.Error())
		}

		assert.Equal(t, http.StatusBadRequest, response.Code, "status code should be equal")
		assert.Equal(t, "error", responseData["status"], "status should be equal")
		assert.Equal(t, "failed to bind input", responseData["message"], " message should be equal")
	})

	t.Run("should reject when usecase call return error", func(t *testing.T) {
		//prepare data for call handler
		data := domain.CreateFarmBind{
			Name: "testName",
		}

		// call mock
		errObject := util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "test_message",
			Err:     errors.New("test_error"),
		}

		mockCall := farmUsecaseMock.Mock.On("Create", data).Return(nil, errObject)

		// call handler
		engine := gin.Default()
		engine.POST("/api/farms", farmHandler.Create)

		dataAsJson, err := json.Marshal(data)
		if err != nil {
			t.Fatal(err)
		}

		response := httptest.NewRecorder()
		requestCall, err := http.NewRequest("POST", "/api/farms", bytes.NewBuffer(dataAsJson))
		if err != nil {
			t.Fatal(err.Error())
		}

		engine.ServeHTTP(response, requestCall)

		//test response
		var responseData map[string]any
		err = json.Unmarshal(response.Body.Bytes(), &responseData)
		if err != nil {
			t.Fatal(err.Error())
		}

		assert.Equal(t, http.StatusInternalServerError, response.Code, "status code should be equal")
		assert.Equal(t, "error", responseData["status"], "status should be equal")
		assert.Equal(t, "test_message", responseData["message"], " message should be equal")
		assert.Equal(t, "test_error", responseData["error"], " error should be equal")

		mockCall.Unset()
	})
}
