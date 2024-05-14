package handlers

import (
	"ERP-API/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSubscriptions(c *gin.Context) {
	subRepo := &models.Subscription{}
	subscriptions := subRepo.ReadAll()
	c.JSON(http.StatusOK, subscriptions)
}

func GetSubscriptionByName(c *gin.Context) {
	subRepo := &models.Subscription{}
	name := c.Param("name")
	subscription := subRepo.ReadByName(name)
	c.JSON(http.StatusOK, subscription)
}
