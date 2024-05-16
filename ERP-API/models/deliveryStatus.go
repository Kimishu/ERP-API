package models

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type DeliveryStatus struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (d *DeliveryStatus) ReadAll() []DeliveryStatus {
	var statuses []DeliveryStatus

	rows, err := Database.Query("SELECT id, name FROM \"DeliveryStatuses\"")
	if err != nil {
		fmt.Println(err)
		return []DeliveryStatus{}
	}

	for rows.Next() {
		var status DeliveryStatus
		if err := rows.Scan(&status.Id, &status.Name); err != nil {
			fmt.Println(err)
			return []DeliveryStatus{}
		}
		statuses = append(statuses, status)
	}

	return statuses
}

func (d *DeliveryStatus) Read(statusId uuid.UUID) DeliveryStatus {
	var status DeliveryStatus

	err := Database.QueryRow("SELECT id, name FROM \"DeliveryStatuses\" WHERE id = $1", statusId).Scan(&status.Id, &status.Name)
	if err != nil {
		fmt.Println(err)
		return DeliveryStatus{}
	}

	return status
}

func (d *DeliveryStatus) Write() error {
	_, err := Database.Exec("INSERT INTO \"DeliveryStatuses\" (id, name) VALUES ($1, $2)", uuid.New(), d.Name)
	if err != nil {
		return errors.New("Failed to create a new delivery status!")
	}
	return nil
}
