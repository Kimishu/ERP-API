package models

import (
	"errors"
	"fmt"
)

type Contract struct {
	ID          int             `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Quantity    int             `json:"quantity"`
	Prod        *Product        `json:"product_name"`
	Status      *ContractStatus `json:"status"`
	Seller      *Enterprise     `json:"seller"`
	Buyer       *Enterprise     `json:"buyer"`
}

func (c *Contract) Read(id int, queryParam string) ([]Contract, error) {
	var contracts []Contract

	//rows, err := Database.Query("SELECT * FROM \"Contracts\" WHERE $1 = $2", queryParam, id)
	rows, err := Database.Query("SELECT * FROM \"Contracts\" WHERE seller_id = 2")
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

		var productId, statusId, sellerId, buyerId int
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

func (c *Contract) Write() {
	fmt.Println("WRITE CONTRACT")
}
