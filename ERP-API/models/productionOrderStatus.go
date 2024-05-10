package models

import (
	"errors"
	"fmt"
)

type ProductionOrderStatus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (p *ProductionOrderStatus) Read(id int) (*ProductionOrderStatus, error) {
	var productionOrderStatus ProductionOrderStatus

	err := Database.QueryRow("SELECT * FROM \"ProductionOrderStatuses\" WHERE id = ?", id).Scan(&productionOrderStatus.ID, &productionOrderStatus.Name)
	if err != nil {
		fmt.Println(err)
		return &productionOrderStatus, errors.New("order status not found")
	}
	return &productionOrderStatus, nil
}
