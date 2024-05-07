package main

import (
	"ERP-API/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Subscription struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func getSubscriptions(c *gin.Context) {

	rows, err := database.Connect().Query("SELECT id, name FROM \"Subscriptions\"")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	var subscriptions []Subscription
	for rows.Next() {
		var s Subscription
		if err := rows.Scan(&s.Id, &s.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		subscriptions = append(subscriptions, s)
	}

	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subscriptions)
}

func getSubscriptionById(c *gin.Context) {

}

func getContracts(c *gin.Context) {

}

func getContractById(c *gin.Context) {

}

func main() {
	router := gin.Default()
	router.GET("/Subscriptions", getSubscriptions)
	router.GET("/Subscriptions/:id", getSubscriptionById)
	router.GET("/Contracts", getContracts)
	router.GET("/Contracts/:id", getContractById)
	router.Run("localhost:8080")
}
