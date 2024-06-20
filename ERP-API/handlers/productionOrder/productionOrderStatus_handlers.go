package productionOrder

import (
	"ERP-API/models/productionOrder"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func GetProductionOrderStatuses(c *gin.Context) {
	subRepo := &productionOrder.ProductionOrderStatus{}
	statuses := subRepo.ReadAll()
	c.JSON(http.StatusOK, statuses)
}

func GetProductionOrderStatusById(c *gin.Context) {
	subRepo := &productionOrder.ProductionOrderStatus{}
	id, _ := uuid.Parse(c.Param("id"))
	status, err := subRepo.Read(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, status)
}

func PostProductionOrderStatus(c *gin.Context) {
	productionOrderStatus := &productionOrder.ProductionOrderStatus{}

	if err := c.ShouldBindJSON(&productionOrderStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := productionOrderStatus.Write()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Production order status created"})
}
