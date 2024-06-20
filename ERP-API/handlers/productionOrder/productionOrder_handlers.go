package productionOrder

import (
	"ERP-API/models/productionOrder"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func GetProductionOrders(c *gin.Context) {
	subRepo := productionOrder.ProductionOrder{}
	enterpriseId, _ := uuid.Parse(c.GetString("enterprise_id"))
	productionOrders := subRepo.ReadAll(enterpriseId)

	c.JSON(http.StatusOK, productionOrders)
}

func GetProductionOrderByID(c *gin.Context) {
	subRepo := productionOrder.ProductionOrder{}
	enterpriseId, _ := uuid.Parse(c.GetString("enterprise_id"))
	orderId, _ := uuid.Parse(c.Param("id"))
	productionOrder, _ := subRepo.Read(enterpriseId, orderId)
	c.JSON(http.StatusOK, productionOrder)
}

func PostProductionOrder(c *gin.Context) {
	productionOrder := productionOrder.ProductionOrder{}

	if err := c.ShouldBindJSON(&productionOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := productionOrder.Write()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Production order created"})
}
