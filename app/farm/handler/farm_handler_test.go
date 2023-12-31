package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
		//prepare request body
		requestBody := domain.FarmBind{
			Name: "testName",
		}

		requestBodyJson, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatal(err)
		}

		// call mock
		mockCallResponse := domain.Farm{
			ID:   "testID",
			Name: requestBody.Name,
		}

		mockCall := farmUsecaseMock.Mock.On("Create", requestBody).Return(mockCallResponse, nil)

		// call handler
		engine := gin.Default()
		engine.POST("/api/farms", farmHandler.Create)

		response := httptest.NewRecorder()
		request, err := http.NewRequest("POST", "/api/farms", bytes.NewBuffer(requestBodyJson))
		if err != nil {
			t.Fatal(err.Error())
		}

		engine.ServeHTTP(response, request)

		//test response
		var responseBody map[string]any
		err = json.Unmarshal(response.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatal(err.Error())
		}

		farmData := responseBody["data"].(map[string]any)
		assert.Equal(t, http.StatusCreated, response.Code, "status code should be equal")
		assert.Equal(t, "success", responseBody["status"], "status should be equal")
		assert.Equal(t, "successfully create farm", responseBody["message"], "message should be equal")
		assert.Equal(t, mockCallResponse.ID, farmData["id"], "farm id should be equal")
		assert.Equal(t, mockCallResponse.Name, farmData["name"], "farm name should be equal")

		mockCall.Unset()
	})

	t.Run("should reject when request invalid", func(t *testing.T) {
		// call handler
		engine := gin.Default()
		engine.POST("/api/farms", farmHandler.Create)

		response := httptest.NewRecorder()
		request, err := http.NewRequest("POST", "/api/farms", nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		engine.ServeHTTP(response, request)

		//test response
		var responseBody map[string]any
		err = json.Unmarshal(response.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatal(err.Error())
		}

		assert.Equal(t, http.StatusBadRequest, response.Code, "status code should be equal")
		assert.Equal(t, "error", responseBody["status"], "status should be equal")
		assert.Equal(t, "failed to bind input", responseBody["message"], " message should be equal")
	})

	t.Run("should reject when usecase call return error", func(t *testing.T) {
		//prepare request body
		requestBody := domain.FarmBind{
			Name: "testName",
		}

		requestBodyJson, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatal(err)
		}

		// call mock
		errObject := util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "testMessage",
			Err:     errors.New("testError"),
		}

		mockCall := farmUsecaseMock.Mock.On("Create", requestBody).Return(nil, errObject)

		// call handler
		engine := gin.Default()
		engine.POST("/api/farms", farmHandler.Create)

		response := httptest.NewRecorder()
		request, err := http.NewRequest("POST", "/api/farms", bytes.NewBuffer(requestBodyJson))
		if err != nil {
			t.Fatal(err.Error())
		}

		engine.ServeHTTP(response, request)

		//test response
		var responseBody map[string]any
		err = json.Unmarshal(response.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatal(err.Error())
		}

		assert.Equal(t, errObject.Code, response.Code, "status code should be equal")
		assert.Equal(t, "error", responseBody["status"], "status should be equal")
		assert.Equal(t, errObject.Message, responseBody["message"], " message should be equal")
		assert.Equal(t, errObject.Err.Error(), responseBody["error"], " error should be equal")

		mockCall.Unset()
	})
}

func TestUpdateFarm(t *testing.T) {
	t.Run("should can be update farm", func(t *testing.T) {
		//prepare request body
		requestBody := domain.FarmBind{
			Name: "testName",
		}

		requestBodyJson, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatal(err)
		}

		// call mock
		mockCallResponse := domain.Farm{
			ID:   "testID",
			Name: requestBody.Name,
		}

		mockCall := farmUsecaseMock.Mock.On("Update", requestBody).Return(mockCallResponse, nil)

		// call handler
		engine := gin.Default()
		engine.PUT("/api/farms/:farmId", farmHandler.Update)

		response := httptest.NewRecorder()
		request, err := http.NewRequest("PUT", "/api/farms/testID", bytes.NewBuffer(requestBodyJson))
		if err != nil {
			t.Fatal(err.Error())
		}

		engine.ServeHTTP(response, request)

		//test response
		var responseBody map[string]any
		err = json.Unmarshal(response.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatal(err.Error())
		}

		farmData := responseBody["data"].(map[string]any)
		assert.Equal(t, http.StatusOK, response.Code, "status code should be equal")
		assert.Equal(t, "success", responseBody["status"], "status should be equal")
		assert.Equal(t, "successfully update farm", responseBody["message"], "message should be equal")
		assert.Equal(t, mockCallResponse.ID, farmData["id"], "farm id should be equal")
		assert.Equal(t, mockCallResponse.Name, farmData["name"], "farm name should be equal")

		mockCall.Unset()
	})

	t.Run("should reject when request invalid", func(t *testing.T) {
		// call handler
		engine := gin.Default()
		engine.PUT("/api/farms/:farmId", farmHandler.Update)

		response := httptest.NewRecorder()
		requestCall, err := http.NewRequest("PUT", "/api/farms/testID", nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		engine.ServeHTTP(response, requestCall)

		//test response
		var responseBody map[string]any
		err = json.Unmarshal(response.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatal(err.Error())
		}

		assert.Equal(t, http.StatusBadRequest, response.Code, "status code should be equal")
		assert.Equal(t, "error", responseBody["status"], "status should be equal")
		assert.Equal(t, "failed to bind input", responseBody["message"], " message should be equal")
	})

	t.Run("should reject when usecase call return error", func(t *testing.T) {
		//prepare request body
		responseBody := domain.FarmBind{
			Name: "testName",
		}

		responseBodyJson, err := json.Marshal(responseBody)
		if err != nil {
			t.Fatal(err)
		}

		// call mock
		errObject := util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "testMessage",
			Err:     errors.New("testError"),
		}

		mockCall := farmUsecaseMock.Mock.On("Update", responseBody).Return(nil, errObject)

		// call handler
		engine := gin.Default()
		engine.PUT("/api/farms/:farmId", farmHandler.Update)

		response := httptest.NewRecorder()
		requestCall, err := http.NewRequest("PUT", "/api/farms/testID", bytes.NewBuffer(responseBodyJson))
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

		assert.Equal(t, errObject.Code, response.Code, "status code should be equal")
		assert.Equal(t, "error", responseData["status"], "status should be equal")
		assert.Equal(t, errObject.Message, responseData["message"], " message should be equal")
		assert.Equal(t, errObject.Err.Error(), responseData["error"], " error should be equal")

		mockCall.Unset()
	})
}

func TestGetFarm(t *testing.T) {
	t.Run("should can get all farm", func(t *testing.T) {
		// call mock
		mockCallResponse := []domain.Farm{
			{
				ID:   "testID1",
				Name: "testName1",
			},
			{
				ID:   "testID2",
				Name: "testName2",
			},
			{
				ID:   "testID3",
				Name: "testName3",
			},
		}

		mockCall := farmUsecaseMock.Mock.On("Get").Return(mockCallResponse, nil)

		// call handler
		engine := gin.Default()
		engine.GET("/api/farms", farmHandler.Get)

		response := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "/api/farms", nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		engine.ServeHTTP(response, request)

		//test response
		var responseBody map[string]any
		err = json.Unmarshal(response.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatal(err.Error())
		}

		assert.Equal(t, http.StatusOK, response.Code, "status code should be equal")
		assert.Equal(t, "success", responseBody["status"], "status should be equal")
		assert.Equal(t, "successfully get all farm", responseBody["message"], "message should be equal")

		farmsData := responseBody["data"].([]interface{})
		for i, v := range farmsData {
			v := v.(map[string]any)
			assert.Equal(t, mockCallResponse[i].Name, v["name"], "name should be equal")
			assert.Equal(t, mockCallResponse[i].ID, v["id"], "id should be equal")
		}

		mockCall.Unset()
	})

	t.Run("should reject when usecase call return error", func(t *testing.T) {
		// call mock
		errObject := util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "testMessage",
			Err:     errors.New("testError"),
		}

		mockCall := farmUsecaseMock.Mock.On("Get").Return(nil, errObject)

		// call handler
		engine := gin.Default()
		engine.GET("/api/farms", farmHandler.Get)

		response := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "/api/farms", nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		engine.ServeHTTP(response, request)

		//test response
		var responseBody map[string]any
		err = json.Unmarshal(response.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatal(err.Error())
		}

		assert.Equal(t, errObject.Code, response.Code, "status code should be equal")
		assert.Equal(t, "error", responseBody["status"], "status should be equal")
		assert.Equal(t, errObject.Err.Error(), responseBody["error"], "error should be equal")
		assert.Equal(t, errObject.Message, responseBody["message"], "message should be equal")

		mockCall.Unset()
	})
}

func TestGetFarmById(t *testing.T) {
	t.Run("should get farm by id", func(t *testing.T) {
		// call mock
		mockCallResponse := domain.FarmApi{
			ID:   "testID",
			Name: "farm1",
			Ponds: []domain.Pond{
				{
					ID:     "pondID1",
					Name:   "pond1",
					FarmID: "testID",
				},
				{
					ID:     "pondID2",
					Name:   "pond2",
					FarmID: "testID",
				},
			},
		}

		mockCall := farmUsecaseMock.Mock.On("GetFarmById", mockCallResponse.ID).Return(mockCallResponse, nil)

		// call handler
		engine := gin.Default()
		engine.GET("/api/farms/:farmId", farmHandler.GetFarmById)

		response := httptest.NewRecorder()
		request, err := http.NewRequest("GET", fmt.Sprintf("/api/farms/%s", mockCallResponse.ID), nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		engine.ServeHTTP(response, request)

		//test response
		var responseBody map[string]any
		err = json.Unmarshal(response.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatal(err.Error())
		}

		farmData := responseBody["data"].(map[string]any)
		assert.Equal(t, http.StatusOK, response.Code, "status code should be equal")
		assert.Equal(t, "success", responseBody["status"], "status should be equal")
		assert.Equal(t, "successfully get farm by id", responseBody["message"], "message should be equal")
		assert.Equal(t, mockCallResponse.ID, farmData["id"], "farm id should be equal")
		assert.Equal(t, mockCallResponse.Name, farmData["name"], "farm name should be equal")

		pondsData := farmData["ponds"].([]interface{})

		for i, v := range pondsData {
			v := v.(map[string]any)
			assert.Equal(t, mockCallResponse.Ponds[i].ID, v["id"], "pond id should be equal")
			assert.Equal(t, mockCallResponse.Ponds[i].FarmID, v["farm_id"], "farm id should be equal")
			assert.Equal(t, mockCallResponse.Ponds[i].Name, v["name"], "pond name should be equal")
		}
		mockCall.Unset()
	})

	t.Run("should reject when usecase call return error", func(t *testing.T) {
		// call mock
		errObject := util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "testMessage",
			Err:     errors.New("testError"),
		}
		mockCall := farmUsecaseMock.Mock.On("GetFarmById", "testID").Return(nil, errObject)

		// call handler
		engine := gin.Default()
		engine.GET("/api/farms/:farmId", farmHandler.GetFarmById)

		response := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "/api/farms/testID", nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		engine.ServeHTTP(response, request)

		//test response
		var responseBody map[string]any
		err = json.Unmarshal(response.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatal(err.Error())
		}

		assert.Equal(t, errObject.Code, response.Code, "status code should be equal")
		assert.Equal(t, "error", responseBody["status"], "status should be equal")
		assert.Equal(t, errObject.Err.Error(), responseBody["error"], "error should be equal")
		assert.Equal(t, errObject.Message, responseBody["message"], "message should be equal")

		mockCall.Unset()
	})
}

func TestDeleteFarm(t *testing.T) {
	t.Run("should can delete farm", func(t *testing.T) {
		// call mock
		mockCall := farmUsecaseMock.Mock.On("Delete", "testID").Return(nil)

		// call handler
		engine := gin.Default()
		engine.DELETE("/api/farms/:farmId", farmHandler.Delete)

		response := httptest.NewRecorder()
		request, err := http.NewRequest("DELETE", "/api/farms/testID", nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		engine.ServeHTTP(response, request)

		//test response
		var responseBody map[string]any
		err = json.Unmarshal(response.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatal(err.Error())
		}

		assert.Equal(t, http.StatusOK, response.Code, "status code should be equal")
		assert.Equal(t, "success", responseBody["status"], "status should be equal")
		assert.Equal(t, "successfully delete farm", responseBody["message"], "message should be equal")
		assert.Nil(t, responseBody["data"], "data should be nil")

		mockCall.Unset()
	})

	t.Run("should reject when usecase call return error", func(t *testing.T) {
		// call mock
		errObject := util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "testMessage",
			Err:     errors.New("testError"),
		}
		mockCall := farmUsecaseMock.Mock.On("Delete", "testID").Return(errObject)

		// call handler
		engine := gin.Default()
		engine.DELETE("/api/farms/:farmId", farmHandler.Delete)

		response := httptest.NewRecorder()
		request, err := http.NewRequest("DELETE", "/api/farms/testID", nil)
		if err != nil {
			t.Fatal(err.Error())
		}

		engine.ServeHTTP(response, request)

		//test response
		var responseBody map[string]any
		err = json.Unmarshal(response.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatal(err.Error())
		}

		assert.Equal(t, errObject.Code, response.Code, "status code should be equal")
		assert.Equal(t, "error", responseBody["status"], "status should be equal")
		assert.Equal(t, errObject.Err.Error(), responseBody["error"], "error should be equal")
		assert.Equal(t, errObject.Message, responseBody["message"], "message should be equal")

		mockCall.Unset()
	})
}
