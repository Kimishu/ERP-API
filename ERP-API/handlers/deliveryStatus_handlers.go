package handlers

import (
	"ERP-API/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func GetDeliveryStatuses(c *gin.Context) {
	subRepo := &models.DeliveryStatus{}
	statuses := subRepo.ReadAll()
	c.JSON(http.StatusOK, statuses)
}

func GetDeliveryStatusById(c *gin.Context) {
	subRepo := &models.DeliveryStatus{}
	statusId, _ := uuid.Parse(c.Param("id"))
	status := subRepo.Read(statusId)
	c.JSON(http.StatusOK, status)
}

func PostDeliveryStatus(c *gin.Context) {
	deliveryStatus := &models.DeliveryStatus{}

	if err := c.ShouldBindJSON(&deliveryStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := deliveryStatus.Write()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Delivery status created"})
}
