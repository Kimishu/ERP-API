package enterprise

import (
	"ERP-API/models"
	"ERP-API/models/subscription"
	"fmt"
	"github.com/google/uuid"
)

type Enterprise struct {
	Id           uuid.UUID                 `json:"id"`
	Name         string                    `json:"name"`
	Email        string                    `json:"email"`
	Subscription subscription.Subscription `json:"subscription"`
}

func (e *Enterprise) Read(id string) Enterprise {
	var enterprise Enterprise
	err := models.Database.QueryRow("SELECT e.id, e.name, e.email, e.subscription_id, s.name FROM \"Enterprises\" e JOIN \"Subscriptions\" s ON e.subscription_id = s.id WHERE e.id = $1", id).Scan(&enterprise.Id, &enterprise.Name, &enterprise.Email, &enterprise.Subscription.Id, &enterprise.Subscription.Name)
	if err != nil {
		fmt.Println(err)
		return enterprise
	}
	return enterprise
}

func (e *Enterprise) ReadPartners(id string) []Enterprise {
	var partners []Enterprise
	rows, err := models.Database.Query("SELECT e.id, e.name FROM \"Partners\" INNER JOIN \"Enterprises\" e ON (\"Partners\".enterprise_first = e.id OR \"Partners\".enterprise_second = e.id) WHERE e.id != $1", id)

	if err != nil {
		fmt.Println(err)
		return []Enterprise{}
	}

	for rows.Next() {
		var partner Enterprise
		if err := rows.Scan(&partner.Id, &partner.Name); err != nil {
			fmt.Println(err)
			return []Enterprise{}
		}
		partners = append(partners, partner)
	}
	return partners
}
