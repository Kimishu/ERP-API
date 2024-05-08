package models

import "time"

type ProductionOrder struct {
	ID        int                   `json:"id"`
	Status    ProductionOrderStatus `json:"status"`
	DateStart time.Time             `json:"date_start"`
	DateEnd   time.Time             `json:"date_end"`
	Prod      Product               `json:"product"`
}

type ProductionOrderStatus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
