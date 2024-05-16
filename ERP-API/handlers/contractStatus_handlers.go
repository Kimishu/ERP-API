package handlers

import (
	"ERP-API/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetContractStatuses(c *gin.Context) {
	subRepo := models.ContractStatus{}
	contractStatuses := subRepo.ReadAll()
	c.JSON(http.StatusOK, contractStatuses)
}

func GetContractStatusById(c *gin.Context) {
	subRepo := models.ContractStatus{}
	contractStatus := subRepo.Read(c.Param("id"))
	c.JSON(http.StatusOK, contractStatus)
}

func PostContractStatus(c *gin.Context) {
	status := &models.ContractStatus{}

	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := status.Write()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Contract status created"})
}
