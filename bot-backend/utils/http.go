package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MakeOkResponse(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func MakeBadRequestResponse(c *gin.Context, code, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":  code,
		"error": message,
	})
}
