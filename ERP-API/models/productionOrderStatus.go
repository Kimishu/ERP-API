package models

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type ProductionOrderStatus struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (p *ProductionOrderStatus) ReadAll() []ProductionOrderStatus {
	var statuses []ProductionOrderStatus
	rows, err := Database.Query("SELECT id, name FROM \"PrudctionsOrderStatuses\"")
	if err != nil {
		return []ProductionOrderStatus{}
	}

	for rows.Next() {
		var status ProductionOrderStatus
		if err := rows.Scan(&status.Id, &status.Name); err != nil {
			return []ProductionOrderStatus{}
		}
		statuses = append(statuses, status)
	}
	return statuses
}

func (p *ProductionOrderStatus) Read(id uuid.UUID) (ProductionOrderStatus, error) {
	var productionOrderStatus ProductionOrderStatus

	err := Database.QueryRow("SELECT * FROM \"ProductionOrderStatuses\" WHERE id = ?", id).Scan(&productionOrderStatus.Id, &productionOrderStatus.Name)
	if err != nil {
		fmt.Println(err)
		return productionOrderStatus, errors.New("order status not found")
	}
	return productionOrderStatus, nil
}

func (p *ProductionOrderStatus) Write() error {
	_, err := Database.Exec("INSERT INTO \"ProductionOrderStatuses\" (id, name) VALUES ($1, $2)", uuid.New(), &p.Name)
	if err != nil {
		return errors.New("Failed to create a new production order status!")
	}
	return nil
}
