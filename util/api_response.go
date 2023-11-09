package util

import "github.com/gin-gonic/gin"

func SuccessResponse(c *gin.Context, code int, message string, data interface{}) {

	c.JSON(code, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})

}

func FailResponse(c *gin.Context, code int, message string, err error) {

	c.JSON(code, gin.H{
		"status":  "error",
		"message": message,
		"error":   err.Error(),
	})

}

type ErrorObject struct {
	Code    int
	Message string
	Err     error
}
