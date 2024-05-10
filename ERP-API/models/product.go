package models

import "fmt"

type Product struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	E           *Enterprise `json:"enterprise"`
}

func (p *Product) Read(id int) *Product {
	var product Product
	var enterpriseId int
	err := Database.QueryRow("SELECT * FROM \"Products\" WHERE id=$1", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &enterpriseId)
	if err != nil {
		fmt.Println(err)
		return &product
	}

	var enterprise Enterprise
	product.E = enterprise.Read(enterpriseId)
	return &product
}
