package handlers

import (
	"ERP-API/models"
	"github.com/gin-gonic/gin"
)

func GetContracts(c *gin.Context) {
	db := models.Database{}
	db.Connect()
	defer db.CloseConnection()

}

func GetContractById(c *gin.Context) {

}

func PostContract(c *gin.Context) {

}
