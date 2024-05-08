package handlers

import (
	"ERP-API/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetSubscriptions(c *gin.Context) {
	subRepo := &models.Subscription{}
	subscriptions := subRepo.ReadAll()
	c.JSON(http.StatusOK, subscriptions)
}

func GetSubscriptionById(c *gin.Context) {
	subRepo := &models.Subscription{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	subscription := subRepo.Read(id)
	c.JSON(http.StatusOK, subscription)
}
