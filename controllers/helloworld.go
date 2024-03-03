package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Helloworld(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}
