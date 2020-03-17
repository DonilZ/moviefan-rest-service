package ui

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func jsonResponse(statusCode int, message string) gin.H {
	return gin.H{
		"statusCode": statusCode,
		"message":    message}
}

func tryBindJSON(c *gin.Context, object interface{}) bool {
	if err := c.ShouldBindJSON(object); err != nil {
		c.JSON(http.StatusBadRequest,
			jsonResponse(http.StatusBadRequest, err.Error()))
		return false
	}

	return true
}

func isEmptyField(c *gin.Context, field *string, fieldDescription string) bool {
	if *field == "" {
		c.JSON(http.StatusBadRequest,
			jsonResponse(http.StatusBadRequest,
				fmt.Sprintf("No %s entered!", fieldDescription)))
		return true
	}

	return false
}
