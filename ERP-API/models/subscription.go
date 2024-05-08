package models

type Subscription struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (s *Subscription) ReadAll() []Subscription {
	db := Database{}
	db.Connect()
	defer db.CloseConnection()

	rows, err := db.Conn.Query("SELECT id, name FROM \"Subscriptions\"")
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

func (s *Subscription) Read(id int) Subscription {
	db := Database{}
	db.Connect()
	defer db.CloseConnection()

	var subscription Subscription

	err := db.Conn.QueryRow("SELECT id, name FROM \"Subscriptions\" WHERE id=$1", id).Scan(&subscription.ID, &subscription.Name)

	if err != nil {
		return Subscription{}
	}

	return subscription
}
