package models

import (
	"errors"
	"fmt"
)

type Contract struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Quantity    int            `json:"quantity"`
	Prod        Product        `json:"product_name"`
	Status      ContractStatus `json:"status"`
	Seller      Enterprise     `json:"seller"`
	Buyer       Enterprise     `json:"buyer"`
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
		var product Product
		var status ContractStatus
		var seller Enterprise
		var buyer Enterprise

		var productId, statusId, sellerId, buyerId string
		if err := rows.Scan(&contract.ID, &contract.Name, &contract.Description, &contract.Quantity, &productId, &statusId, &sellerId, &buyerId); err != nil {
			fmt.Println(err)
			return []Contract{}, errors.New("empty contracts")
		}
		contract.Prod = product.Read(productId)
		contract.Status = status.Read(statusId)
		contract.Seller = seller.Read(sellerId)
		contract.Buyer = buyer.Read(buyerId)

		contracts = append(contracts, contract)
	}

	return contracts, nil
}

func (c *Contract) Write() error {
	_, err := Database.Exec("INSERT INTO \"Contracts\" (name, description, quantity, product_id, status_id, seller_id, buyer_id) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		&c.Name, &c.Description, &c.Quantity, &c.Prod.ID, &c.Status.ID, &c.Seller.ID, &c.Buyer.ID)
	if err != nil {
		fmt.Println(err)
		return errors.New("Failed to create a new contract")
	}
	return nil
}
