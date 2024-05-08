package models

type Contract struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Quantity    int            `json:"quantity"`
	Prod        Product        `json:"product_name"`
	Status      ContractStatus `json:"status"`
	Seller      Enterprise     `json:"seller"`
	Buyer       Enterprise     `json:"buyer"`
}

type ContractStatus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
