package handlers

import (
	"ERP-API/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProducts(c *gin.Context) {
	subRepo := &models.Product{}
	enterpriseId := c.GetString("enterprise_id")

	products := subRepo.ReadByEnterprise(enterpriseId)
	c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	subRepo := &models.Product{}
	product := subRepo.Read(c.Param("id"))
	c.JSON(http.StatusOK, product)
}

func PostProduct(c *gin.Context) {
	var product = models.Product{
		E: models.Enterprise{
			ID: c.GetString("enterprise_id"),
		},
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := product.Write()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Product created"})
}
