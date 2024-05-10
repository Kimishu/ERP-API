package handlers

import (
	"ERP-API/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetContracts(c *gin.Context) {
	subRepo := &models.Contract{}
	contracts, _ := subRepo.Read(2, "seller_id")
	c.JSON(http.StatusOK, contracts)
}

func GetContractById(c *gin.Context) {

}

func PostContract(c *gin.Context) {

}
