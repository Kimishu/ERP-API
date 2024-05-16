package handlers

import (
	"ERP-API/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func GetDebtors(c *gin.Context) {
	subRepo := models.Debtor{}
	enterpriseId, _ := uuid.Parse(c.GetString("enterprise_id"))
	debtors := subRepo.ReadByEnterprise(enterpriseId)
	c.JSON(http.StatusOK, debtors)
}

func GetDebtorByID(c *gin.Context) {
	subRepo := models.Debtor{}
	enterpriseId, _ := uuid.Parse(c.Param("id"))
	debtor := subRepo.Read(enterpriseId)
	c.JSON(http.StatusOK, debtor)
}

func PostDebtor(c *gin.Context) {
	debtor := models.Debtor{}

	if err := c.ShouldBindJSON(&debtor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := debtor.Write()

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "New debtor has been created"})
}
