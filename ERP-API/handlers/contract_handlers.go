package handlers

import (
	"ERP-API/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetContracts(c *gin.Context) {
	subRepo := &models.Contract{}
	enterpriseId := c.GetString("enterprise_id")

	importContracts, _ := subRepo.Read(enterpriseId, "buyer_id")
	exportContracts, _ := subRepo.Read(enterpriseId, "seller_id")

	c.JSON(http.StatusOK, gin.H{
		"importContracts": importContracts,
		"exportContracts": exportContracts,
	})
}

func GetImportContracts(c *gin.Context) {
	subRepo := &models.Contract{}
	enterpriseId := c.GetString("enterprise_id")
	contracts, _ := subRepo.Read(enterpriseId, "buyer_id")
	c.JSON(http.StatusOK, contracts)
}

func GetExportContracts(c *gin.Context) {
	subRepo := &models.Contract{}
	enterpriseId := c.GetString("enterprise_id")
	contracts, _ := subRepo.Read(enterpriseId, "seller_id")
	c.JSON(http.StatusOK, contracts)
}

func GetContractById(c *gin.Context) {

}

// Необходимо добавить обработку нахождения предприятия, с которым заключается контракт
func PostContract(c *gin.Context) {
	var contract = models.Contract{}
}
