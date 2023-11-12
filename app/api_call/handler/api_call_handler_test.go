package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	api_call_mock "github.com/reyhanmichiels/AquaFarmManagement/app/api_call/mock"
	"github.com/reyhanmichiels/AquaFarmManagement/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var apiCallUsecaseMock = api_call_mock.ApiCallUsecaseMock{
	Mock: mock.Mock{},
}

var apiCallHandler = NewApiCallHandler(&apiCallUsecaseMock)

func TestGet(t *testing.T) {
	t.Run("should can get api calls", func(t *testing.T) {
		// call mock
		mockResponse := make(map[string]map[string]int, 0)
		mockResponse["endpoint1"] = map[string]int{
			"count":             1,
			"unique_user_agent": 1,
		}
		mockResponse["endpoint2"] = map[string]int{
			"count":             2,
			"unique_user_agent": 1,
		}
		mockCall := apiCallUsecaseMock.Mock.On("Get").Return(mockResponse, nil)

		// call handler
		engine := gin.Default()
		engine.GET("/api/api-calls", apiCallHandler.Get)

		response := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "/api/api-calls", nil)
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
		assert.Equal(t, http.StatusOK, response.Code, "status code should be equal")
		assert.Equal(t, "success", responseBody["status"], "status should be equal")
		assert.Equal(t, "successfully get all api call", responseBody["message"], "message should be equal")

		apiCallData := responseBody["data"].(map[string]interface{})
		assert.Equal(t, len(mockResponse), len(apiCallData), "length should be equal")

		endpoint1Data := apiCallData["endpoint1"].(map[string]any)
		assert.Equal(t, float64(mockResponse["endpoint1"]["count"]), endpoint1Data["count"], "endpoint1 count should be equal")
		assert.Equal(t, float64(mockResponse["endpoint1"]["unique_user_agent"]), endpoint1Data["unique_user_agent"], "endpoint1 unique user agent should be equal")

		endpoint2Data := apiCallData["endpoint2"].(map[string]any)
		assert.Equal(t, float64(mockResponse["endpoint2"]["count"]), endpoint2Data["count"], "endpoint2 count should be equal")
		assert.Equal(t, float64(mockResponse["endpoint2"]["unique_user_agent"]), endpoint2Data["unique_user_agent"], "endpoint2 unique user agent should be equal")

		mockCall.Unset()
	})

	t.Run("should reject when usecase call return error", func(t *testing.T) {
		// call mock
		errObject := util.ErrorObject{
			Code:    http.StatusInternalServerError,
			Err:     errors.New("testError"),
			Message: "test message",
		}
		mockCall := apiCallUsecaseMock.Mock.On("Get").Return(nil, errObject)

		// call handler
		engine := gin.Default()
		engine.GET("/api/api-calls", apiCallHandler.Get)

		response := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "/api/api-calls", nil)
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
