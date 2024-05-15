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

func (p *ProductionOrderStatus) Read(id int) (*ProductionOrderStatus, error) {
	var productionOrderStatus ProductionOrderStatus

	err := Database.QueryRow("SELECT * FROM \"ProductionOrderStatuses\" WHERE id = ?", id).Scan(&productionOrderStatus.Id, &productionOrderStatus.Name)
	if err != nil {
		fmt.Println(err)
		return &productionOrderStatus, errors.New("order status not found")
	}
	return &productionOrderStatus, nil
}
