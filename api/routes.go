package api

import (
	"investify/api/controller"
	"investify/api/middleware"
	"investify/api/services"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRouter(server *Server) *gin.Engine {
	router := server.router

	// Log statement for debugging

	// router.POST("/test", Test)

	//create a owner service
	//create a investor service
	//create a restaurant service

	v1 := router.Group("/api/v1")
	{
		// Log statement for debugging
		log.Println("Setting up API version 1 routes...")

		//Services
		//STart the s3 service over here

		userService := services.NewUserService(server.store, server.s3)
		// ownerService := services.NewOwnerService(server.store)
		investorService := services.NewInvestorService(server.store)
		businessService := services.NewBusinessService(server.store, server.s3)
		quickCodesService := services.NewQuickCodesService(server.store)

		// Controllers
		quickCodesController := controller.NewQuickCodesController(quickCodesService)
		userController := controller.NewUserController(server.store, userService, server.s3)
		// ownerController := controller.NewOwnerController(server.store, ownerService)
		investorController := controller.NewInvestorController(server.store, investorService)
		businessController := controller.NewBusinessController(server.store, businessService, server.s3)
		// extraController := controller.NewExtraController(server.store, server.s3)
		router.GET("/", userController.Test)

		// Define routes for users
		users := v1.Group("/users")
		{
			// Log statement for debugging
			log.Println("Setting up user routes...")

			users.POST("/signup", userController.CreateUser)
			users.POST("/login", userController.LoginUser)
			users.GET("/test", userController.Test)
			users.POST("/logout", userController.LogOut)
			users.POST("/uploadImage", middleware.JWTAuthAny(), userController.UploadImage)
			users.GET("/getProfileImage", middleware.JWTAuthAny(), userController.GetImage)
			users.DELETE("/deleteProfileImage", middleware.JWTAuthAny(), userController.DeleteImage)
		}
		// owner := v1.Group("/owner")
		// {

		// }
		investor := v1.Group("/investor")
		{
			investor.Use(middleware.JWTAuthInvestor())
			investor.GET("/feed", investorController.GetBusinessFeedController)
			investor.GET("/:id", investorController.GetInvestorByIdController)
			investor.GET("/business/:id", businessController.GetBusinessByIdController)

		}
		business := v1.Group("/business")
		{
			business.Use(middleware.JWTOwnerAuth())
			business.POST("/createBusiness", businessController.CreateBusiness)
			business.GET("/:id", businessController.GetBusinessByIdController)
			business.GET("/owner", businessController.GetBusinessByOwnerController)
			business.GET("/feed", businessController.GetInvestorFeedController)
			business.POST("/uploadBusinessImage/:id", businessController.UploadImage)
			business.GET("/getBusinessImages/:id", businessController.GetImageByBusinessId)
			business.DELETE("/deleteImage/:id", businessController.DeleteImageById)
			business.POST("/filter/", businessController.FilterBusinesses) // Corrected endpoint
			//send business id with the  image to upload to that specific business
		}
		quick_codes := v1.Group("/quick_codes")
		{
			quick_codes.GET("/states", quickCodesController.GetAllStates)
			quick_codes.GET("/domains", quickCodesController.GetAllDomains)

		}

		// extra := v1.Group("/extra")
		// {
		// 	extra.Use(middleware.JWTAuthAny())
		// 	extra.POST("/uploadImage", extraController.UploadImage)

		// }

	}

	// Log statement for debugging
	log.Println("Router setup complete.")

	return router
}
