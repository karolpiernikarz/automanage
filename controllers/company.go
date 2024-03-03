package controllers

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/karolpiernikarz/automanage/cache/dbcache"
	"github.com/karolpiernikarz/automanage/models"
	"github.com/spf13/viper"
)

func CompanyPage(c *gin.Context) {
	// get the request domain
	domain := c.Request.Host
	redirectUrl := viper.GetString("app.companyUrl")
	// get the company id from the domain
	companies, err := dbcache.Companies()
	if err != nil {
		c.Redirect(302, redirectUrl)
		return
	}
	restaurants, err := dbcache.Restaurants()
	if err != nil {
		c.Redirect(302, redirectUrl)
		return
	}

	var response models.CompanyHtmlResponse
	for _, company := range companies {
		companyDomain, err := url.Parse(company.Domain)
		if err != nil {
			continue
		}
		if companyDomain.Host == domain {
			for _, restaurant := range restaurants {
				if restaurant.CompanyId == company.Id {
					if response.Logo == "" {
						response.Logo = viper.GetString("app.adminUrl") + "/storage/" + restaurant.Logo
					}

					// check if restaurant is active
					if restaurant.IsActive == 0 {
						continue
					}

					var restaurantResponse models.CompanyHtmlResponseRestaurants
					restaurantResponse.Name = restaurant.Name
					restaurantResponse.Address = restaurant.Address
					restaurantResponse.Url = restaurant.Website
					response.Restaurants = append(response.Restaurants, restaurantResponse)
				}
			}
			response.CompanyName = company.ChainName
			response.CompanyDesc = company.Description
			response.BackgroundUrl = viper.GetString("app.adminUrl") + "/storage/" + company.BackgroundImage
			break
		}
	}
	if len(response.Restaurants) == 0 {
		// if no restaurants found, redirect
		c.Redirect(302, redirectUrl)
		return
	}
	gridLength := []int{4, 3, 2, 1}
	if len(response.Restaurants) == 3 {
		gridLength = []int{3, 3, 2, 1}
	} else if len(response.Restaurants) == 2 {
		gridLength = []int{2, 2, 2, 1}
	} else if len(response.Restaurants) == 1 {
		gridLength = []int{1, 1, 1, 1}
	}
	// render the html
	c.HTML(http.StatusOK, "company.tmpl", gin.H{
		"data":       response,
		"gridLength": gridLength,
	})
}
