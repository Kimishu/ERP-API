package subscription

import (
	"ERP-API/models"
	"errors"
	"github.com/google/uuid"
)

type Subscription struct {
	Id   *uuid.UUID `json:"id"`
	Name string     `json:"name"`
}

func (s *Subscription) ReadAll() []Subscription {
	rows, err := models.Database.Query("SELECT id, name FROM \"Subscriptions\"")
	if err != nil {
		return []Subscription{}
	}
	defer rows.Close()

	var subscriptions []Subscription
	for rows.Next() {
		var s Subscription
		if err := rows.Scan(&s.Id, &s.Name); err != nil {
			return []Subscription{}
		}
		subscriptions = append(subscriptions, s)
	}
	return subscriptions
}

func (s *Subscription) ReadByName(name string) Subscription {
	var subscription Subscription
	err := models.Database.QueryRow("SELECT id, name FROM \"Subscriptions\" WHERE name=$1", name).Scan(&subscription.Id, &subscription.Name)
	if err != nil {
		return Subscription{}
	}
	return subscription
}

func (s *Subscription) Read(id string) Subscription {
	var subscription Subscription
	err := models.Database.QueryRow("SELECT id, name FROM \"Subscriptions\" WHERE id=$1", id).Scan(&subscription.Id, &subscription.Name)
	if err != nil {
		return Subscription{}
	}
	return subscription
}

func (s *Subscription) Write() error {
	_, err := models.Database.Exec("INSERT INTO \"Subscriptions\" (id, name) VALUES ($1, $2)", uuid.New(), &s.Name)
	if err != nil {
		return errors.New("Failed to create a new subscription")
	}
	return nil
}
