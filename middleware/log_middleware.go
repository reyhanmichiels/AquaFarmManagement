package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/reyhanmichiels/AquaFarmManagement/infrastructure"
	"github.com/sirupsen/logrus"
)

func LogEvent(c *gin.Context) {
	ip := c.ClientIP()
	endpoint := c.Request.URL.Path
	method := c.Request.Method

	infrastructure.Logger.WithFields(logrus.Fields{
		"METHOD":   method,
		"ENDPOINT": endpoint,
		"IP":       ip,
	}).Info("Incoming HTTP Request")

	c.Next()
}
