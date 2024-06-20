package productionOrder

import (
	"ERP-API/models"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type ProductionOrder struct {
	Id        uuid.UUID `json:"id"`
	StatusId  uuid.UUID `json:"status"`
	DateStart time.Time `json:"date_start"`
	DateEnd   time.Time `json:"date_end"`
	Quantity  uint      `json:"quantity"`
	ProductId uuid.UUID `json:"product"`
}

func (po *ProductionOrder) Read(enterpriseId uuid.UUID, orderId uuid.UUID) (ProductionOrder, error) {
	var productionOrder ProductionOrder
	err := models.Database.QueryRow("SELECT po.id, po.status_id, po.date_start, po.date_end, po.quantity, po.product_id FROM \"ProductionOrders\" po JOIN \"Products\" p ON po.product_id = $1 AND p.enterprise_id = $2", orderId, enterpriseId).Scan(&productionOrder.Id,
		&productionOrder.StatusId, &productionOrder.DateStart, &productionOrder.DateEnd, &productionOrder.Quantity, &productionOrder.ProductId)
	if err != nil {
		return ProductionOrder{}, errors.New("Requested order doesn't exist or doesn't belongs to you")
	}
	return productionOrder, nil
}

func (po *ProductionOrder) ReadAll(enterpriseId uuid.UUID) []ProductionOrder {
	var productionOrders []ProductionOrder
	rows, err := models.Database.Query("SELECT id, status_id, date_start, date_end, quantity, product_id FROM \"ProductionOrders\""+
		"WHERE product_id = (SELECT product_id FROM \"Products\" WHERE enterprise_id = $1)", enterpriseId)
	if err != nil {
		fmt.Println(err)
		return productionOrders
	}

	for rows.Next() {
		var productionOrder ProductionOrder
		if err := rows.Scan(&productionOrder.Id, &productionOrder.StatusId, &productionOrder.DateStart, &productionOrder.DateEnd, &productionOrder.Quantity, &productionOrder.ProductId); err != nil {
			fmt.Println(err)
			return productionOrders
		}
		productionOrders = append(productionOrders, productionOrder)
	}

	return productionOrders
}

func (po *ProductionOrder) Write() error {
	po.DateStart = time.Now()
	po.DateEnd = po.DateStart.Add(time.Hour * 24 * 10)
	//uuid.New()
	_, err := models.Database.Exec("INSERT INTO \"ProductionOrders\" (id, status_id, date_start, date_end, quantity, product_id)")
	if err != nil {
		return errors.New("Failed to create a new production order!")
	}
	return nil
}
