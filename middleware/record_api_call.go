package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/reyhanmichiels/AquaFarmManagement/domain"
	"github.com/reyhanmichiels/AquaFarmManagement/infrastructure/database"
)

func RecordApiCallMiddleware(c *gin.Context) {
	apiCall := domain.ApiCall{
		IpAdress: c.ClientIP(),
		Endpoint: c.Request.URL.Path,
		Method:   c.Request.Method,
	}

	database.DB.Create(&apiCall)
	c.Next()
}
