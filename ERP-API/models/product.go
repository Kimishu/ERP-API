package models

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type Product struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	EnterpriseId uuid.UUID `json:"enterprise"`
}

func (p *Product) Read(id string) Product {
	var product Product
	err := Database.QueryRow("SELECT name, description, price, enterprise_id FROM \"Products\" WHERE id=$1", id).Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.EnterpriseId)
	if err != nil {
		fmt.Println(err)
		return product
	}
	return product
}

func (p *Product) ReadByEnterprise(enterpriseId string) []Product {
	var products []Product
	rows, err := Database.Query("SELECT name, description, price FROM \"Products\" WHERE enterprise_id = $1", enterpriseId)
	if err != nil {
		fmt.Println(err)
		return products
	}

	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.Name, &product.Description, &product.Price); err != nil {
			fmt.Println(err)
			return products
		}
		products = append(products, product)
	}

	return products
}

func (p *Product) Write() error {
	_, err := Database.Exec("INSERT INTO \"Products\" (id, name, description, price, enterprise_id) VALUES ($1, $2, $3, $4, $5)", uuid.New(), &p.Name, &p.Description, &p.Price, &p.EnterpriseId)
	if err != nil {
		fmt.Println(err)
		return errors.New("Failed to create a new product")
	}
	return nil
}
