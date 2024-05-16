package models

import (
	"fmt"
	"github.com/google/uuid"
)

type Enterprise struct {
	Id             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	SubscriptionId uuid.UUID `json:"subscription_id"`
}

func (e *Enterprise) Read(id string) Enterprise {
	var enterprise Enterprise
	err := Database.QueryRow("SELECT id, name, email, subscription_id FROM \"Enterprises\" WHERE id = $1", id).Scan(&enterprise.Id, &enterprise.Name, &enterprise.Email, &enterprise.SubscriptionId)
	if err != nil {
		fmt.Println(err)
		return enterprise
	}
	return enterprise
}
