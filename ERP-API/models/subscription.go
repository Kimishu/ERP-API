package models

type Subscription struct {
	ID   string `json:"id"`
	Name string `json:"name"`
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
		if err := rows.Scan(&s.ID, &s.Name); err != nil {
			return []Subscription{}
		}
		subscriptions = append(subscriptions, s)
	}

	return subscriptions
}

func (s *Subscription) ReadByName(name string) Subscription {
	var subscription Subscription

	err := Database.QueryRow("SELECT id, name FROM \"Subscriptions\" WHERE name=$1", name).Scan(&subscription.ID, &subscription.Name)
	if err != nil {
		return Subscription{}
	}

	return subscription
}

func (s *Subscription) ReadById(id string) Subscription {
	var subscription Subscription

	err := Database.QueryRow("SELECT id, name FROM \"Subscriptions\" WHERE id=$1", id).Scan(&subscription.ID, &subscription.Name)
	if err != nil {
		return Subscription{}
	}

	return subscription
}
