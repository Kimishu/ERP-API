package main

import (
	"ERP-API/handlers/contract"
	"ERP-API/handlers/delivery"
	"ERP-API/handlers/enterprise"
	"ERP-API/handlers/product"
	"ERP-API/handlers/productionOrder"
	"ERP-API/handlers/subscription"
	"ERP-API/middleware"
	"ERP-API/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDB()
	defer models.Database.Close()

	router := gin.Default()

	//Enterprises (public)
	router.POST("/login", enterprise.Login)
	router.POST("/register", enterprise.Register)
	//Subscriptions
	router.GET("/subscriptions", subscription.GetSubscriptions)
	router.GET("/subscriptions/:id", subscription.GetSubscriptionByName)
	//Contract Statuses
	router.GET("/contracts/statuses", contract.GetContractStatuses)
	router.GET("/contracts/statuses/:id", contract.GetContractStatusById)
	//Delivery Statuses
	router.GET("/deliveries/statuses", delivery.GetDeliveryStatuses)
	router.GET("/deliveries/statuses/:id", delivery.GetDeliveryStatusById)

	protectedRoutes := router.Group("/")
	protectedRoutes.Use(middleware.AuthenticationMiddleware())
	{
		//Enterprises
		protectedRoutes.GET("/profile", enterprise.Profile)
		protectedRoutes.GET("/enterprise/:id", enterprise.Enterprise)
		//Subscriptions
		protectedRoutes.POST("/subscriptions", subscription.PostSubscription)
		//Contracts
		protectedRoutes.GET("/contracts", contract.GetContracts)
		protectedRoutes.GET("/contracts/import", contract.GetImportContracts)
		protectedRoutes.GET("/contracts/export", contract.GetExportContracts)
		//Contract Statuses
		protectedRoutes.POST("/contracts/statuses", contract.PostContractStatus)
		//requests?
		//...
		protectedRoutes.GET("/contracts/:id", contract.GetContractById)
		protectedRoutes.POST("/contracts", contract.PostContract)
		//...
		//Products
		protectedRoutes.GET("/products", product.GetProducts)
		protectedRoutes.GET("/products/:id", product.GetProductByID)
		protectedRoutes.POST("/products", product.PostProduct)
		//ProductionOrders
		protectedRoutes.GET("/productionOrders", productionOrder.GetProductionOrders)
		protectedRoutes.GET("/productionOrders/:id", productionOrder.GetProductionOrderByID)
		protectedRoutes.POST("/productionOrders", productionOrder.PostProductionOrder)
		//Deliveries
		protectedRoutes.GET("/deliveries", delivery.GetDeliveries)
		protectedRoutes.GET("/deliveries/:id", delivery.GetDeliveryByID)
		//Delivery Statuses
		protectedRoutes.POST("/deliveries/statuses", delivery.PostDeliveryStatus)
		//Debtors
		protectedRoutes.GET("/debtors", enterprise.GetDebtors)
		protectedRoutes.GET("/debtors/:id", enterprise.GetDebtorByID)
		protectedRoutes.POST("/debtors", enterprise.PostDebtor)
		//Partners
		protectedRoutes.GET("/partners", enterprise.GetPartners)
		protectedRoutes.GET("/partners/incoming", enterprise.GetIncomingPartnerRequests)
		protectedRoutes.GET("/partners/outgoing", enterprise.GetOutgoingPartnerRequests)
		protectedRoutes.POST("/partners", enterprise.PostPartnerRequest)
		protectedRoutes.POST("/accept-partner", enterprise.AcceptPartnerRequest)
		protectedRoutes.POST("/decline-partner", enterprise.DeclinePartnerRequest)
	}

	router.Run("localhost:8080")
}
