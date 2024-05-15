package main

import (
	"ERP-API/handlers"
	"ERP-API/middleware"
	"ERP-API/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDB()
	defer models.Database.Close()

	router := gin.Default()

	//Enterprises (public)
	router.POST("/login", handlers.Login)
	router.POST("/register", handlers.Register)
	//Subscriptions
	router.GET("/subscriptions", handlers.GetSubscriptions)
	router.GET("/subscriptions/:id", handlers.GetSubscriptionByName)

	protectedRoutes := router.Group("/")
	protectedRoutes.Use(middleware.AuthenticationMiddleware())
	{
		//Enterprises
		protectedRoutes.GET("/profile", handlers.Profile)
		//Contracts
		protectedRoutes.GET("/contracts", handlers.GetContracts)
		protectedRoutes.GET("/contracts/import", handlers.GetImportContracts)
		protectedRoutes.GET("/contracts/export", handlers.GetExportContracts)
		//requests?
		//...
		protectedRoutes.GET("/contracts/:id", handlers.GetContractById)
		protectedRoutes.POST("/contracts", handlers.PostContract)
		//...
		//Products
		protectedRoutes.GET("/products", handlers.GetProducts)
		protectedRoutes.GET("/products/:id", handlers.GetProductByID)
		protectedRoutes.POST("/products", handlers.PostProduct)
		//ProductionOrders
		protectedRoutes.GET("/productionOrders", handlers.GetProductionOrders)
		protectedRoutes.GET("/productionOrders/:id", handlers.GetProductionOrderByID)
		protectedRoutes.POST("/productionOrders", handlers.PostProductionOrder)
		//Deliveries
		protectedRoutes.GET("/deliveries", handlers.GetDeliveries)
		protectedRoutes.GET("/deliveries/:id", handlers.GetDeliveryByID)
		//Debtors
		protectedRoutes.GET("/debtors", handlers.GetDebtors)
		protectedRoutes.GET("/debtors/:id", handlers.GetDebtorByID)
		protectedRoutes.POST("/debtors", handlers.PostDebtor)
		//Partners
		//protectedRoutes.GET("/partners", handlers.)
	}

	router.Run("localhost:8080")
}
