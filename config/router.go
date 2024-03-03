package config

import (
	"html/template"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/karolpiernikarz/automanage/controllers"
	controllers_newadmin "github.com/karolpiernikarz/automanage/controllers/newadmin"
	"github.com/karolpiernikarz/automanage/controllers/restaurantapi"
	controllers_test "github.com/karolpiernikarz/automanage/controllers/test"
	"github.com/karolpiernikarz/automanage/middlewares"
	"github.com/karolpiernikarz/automanage/templates"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var router *gin.Engine

func InitRoutes() {
	// set gin mode
	if viper.GetString("app.env") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// init router
	router = gin.New()

	// init global middlewares
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	// use pprof if gin in debug mode
	if viper.GetString("app.env") == "development" {
		pprof.Register(router)
	}

	// set html template
	templateFs := templates.GetTemplateFs()
	templ := template.Must(template.New("").ParseFS(templateFs, "*.tmpl"))
	router.SetHTMLTemplate(templ)

	//router.LoadHTMLGlob("./templates/*")
	//router.Static("/view", "./view")

	newMachhubAdmin := router.Group("/api/nma")
	newMachhubAdmin.Use(middlewares.TokenAuthMiddleware())
	newMachhubAdmin.Use(middlewares.ZeroLogMiddleware)
	{
		newMachhubAdmin.GET("migrate/restaurant", controllers_newadmin.MigrateRestaurant)
	}

	machhubadmin := router.Group("/api/ma")
	machhubadmin.Use(middlewares.TokenAuthMiddleware())
	machhubadmin.Use(middlewares.ZeroLogMiddleware)
	{
		machhubadmin.GET("hello-world", controllers.Helloworld)
		machhubadmin.GET("orderboxhealth", controllers.OrderBoxHealthcheck)
		machhubadmin.GET("restaurant/address", controllers.GetRestaurantAddress)
		machhubadmin.GET("restaurant/addresses", controllers.GetRestaurantAddresses)
		machhubadmin.GET("restaurant/attribute", controllers.GetRestaurantAttribute)
		machhubadmin.GET("restaurant/attributes", controllers.GetRestaurantAttributes)
		machhubadmin.GET("restaurant/cart", controllers.GetRestaurantCart)
		machhubadmin.GET("restaurant/carts", controllers.GetRestaurantCarts)
		machhubadmin.GET("restaurant/category", controllers.GetRestaurantCategory)
		machhubadmin.GET("restaurant/categories", controllers.GetRestaurantCategories)
		machhubadmin.GET("restaurant/coupon", controllers.GetRestaurantCoupon)
		machhubadmin.GET("restaurant/coupons", controllers.GetRestaurantCoupons)
		machhubadmin.GET("restaurant/customer", controllers.GetRestaurantCustomer)
		machhubadmin.GET("restaurant/customers", controllers.GetRestaurantCustomers)
		machhubadmin.GET("restaurant/failed_jobs", controllers.GetRestaurantFailedJobs)
		machhubadmin.GET("restaurant/migrations", controllers.GetRestaurantMigrations)
		machhubadmin.GET("restaurant/order_actions", controllers.GetRestaurantOrderActions)
		machhubadmin.GET("restaurant/order_products", controllers.GetRestaurantOrderProducts)
		machhubadmin.GET("restaurant/order", controllers.GetRestaurantOrder)
		machhubadmin.GET("restaurant/orders", controllers.GetRestaurantOrders)
		machhubadmin.GET("restaurants/orders", controllers.GetAllRestaurantOrdersWithTime)
		machhubadmin.GET("restaurants/supportresponse", controllers.GetSupportResponse)
		machhubadmin.GET("restaurant/password_resets", controllers.GetRestaurantPasswordResets)
		machhubadmin.GET("restaurant/product", controllers.GetRestaurantProduct)
		machhubadmin.GET("restaurant/productextragroups", controllers.GetRestaurantProductExtraGroups)
		machhubadmin.GET("restaurant/products", controllers.GetRestaurantProducts)
		machhubadmin.GET("restaurant/productvariantoptions", controllers.GetRestaurantProductVariantOptions)
		machhubadmin.GET("restaurant/productvariants", controllers.GetRestaurantProductVariants)
		machhubadmin.GET("restaurant/settings", controllers.GetRestaurantSettings)
		machhubadmin.GET("restaurant/tlogs/:restaurantId", controllers.GetRestaurantTerminalLogs)
		machhubadmin.GET("restaurant/slogs/:restaurantId", controllers.GetRestaurantStorageLogs)
		machhubadmin.GET("restaurant/restart/:restaurantId", controllers.RestartRestaurantApp)
		machhubadmin.GET("restaurant/variables/:restaurantId", controllers.GetRestaurantVariables)
		machhubadmin.GET("restaurant/isopen/:restaurantId", controllers.GetRestaurantIsOpen)
		machhubadmin.GET("restaurants/reports/create", controllers.CreateRestaurantOrderReport)
		machhubadmin.GET("restaurants/reports/createwithtime", controllers.CreateRestaurantOrderReportWithTime)
		machhubadmin.GET("restaurants/reports/restaurantreport", controllers.CreateRestaurantReports)
		machhubadmin.GET("restaurants/count", controllers.GetRestaurantCount)
		machhubadmin.GET("restaurant/lastorders/:restaurantId", controllers.GetRestaurantLastNOrders)
		machhubadmin.GET("restaurant/settingsbyname/:restaurantId", controllers.GetRestaurantSettingsByName)
		machhubadmin.POST("restaurant", controllers.RestaurantCreate)
		machhubadmin.POST("tools/isdomainexist", controllers.IsDomainExist)
		machhubadmin.GET("company/orders/:companyId", controllers.GetRestaurantOrdersFromCompanyId)
		machhubadmin.POST("company/tables/:companyId", controllers.GetRestaurantTablesFromCompanyId)
		machhubadmin.POST("restaurant/tables/:restaurantId", controllers.GetRestaurantTablesFromRestaurantId)
		machhubadmin.POST("restaurants/tables/:restaurantIds", controllers.GetRestaurantTablesFromAllRestaurants)
		machhubadmin.GET("restaurants/open", controllers.GetRestaurantsOpen)
		machhubadmin.GET("courierlist", controllers.CourierList)
		machhubadmin.GET("restaurants", controllers.GetAllRestaurants)
		machhubadmin.GET("companies", controllers.GetAllCompanies)
		machhubadmin.GET("orders/search/byemail", controllers.GetOrdersFromEmail)
	}

	restaurant := router.Group("/api/restaurant")
	restaurant.Use(middlewares.RestaurantTokenAuthMiddleware())
	restaurant.Use(middlewares.ZeroLogMiddleware)
	{
		restaurant.POST("order", restaurantapi.Order)
		restaurant.POST("error", restaurantapi.Error)
	}

	orderbox := router.Group("/api/order/box")
	orderbox.Use(middlewares.ZeroLogMiddlewareForOrderbox)
	orderbox.Use(middlewares.SetRestaurantIdFromTerminalId())
	{
		orderbox.GET("list", controllers.OrderBoxList, middlewares.SetOrderBoxAlive())
		orderbox.GET("callback", controllers.OrderBoxCallBack, middlewares.CallbackHandler())
	}

	html := router.Group("/html")
	html.Use(middlewares.ZeroLogMiddleware)
	{
		html.GET("company", controllers.CompanyPage)
	}

	test := router.Group("/api/test")
	test.Use(middlewares.TokenAuthMiddlewareTest())
	test.Use(middlewares.ZeroLogMiddleware)
	{
		test.GET("restaurants/supportresponse", controllers_test.GetAllRestaurantOrdersWithTime)
		test.GET("restaurants/orderboxhealthid", controllers_test.OrderBoxHealthById)
		test.POST("restaurant/tables/:restaurantId", controllers_test.RestaurantTables)
	}

	aws := router.Group("/api/aws")
	aws.Use(middlewares.ZeroLogMiddleware)
	{
		aws.POST("ses/callback"+viper.GetString("aws.ses_callback"), middlewares.ZeroLogMiddlewareForStoreRawJson, controllers.AwsSesWebhook)
	}

	wolt := router.Group("/api/wolt")
	wolt.Use(middlewares.ZeroLogMiddleware)
	{
		wolt.POST("webhook", middlewares.ZeroLogMiddlewareForStoreRawJson, controllers.WoltWebHook)
		wolt.POST("webhooktest", middlewares.ZeroLogMiddlewareForStoreRawJson, controllers.WoltWebHookTest)
	}

	// run router
	err := router.Run(":" + viper.GetString("app.port"))
	if err != nil {
		log.Fatal().Err(err).Msg("Error while running router")
	}
}
