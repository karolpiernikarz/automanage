package newAdmin

import (
	"strconv"
	"time"

	"github.com/karolpiernikarz/automanage/models"
)

func MigrateOrder(order models.RestaurantOrders, oldRestaurantId string, newRestaurantId int) (newOrder Order) {
	newOrder.OrderNumber = oldRestaurantId + "-" + order.OrderNumber
	newOrder.CustomerId = order.CustomerId
	newOrder.RestaurantId = newRestaurantId
	newOrder.Customer = make(map[string]interface{})
	newOrder.Customer["name"] = order.Customer.Data().Name
	newOrder.Customer["surname"] = order.Customer.Data().Surname
	newOrder.Customer["phone"] = order.Customer.Data().Phone
	newOrder.Customer["email"] = order.Customer.Data().Email
	newOrder.Address = make(map[string]interface{})
	newOrder.Address["name"] = order.Address.Data().Name
	newOrder.Address["surname"] = order.Address.Data().Name
	newOrder.Address["email"] = order.Address.Data().Email
	newOrder.Address["placeid"] = ""
	newOrder.Address["lat"] = order.Address.Data().Lat
	newOrder.Address["long"] = order.Address.Data().Long
	newOrder.Address["detail"] = order.Address.Data().Detail
	newOrder.Address["phone"] = order.Address.Data().Phone
	newOrder.Prices = make(map[string]interface{})

	subtotalFloat, _ := strconv.ParseFloat(order.Prices.Data().SubTotal.String(), 64)
	totalFloat, _ := strconv.ParseFloat(order.Prices.Data().Total.String(), 64)
	deliveryFloat, _ := strconv.ParseFloat(order.Prices.Data().Delivery.String(), 64)
	bagFloat, _ := strconv.ParseFloat(order.Prices.Data().Bag.String(), 64)
	paymentFeeFloat, _ := strconv.ParseFloat(order.Prices.Data().PaymentFee.String(), 64)
	serviceFeeFloat, _ := strconv.ParseFloat(order.Prices.Data().ServiceFee.String(), 64)
	bankFeeFloat, _ := strconv.ParseFloat(order.Prices.Data().BankFee.String(), 64)
	currierFloat, _ := strconv.ParseFloat(order.Prices.Data().Currier.String(), 64)

	newOrder.Prices["sub_total"] = subtotalFloat
	newOrder.Prices["total"] = totalFloat
	newOrder.Prices["delivery"] = deliveryFloat
	newOrder.Prices["bag"] = bagFloat
	newOrder.Prices["payment_fee"] = paymentFeeFloat
	newOrder.Prices["serviceFee"] = serviceFeeFloat
	newOrder.Prices["bankFee"] = bankFeeFloat
	newOrder.Prices["currier"] = currierFloat

	discountTotalFloat, _ := strconv.ParseFloat(order.Prices.Data().Discount.Data().Total.String(), 64)

	newOrder.Prices["discount"] = map[string]interface{}{
		"coupon_id": order.Prices.Data().Discount.Data().CouponId,
		"type":      order.Prices.Data().Discount.Data().Type,
		"code":      order.Prices.Data().Discount.Data().Code,
		"discount":  order.Prices.Data().Discount.Data().Discount.String(),
		"total":     discountTotalFloat,
	}
	newOrder.Payment = make(map[string]interface{})
	newOrder.Payment["id"] = order.Payment.Data().Id
	newOrder.Payment["amount"] = order.Payment.Data().Amount
	newOrder.Payment["currency"] = order.Payment.Data().Currency
	newOrder.Payment["card"] = map[string]interface{}{
		"name":      order.Payment.Data().Card.Data().Name,
		"exp_month": order.Payment.Data().Card.Data().ExpMonth,
		"exp_year":  order.Payment.Data().Card.Data().ExpYear,
		"last4":     order.Payment.Data().Card.Data().Last4,
		"type":      order.Payment.Data().Card.Data().Type,
	}
	newOrder.Payment["receipt_url"] = ""
	newOrder.Payment["status"] = ""
	newOrder.Payment["refunded"] = ""
	newOrder.PaymentId = order.PaymentId
	newOrder.PaymentType = order.PaymentType
	newOrder.Status = order.Status
	newOrder.Type = order.Type
	newOrder.Note = order.Note
	orderDate, _ := time.Parse("2006-01-02 15:04:05", order.Date)
	newOrder.Date = orderDate
	newOrder.Currier = JSONBArray{}
	newOrder.IsPreOrder = order.IsPreOrder
	createdAt, _ := time.Parse("2006-01-02 15:04:05", order.CreatedAt)
	newOrder.CreatedAt = createdAt
	updatedAt, _ := time.Parse("2006-01-02 15:04:05", order.UpdatedAt)
	newOrder.UpdatedAt = updatedAt
	return
}

func MigrateProduct(product models.RestaurantProducts, menuId int, categoryId int) (newProduct Product) {
	newProduct.MenuId = menuId
	newProduct.CategoryId = categoryId
	newProduct.Name = product.Name
	newProduct.Slug = product.Slug
	newProduct.Image = product.Image
	newProduct.WithoutDiscountPrice = product.WithoutDiscountPrice
	newProduct.Price = product.Price
	newProduct.Description = product.Materials
	newProduct.Sort = product.Sort
	newProduct.IsActive = product.IsActive
	newProduct.Keywords = product.Keywords
	newProduct.Ingredients = "[\"\"]"
	newProduct.Discount = JSONB{
		"type":   1,
		"amount": nil,
	}
	newProduct.Type = "[\"0\"]"
	newProduct.Unit = JSONB{
		"type":    1,
		"amount":  nil,
		"portion": nil,
		"calorie": nil,
	}
	newProduct.Allergen = "[\"\"]"
	parsedCreatedAt, _ := time.Parse("2006-01-02 15:04:05", product.CreatedAt)
	newProduct.CreatedAt = parsedCreatedAt
	parsedUpdatedAt, _ := time.Parse("2006-01-02 15:04:05", product.UpdatedAt)
	newProduct.UpdatedAt = parsedUpdatedAt
	return
}

func MigrateVariant(variant models.RestaurantProduct_Variants, productId int) (newVariant Variant) {
	newVariant.ProductId = productId
	newVariant.Name = variant.Name
	parsedCreatedAt, _ := time.Parse("2006-01-02 15:04:05", variant.CreatedAt)
	newVariant.CreatedAt = parsedCreatedAt
	parsedUpdatedAt, _ := time.Parse("2006-01-02 15:04:05", variant.UpdatedAt)
	newVariant.UpdatedAt = parsedUpdatedAt
	return
}

func MigrateVariantOption(variantOption models.RestaurantProduct_Variant_Options, variantId int) (newVariantOption VariantOption) {
	newVariantOption.VariantId = variantId
	newVariantOption.Name = variantOption.Name
	newVariantOption.Price = variantOption.Price
	parsedCreatedAt, _ := time.Parse("2006-01-02 15:04:05", variantOption.CreatedAt)
	newVariantOption.CreatedAt = parsedCreatedAt
	parsedUpdatedAt, _ := time.Parse("2006-01-02 15:04:05", variantOption.UpdatedAt)
	newVariantOption.UpdatedAt = parsedUpdatedAt
	return
}

func MigrateExtraGroup(extraGroup models.RestaurantExtra_Groups, menuId int) (newExtraGroup ExtraGroup) {
	newExtraGroup.MenuId = menuId
	newExtraGroup.Name = extraGroup.Name
	newExtraGroup.DisplayName = extraGroup.DisplayName
	limit := 1
	limit, _ = strconv.Atoi(extraGroup.Limit)
	newExtraGroup.Limit = limit
	newExtraGroup.Sort = extraGroup.Order
	newExtraGroup.IsActive = 1
	parsedCreatedAt, _ := time.Parse("2006-01-02 15:04:05", extraGroup.CreatedAt)
	newExtraGroup.CreatedAt = parsedCreatedAt
	parsedUpdatedAt, _ := time.Parse("2006-01-02 15:04:05", extraGroup.UpdatedAt)
	newExtraGroup.UpdatedAt = parsedUpdatedAt
	return
}

func MigrateExtra(extra models.RestaurantExtras, extraGroupId int) (newExtra Extra) {
	newExtra.GroupId = extraGroupId
	newExtra.Name = extra.Name
	newExtra.Price = extra.Price
	newExtra.Sort = 1
	newExtra.IsDisabled = extra.IsDisabled
	newExtra.IsDefault = extra.IsDefault
	newExtra.Limit = 1
	parsedCreatedAt, _ := time.Parse("2006-01-02 15:04:05", extra.CreatedAt)
	newExtra.CreatedAt = parsedCreatedAt
	parsedUpdatedAt, _ := time.Parse("2006-01-02 15:04:05", extra.UpdatedAt)
	newExtra.UpdatedAt = parsedUpdatedAt
	return
}

func MigrateRestaurantSetting(setting models.RestaurantSettings, restaurantId int) (newSetting RestaurantSetting) {
	newSetting.RestaurantId = restaurantId
	newSetting.Name = setting.Name
	newSetting.Value = setting.Value
	parsedCreatedAt, _ := time.Parse("2006-01-02 15:04:05", setting.CreatedAt)
	newSetting.CreatedAt = parsedCreatedAt
	parsedUpdatedAt, _ := time.Parse("2006-01-02 15:04:05", setting.UpdatedAt)
	newSetting.UpdatedAt = parsedUpdatedAt
	return
}

func MigrateRestaurant(restaurant models.Restaurant) (newRestaurant Restaurant) {
	newRestaurant.CompanyId = restaurant.CompanyId
	newRestaurant.Name = restaurant.Name
	slug := CreateSlug(restaurant.Name)
	newRestaurant.Slug = slug
	newRestaurant.Logo = restaurant.Logo
	newRestaurant.Address = JSONB{
		"detail":    restaurant.Address,
		"post_code": 1,
	}
	newRestaurant.Phone = restaurant.Phone
	newRestaurant.Commission = JSONB{
		"pickup":     "0",
		"delivery":   "0",
		"restaurant": "0",
	}
	newRestaurant.Bank = JSONB{
		"name":         "",
		"reg_number":   "",
		"konto_number": "",
		"swift":        "",
		"iban":         "",
	}

	newRestaurant.PlaceId = ""
	newRestaurant.Lat = "0"
	newRestaurant.Long = "0"
	newRestaurant.IsActive = restaurant.IsActive
	newRestaurant.PlatformIsActive = 1
	newRestaurant.WebIsActive = 1
	parsedCreatedAt, _ := time.Parse("2006-01-02 15:04:05", restaurant.CreatedAt)
	newRestaurant.CreatedAt = parsedCreatedAt
	parsedUpdatedAt, _ := time.Parse("2006-01-02 15:04:05", restaurant.UpdatedAt)
	newRestaurant.UpdatedAt = parsedUpdatedAt
	return
}

func MigrateCompany(company models.Companies) (newCompany Company) {
	newCompany.Name = company.Name
	newCompany.ChainName = company.ChainName
	newCompany.Notes = ""

	newCompany.TaxNumber = company.TaxNumber
	newCompany.Contact = JSONB{
		"name":  company.Name,
		"phone": company.Phone,
		"email": company.Email,
	}
	newCompany.Domain = company.Domain
	newCompany.Social = JSONB{
		"facebook":  company.Social.Data().Facebook,
		"instagram": company.Social.Data().Instagram,
		"x":         "",
	}

	newCompany.Settings = JSONB{
		"test": "test",
	}
	newCompany.IsActive = company.IsActive
	parsedCreatedAt, _ := time.Parse("2006-01-02 15:04:05", company.CreatedAt)
	newCompany.CreatedAt = parsedCreatedAt
	parsedUpdatedAt, _ := time.Parse("2006-01-02 15:04:05", company.UpdatedAt)
	newCompany.UpdatedAt = parsedUpdatedAt
	return
}

func MigrateCategory(category models.RestaurantCategories, menuId int) (newCategory Category) {
	newCategory.MenuId = menuId
	newCategory.Name = category.Name
	newCategory.Description = category.Description
	newCategory.Slug = category.Slug
	newCategory.Icon = category.Icon
	newCategory.Banner = category.Banner
	newCategory.Sort = category.Sort
	newCategory.IsActive = category.IsActive
	newCategory.Discount = 0
	newCategory.Hours = JSONB{
		"monday":    nil,
		"tuesday":   nil,
		"wednesday": nil,
		"thursday":  nil,
		"friday":    nil,
		"saturday":  nil,
		"sunday":    nil,
	}
	parsedCreatedAt, _ := time.Parse("2006-01-02 15:04:05", category.CreatedAt)
	newCategory.CreatedAt = parsedCreatedAt
	parsedUpdatedAt, _ := time.Parse("2006-01-02 15:04:05", category.UpdatedAt)
	newCategory.UpdatedAt = parsedUpdatedAt
	return
}

func MigrateOrderProduct(orderProduct models.RestaurantOrder_Products, orderId int, productId int) (newOrderProduct OrderProduct) {
	newOrderProduct.OrderId = orderId
	newOrderProduct.ProductId = productId
	newOrderProduct.Qty = orderProduct.Qty
	newOrderProduct.Price = orderProduct.Price
	newOrderProduct.UnitPrice = orderProduct.UnitPrice
	variants := JSONBArray{}
	for _, variant := range *orderProduct.Variants.Data() {
		variants = append(variants, JSONB{
			"variant": JSONB{
				"id":   variant.Variant.Data().Id,
				"name": variant.Variant.Data().Name,
			},
			"option": JSONB{
				"id":   variant.Option.Data().Id,
				"name": variant.Option.Data().Name,
			},
			"price": variant.Price,
			"formatted": JSONB{
				"non_price": variant.Formatted.Data().NonPrice,
				"add_price": variant.Formatted.Data().AddPrice,
			},
		})
	}
	newOrderProduct.Variants = variants
	extras := JSONBArray{}
	for _, extra := range *orderProduct.Extras.Data() {
		extras = append(extras, JSONB{
			"id":              extra.Id,
			"name":            extra.Name,
			"price":           extra.Price,
			"formatted_price": extra.FormattedPrice,
			"group": JSONB{
				"id":           extra.Group.Data().Id,
				"name":         extra.Group.Data().Name,
				"display_name": extra.Group.Data().DisplayName,
				"limit":        extra.Group.Data().Limit,
			},
		})
	}
	newOrderProduct.Extras = extras
	newOrderProduct.Note = orderProduct.Comment
	newOrderProduct.CreatedAt = time.Now()
	newOrderProduct.UpdatedAt = time.Now()
	return
}
