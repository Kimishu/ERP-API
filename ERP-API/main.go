package main

import (
	"ERP-API/handlers"
	"ERP-API/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDB()
	defer models.Database.Close()

	router := gin.Default()
	//Subscriptions
	router.GET("/Subscriptions", handlers.GetSubscriptions)
	router.GET("/Subscriptions/:id", handlers.GetSubscriptionById)
	//Contracts
	router.GET("/Contracts", handlers.GetContracts)
	router.GET("/Contracts/:id", handlers.GetContractById)
	router.POST("/Contracts", handlers.PostContract)
	//Products
	router.GET("/Products", handlers.GetProducts)
	router.GET("/Products/:id", handlers.GetProductByID)
	router.POST("/Products", handlers.PostProduct)
	//ProductionOrders
	router.GET("/ProductionOrders", handlers.GetProductionOrders)
	router.GET("/ProductionOrders/:id", handlers.GetProductionOrderByID)
	router.POST("/ProductionOrders", handlers.PostProductionOrder)
	//Deliveries
	router.GET("/Deliveries", handlers.GetDeliveries)
	router.GET("/Deliveries/:id", handlers.GetDeliveryByID)
	//Debtors
	router.GET("/Debtors", handlers.GetDebtors)
	router.GET("/Debtors/:id", handlers.GetDebtorByID)
	router.POST("/Debtors", handlers.PostDebtor)

	router.Run("localhost:8080")
}
