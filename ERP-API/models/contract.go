package models

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type Contract struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	ProductId   uuid.UUID `json:"product_name"`
	StatusId    uuid.UUID `json:"status"`
	SellerId    uuid.UUID `json:"seller"`
	BuyerId     uuid.UUID `json:"buyer"`
}

func (c *Contract) Read(id string, queryParam string) ([]Contract, error) {
	var contracts []Contract

	rows, err := Database.Query("SELECT * FROM \"Contracts\" WHERE $1 = $2", queryParam, id)
	if err != nil {
		return contracts, errors.New("enterprise hasn't any contracts")
	}

	defer rows.Close()

	for rows.Next() {
		var contract Contract

		if err := rows.Scan(&contract.Id, &contract.Name, &contract.Description, &contract.Quantity, &contract.ProductId, &contract.StatusId, &contract.SellerId, &contract.BuyerId); err != nil {
			fmt.Println(err)
			return []Contract{}, errors.New("empty contracts")
		}
		contracts = append(contracts, contract)
	}

	return contracts, nil
}

func (c *Contract) Write() error {
	_, err := Database.Exec("INSERT INTO \"Contracts\" (name, description, quantity, product_id, status_id, seller_id, buyer_id) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		&c.Name, &c.Description, &c.Quantity, &c.ProductId, &c.StatusId, &c.SellerId, &c.BuyerId)
	if err != nil {
		fmt.Println(err)
		return errors.New("Failed to create a new contract")
	}
	return nil
}
