package handlers

import (
	"ERP-API/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDebtors(c *gin.Context) {
	subRepo := models.Debtor{}
	enterpriseId := c.GetString("enterprise_id")

	debtors := subRepo.ReadByEnterprise(enterpriseId)
	c.JSON(http.StatusOK, debtors)
}

func GetDebtorByID(c *gin.Context) {

}

func PostDebtor(c *gin.Context) {

}
