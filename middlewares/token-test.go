package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func TokenAuthMiddlewareTest() gin.HandlerFunc {
	return func(c *gin.Context) {
		validateTokenTest(c)
		c.Next()
	}
}

func validateTokenTest(c *gin.Context) {
	token := c.PostForm("apitoken")
	if token == "" {
		token = c.Query("apitoken")
	}
	if token == "" {
		c.AbortWithStatus(401)
	} else if checkTokenTest(token) {
		c.Next()
	} else {
		c.AbortWithStatus(401)
	}
}

func checkTokenTest(token string) bool {
	return token == viper.GetString("test.apitoken")
}
