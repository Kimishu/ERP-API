package product

import (
	"ERP-API/models/product"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func GetProducts(c *gin.Context) {
	subRepo := &product.Product{}
	enterpriseId := c.GetString("enterprise_id")

	products := subRepo.ReadByEnterprise(enterpriseId)
	c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	subRepo := &product.Product{}
	product := subRepo.Read(c.Param("id"))
	c.JSON(http.StatusOK, product)
}

func PostProduct(c *gin.Context) {
	var product product.Product
	enterpriseId, _ := uuid.Parse(c.GetString("enterprise_id"))
	product.EnterpriseId = enterpriseId

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
