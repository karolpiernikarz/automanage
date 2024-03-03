package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		validateToken(c)
		c.Next()
	}
}

func validateToken(c *gin.Context) {
	token := c.PostForm("apitoken")
	if token == "" {
		token = c.Query("apitoken")
	}
	if token == "" {
		c.AbortWithStatus(401)
	} else if checkToken(token) {
		c.Next()
	} else {
		c.AbortWithStatus(401)
	}
}

func checkToken(token string) bool {
	return token == viper.GetString("app.apitoken")
}
