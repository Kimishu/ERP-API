package models

import "fmt"

type Enterprise struct {
	ID    string       `json:"id"`
	Name  string       `json:"name"`
	Email string       `json:"email"`
	Sub   Subscription `json:"subscription"`
}

func (e *Enterprise) Read(id string) Enterprise {
	var enterprise Enterprise
	var subscriptionId string
	err := Database.QueryRow("SELECT id, name, email, subscription_id FROM \"Enterprises\" WHERE id = $1", id).Scan(&enterprise.ID, &enterprise.Name, &enterprise.Email, &subscriptionId)
	if err != nil {
		fmt.Println(err)
		return enterprise
	}
	var subscription Subscription
	enterprise.Sub = subscription.ReadById(subscriptionId)
	return enterprise
}
