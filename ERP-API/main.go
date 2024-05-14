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
	router.POST("/Login", handlers.Login)
	router.POST("/Register", handlers.Register)

	protectedRoutes := router.Group("/")
	protectedRoutes.Use(middleware.AuthenticationMiddleware())
	{
		//Subscriptions
		protectedRoutes.GET("/Subscriptions", handlers.GetSubscriptions)
		protectedRoutes.GET("/Subscriptions/:id", handlers.GetSubscriptionByName)
		//Contracts
		protectedRoutes.GET("/Contracts", handlers.GetContracts)
		protectedRoutes.GET("/Contracts/:id", handlers.GetContractById)
		protectedRoutes.POST("/Contracts", handlers.PostContract)
		//Products
		protectedRoutes.GET("/Products", handlers.GetProducts)
		protectedRoutes.GET("/Products/:id", handlers.GetProductByID)
		protectedRoutes.POST("/Products", handlers.PostProduct)
		//ProductionOrders
		protectedRoutes.GET("/ProductionOrders", handlers.GetProductionOrders)
		protectedRoutes.GET("/ProductionOrders/:id", handlers.GetProductionOrderByID)
		protectedRoutes.POST("/ProductionOrders", handlers.PostProductionOrder)
		//Deliveries
		protectedRoutes.GET("/Deliveries", handlers.GetDeliveries)
		protectedRoutes.GET("/Deliveries/:id", handlers.GetDeliveryByID)
		//Debtors
		protectedRoutes.GET("/Debtors", handlers.GetDebtors)
		protectedRoutes.GET("/Debtors/:id", handlers.GetDebtorByID)
		protectedRoutes.POST("/Debtors", handlers.PostDebtor)
	}

	router.Run("localhost:8080")
}
