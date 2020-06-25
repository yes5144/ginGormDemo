package middlewares

import (
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yes5144/ginGormDemo/models"
	"github.com/yes5144/ginGormDemo/utils"
)

// JwtMiddleware xxx
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get tokenString from athorization header
		tokenString := c.GetHeader("Authorization")
		log.Println(tokenString)
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			utils.JwtFail(c, nil, "Authory not enough ")
			c.Abort()
			return
		}
		tokenString = tokenString[7:]

		// parseToken
		token, claims, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			utils.JwtFail(c, nil, "authory nt enough")
			c.Abort()
			return
		}
		// userId xxx
		userId := claims.UserID
		var user models.User
		user.SelectIds(strconv.FormatUint(userId, 10))

		// user exist
		if userId == 0 {
			utils.JwtFail(c, nil, "no authorization")
			c.Abort()
			return
		}

		// write into context
		c.Set("user", user)
		c.Next()
	}
}
