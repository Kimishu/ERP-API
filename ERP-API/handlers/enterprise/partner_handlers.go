package enterprise

import (
	"ERP-API/models/enterprise"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPartners(c *gin.Context) {
	subRepo := &enterprise.Enterprise{}
	enterpriseId := c.GetString("enterprise_id")
	partners := subRepo.ReadPartners(enterpriseId)
	c.JSON(http.StatusOK, partners)
}

func GetIncomingPartnerRequests(c *gin.Context) {
	subRepo := &enterprise.PartnerRequest{}
	enterpriseId := c.GetString("enterprise_id")
	incomingRequests := subRepo.ReadIncoming(enterpriseId)
	c.JSON(http.StatusOK, incomingRequests)
}

func GetOutgoingPartnerRequests(c *gin.Context) {
	subRepo := &enterprise.PartnerRequest{}
	enterpriseId := c.GetString("enterprise_id")
	outgoingRequests := subRepo.ReadOutgoing(enterpriseId)
	c.JSON(http.StatusOK, outgoingRequests)
}

func PostPartnerRequest(c *gin.Context) {
	partnerRequest := &enterprise.PartnerRequest{}

	if err := c.ShouldBindJSON(&partnerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := partnerRequest.Write()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Partner request created"})
}

func AcceptPartnerRequest(c *gin.Context) {
	partnerRequest := &enterprise.PartnerRequest{}

	if err := c.ShouldBindJSON(&partnerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := partnerRequest.Accept(c.GetString("enterprise_id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Partner request accepted"})
}

func DeclinePartnerRequest(c *gin.Context) {
	partnerRequest := &enterprise.PartnerRequest{}

	if err := c.ShouldBindJSON(&partnerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := partnerRequest.Decline(c.GetString("enterprise_id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Partner request declined"})
}
