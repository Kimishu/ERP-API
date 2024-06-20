package subscription

import (
	"ERP-API/models/subscription"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSubscriptions(c *gin.Context) {
	subRepo := &subscription.Subscription{}
	subscriptions := subRepo.ReadAll()
	c.JSON(http.StatusOK, subscriptions)
}

func GetSubscriptionByName(c *gin.Context) {
	subRepo := &subscription.Subscription{}
	name := c.Param("name")
	subscription := subRepo.ReadByName(name)
	c.JSON(http.StatusOK, subscription)
}

func PostSubscription(c *gin.Context) {
	subscription := &subscription.Subscription{}
	if err := c.ShouldBindJSON(subscription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := subscription.Write()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Subscription created"})
}
