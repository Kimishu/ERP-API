package models

import (
	"errors"
	"fmt"
)

type Product struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	E           Enterprise `json:"enterprise"`
}

func (p *Product) Read(id string) Product {
	var product Product
	var enterpriseId string
	err := Database.QueryRow("SELECT name, description, price, enterprise_id FROM \"Products\" WHERE id=$1", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &enterpriseId)
	if err != nil {
		fmt.Println(err)
		return product
	}

	var enterprise Enterprise
	product.E = enterprise.Read(enterpriseId)
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
	_, err := Database.Exec("INSERT INTO \"Products\" (name, description, price, enterprise_id) VALUES ($1, $2, $3, $4)", &p.Name, &p.Description, &p.Price, &p.E.ID)
	if err != nil {
		fmt.Println(err)
		return errors.New("Failed to create a new product")
	}
	return nil
}
