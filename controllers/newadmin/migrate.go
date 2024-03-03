package controllers_newadmin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/karolpiernikarz/automanage/helpers"
	newAdmin "github.com/karolpiernikarz/automanage/services/newadmin"
)

func MigrateRestaurant(c *gin.Context) {
	oldRestarauntid := c.Query("oldid")
	newRestaurantIdString := c.Query("newid")
	menuIdString := c.Query("menuid")

	newRestaurantId, err := strconv.Atoi(newRestaurantIdString)
	if err != nil {
		c.JSON(400, gin.H{"error": "newid must be an integer"})
		return
	}
	menuId, err := strconv.Atoi(menuIdString)
	if err != nil {
		c.JSON(400, gin.H{"error": "menuId must be an integer"})
		return
	}

	categories := helpers.GetRestaurantCategories(oldRestarauntid)
	products := helpers.GetRestaurantProducts(oldRestarauntid)
	variants := helpers.GetRestaurantProductVariants(oldRestarauntid)
	variant_options := helpers.GetRestaurantProductVariantOptions(oldRestarauntid)
	extragroups := helpers.GetRestaurantExtraGroups(oldRestarauntid)
	extras := helpers.GetRestaurantExtras(oldRestarauntid)
	settings := helpers.GetRestaurantSettings(oldRestarauntid)
	productextragroups := helpers.GetRestaurantProductExtraGroups(oldRestarauntid)
	orderProducts := helpers.GetRestaurantOrderProducts(oldRestarauntid)
	orders := helpers.GetOrdersFromRestaurant(oldRestarauntid)

	extrasandproductsmatch := make(map[int]int)
	variantoptionidmatch := make(map[int]int)
	productsmatch := make(map[int]int)

	for _, setting := range settings {
		newsetting := newAdmin.MigrateRestaurantSetting(setting, newRestaurantId)
		err := newsetting.Create()
		if err != nil {
			panic(err)
		}
	}

	for _, extragroup := range extragroups {
		newextragroup := newAdmin.MigrateExtraGroup(extragroup, menuId)
		err := newextragroup.Create()
		if err != nil {
			panic(err)
		}
		extrasandproductsmatch[extragroup.Id] = newextragroup.Id
		for _, extra := range extras {
			if extra.GroupId == extragroup.Id {
				newextra := newAdmin.MigrateExtra(extra, newextragroup.Id)
				err := newextra.Create()
				if err != nil {
					panic(err)
				}
			}
		}
	}

	for _, category := range categories {
		newcategory := newAdmin.MigrateCategory(category, menuId)
		err := newcategory.Create()
		if err != nil {
			panic(err)
		}
		for _, product := range products {
			if product.CategoryId == category.Id {
				newproduct := newAdmin.MigrateProduct(product, menuId, newcategory.Id)
				err := newproduct.Create()
				if err != nil {
					panic(err)
				}
				productsmatch[product.Id] = newproduct.Id
				for _, variant := range variants {
					if variant.ProductId == product.Id {
						newvariant := newAdmin.MigrateVariant(variant, newproduct.Id)
						err := newvariant.Create()
						if err != nil {
							panic(err)
						}
						for _, variant_option := range variant_options {
							if variant_option.ProductVariantId == variant.Id {
								newvariant_option := newAdmin.MigrateVariantOption(variant_option, newvariant.Id)
								err := newvariant_option.Create()
								if err != nil {
									panic(err)
								}
								variantoptionidmatch[variant_option.Id] = newvariant_option.Id
							}
						}
					}
				}
			}
		}
	}

	for _, productextragroup := range productextragroups {
		productVariantOptionId := variantoptionidmatch[productextragroup.ProductVariantOptionId]
		productId := productsmatch[productextragroup.ProductId]
		extraGroupId := extrasandproductsmatch[productextragroup.ExtraGroupId]

		if productVariantOptionId == 0 || productId == 0 || extraGroupId == 0 {
			continue
		}

		newproductextragroup := newAdmin.ProductExtraGroups{}
		newproductextragroup.ProductId = productId
		newproductextragroup.ProductVariantOptionId = productVariantOptionId
		newproductextragroup.ExtraGroupId = extraGroupId
		newproductextragroup.Sort = 1
		err := newproductextragroup.Create()
		if err != nil {
			panic(err)
		}
	}

	for _, order := range orders {
		neworder := newAdmin.MigrateOrder(order, oldRestarauntid, newRestaurantId)
		err := neworder.Create()
		if err != nil {
			panic(err)
		}
		for _, orderProduct := range orderProducts {
			if orderProduct.OrderId == order.Id {
				productId := productsmatch[orderProduct.ProductId]
				neworderProduct := newAdmin.MigrateOrderProduct(orderProduct, neworder.Id, productId)
				err := neworderProduct.Create()
				if err != nil {
					panic(err)
				}
			}
		}
	}
	c.JSON(200, gin.H{"message": "success"})

}
