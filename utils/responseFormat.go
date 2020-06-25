package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// responseFormat xxx
func responseFormat(c *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

// Success xxx
func Success(c *gin.Context, data gin.H, msg string) {
	responseFormat(c, http.StatusOK, 200, data, msg)
}

// Fail xxx
func Fail(c *gin.Context, data gin.H, msg string) {
	responseFormat(c, http.StatusOK, 400, data, msg)
}

// JwtFail xxx
func JwtFail(c *gin.Context, data gin.H, msg string) {
	responseFormat(c, http.StatusUnauthorized, 400, data, msg)
}
