package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/yes5144/ginGormDemo/apis/v1"
	"github.com/yes5144/ginGormDemo/middlewares"
)

// InitRouter xxx
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode("debug")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "ping",
			"code": http.StatusOK,
		})
	})

	//
	r.POST("/api/auth/register", v1.Register)
	r.POST("/api/auth/login", v1.Login)
	r.GET("/api/auth/userinfo", middlewares.JwtMiddleware(), v1.UserInfo)

	return r
}
