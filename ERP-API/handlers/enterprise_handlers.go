package handlers

import (
	"ERP-API/models"
	"ERP-API/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

//		Authorization example
//	1. Enterprise emails
//	2.
//

type loginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var data loginData

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	var userId, hashedPassword string
	err := models.Database.QueryRow("SELECT id, password FROM \"Enterprises\" WHERE email = $1", data.Email).Scan(&userId, &hashedPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(data.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid login credentials"})
		return
	}

	token, _ := utils.GenerateToken(userId)

	c.JSON(http.StatusOK, token)
}

type registerData struct {
	Ent      *models.Enterprise `json:"enterprise"`
	Password string             `json:"password"`
}

func Register(c *gin.Context) {
	var data registerData
	var userId string

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if ok := models.Database.QueryRow("SELECT id FROM \"Enterprises\" WHERE email = $1", data.Ent.Email).Scan(&userId); ok == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Enterprise with this email already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Hash error"})
		return
	}
	data.Password = string(hashedPassword)

	subscriptionId := data.Ent.Sub.ReadByName(data.Ent.Sub.Name).ID
	data.Ent.Sub.ID = subscriptionId

	_, err = models.Database.Exec("INSERT INTO \"Enterprises\" (name, email, password, subscription_id) VALUES ($1, $2, $3, $4)",
		data.Ent.Name, data.Ent.Email, data.Password, data.Ent.Sub.ID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = models.Database.QueryRow("SELECT id FROM \"Enterprises\" WHERE email = $1", data.Ent.Email).Scan(&userId)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Not found"})
		return
	}

	token, _ := utils.GenerateToken(userId)

	c.JSON(http.StatusOK, token)
}
