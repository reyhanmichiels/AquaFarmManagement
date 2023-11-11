package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	pond_mock "github.com/reyhanmichiels/AquaFarmManagement/app/pond/mock"
	"github.com/reyhanmichiels/AquaFarmManagement/domain"
	"github.com/reyhanmichiels/AquaFarmManagement/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var pondUsecaseMock = pond_mock.PondUsecaseMock{
	Mock: mock.Mock{},
}

var pondHandler = NewPondHandler(&pondUsecaseMock)

func TestCreate(t *testing.T) {
	t.Run("should can create pond", func(t *testing.T) {
		// prepare request body
		requestBody := domain.PondBind{
			Name:   "pondName",
			FarmID: "farmID",
		}

		requestBodyJson, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatal(err)
		}

		// call mock
		mockResponse := domain.Pond{
			ID:     "pondID",
			Name:   requestBody.Name,
			FarmID: requestBody.FarmID,
		}
		mockCall := pondUsecaseMock.Mock.On("Create", requestBody).Return(mockResponse, nil)

		// call handler
		engine := gin.Default()
		engine.POST("/api/ponds", pondHandler.Create)

		response := httptest.NewRecorder()
		request, err := http.NewRequest("POST", "/api/ponds", bytes.NewBuffer(requestBodyJson))
		if err != nil {
			t.Fatal(err)
		}

		engine.ServeHTTP(response, request)

		// parsing response body
		var responseBody map[string]any
		err = json.Unmarshal(response.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatal(err)
		}

		// test response
		assert.Equal(t, http.StatusCreated, response.Code, "status code should be equal")
		assert.Equal(t, "success", responseBody["status"], "status should be equal")
		assert.Equal(t, "successfully create pond", responseBody["message"], "message should be equal")

		pondData := responseBody["data"].(map[string]any)

		assert.Equal(t, mockResponse.ID, pondData["id"], "pond id should be equal")
		assert.Equal(t, mockResponse.Name, pondData["name"], "pond name should be equal")
		assert.Equal(t, mockResponse.FarmID, pondData["farm_id"], "farm id should be equal")

		mockCall.Unset()
	})

	t.Run("should reject when request invalid", func(t *testing.T) {
		// call handler
		engine := gin.Default()
		engine.POST("/api/ponds", pondHandler.Create)

		response := httptest.NewRecorder()
		request, err := http.NewRequest("POST", "/api/ponds", nil)
		if err != nil {
			t.Fatal(err)
		}

		engine.ServeHTTP(response, request)

		// parsing response body
		var responseBody map[string]any
		err = json.Unmarshal(response.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatal(err)
		}

		// test response
		assert.Equal(t, http.StatusBadRequest, response.Code, "status code should be equal")
		assert.Equal(t, "error", responseBody["status"], "status should be equal")
		assert.Equal(t, "failed to bind request", responseBody["message"], "message should be equal")
	})

	t.Run("should reject when usecase call return error", func(t *testing.T) {
		// prepare request body
		requestBody := domain.PondBind{
			Name:   "pondName",
			FarmID: "farmID",
		}

		requestBodyJson, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatal(err)
		}

		// call mock
		errObject := util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Err:     errors.New("testError"),
			Message: "test message",
		}
		mockCall := pondUsecaseMock.Mock.On("Create", requestBody).Return(nil, errObject)

		// call handler
		engine := gin.Default()
		engine.POST("/api/ponds", pondHandler.Create)

		response := httptest.NewRecorder()
		request, err := http.NewRequest("POST", "/api/ponds", bytes.NewBuffer(requestBodyJson))
		if err != nil {
			t.Fatal(err)
		}

		engine.ServeHTTP(response, request)

		// parsing response body
		var responseBody map[string]any
		err = json.Unmarshal(response.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatal(err)
		}

		// test response
		assert.Equal(t, errObject.Code, response.Code, "status code should be equal")
		assert.Equal(t, "error", responseBody["status"], "status should be equal")
		assert.Equal(t, errObject.Message, responseBody["message"], "message should be equal")
		assert.Equal(t, errObject.Err.Error(), responseBody["error"], "error should be equal")

		mockCall.Unset()
	})
}
