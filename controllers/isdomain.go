package controllers

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	isd "github.com/jbenet/go-is-domain"
	"github.com/karolpiernikarz/automanage/helpers"
)

func IsDomainExist(c *gin.Context) {
	restaurantDomain := c.PostForm("domain")
	//check if domain variables is correct
	parsedUrl, err := url.Parse(restaurantDomain)
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusInternalServerError, Error: []string{err.Error()}})
		return
	}
	domain := ""
	if parsedUrl.Scheme != "" {
		domain = parsedUrl.Hostname()
	} else {
		domain = parsedUrl.String()
	}
	if !isd.IsDomain(domain) {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusBadRequest, Error: []string{"not a domain"}})
		return
	}
	if helpers.IsDomainExist(domain) {
		helpers.SendResponse(c, helpers.Response{
			Status:  http.StatusOK,
			Message: []string{"Ok"},
			Error:   []string{},
		})
		return
	}
	helpers.SendResponse(c, helpers.Response{
		Status:  http.StatusNotFound,
		Message: []string{"does not exist"},
		Error:   []string{},
	})
}
