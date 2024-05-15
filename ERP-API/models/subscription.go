package models

import "github.com/google/uuid"

type Subscription struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (s *Subscription) ReadAll() []Subscription {
	rows, err := Database.Query("SELECT id, name FROM \"Subscriptions\"")
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
	err := Database.QueryRow("SELECT id, name FROM \"Subscriptions\" WHERE name=$1", name).Scan(&subscription.Id, &subscription.Name)
	if err != nil {
		return Subscription{}
	}
	return subscription
}

func (s *Subscription) Read(id string) Subscription {
	var subscription Subscription
	err := Database.QueryRow("SELECT id, name FROM \"Subscriptions\" WHERE id=$1", id).Scan(&subscription.Id, &subscription.Name)
	if err != nil {
		return Subscription{}
	}
	return subscription
}
