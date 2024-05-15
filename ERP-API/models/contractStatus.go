package models

import "fmt"

type ContractStatus struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *ContractStatus) Read(id string) ContractStatus {
	var contractStatus ContractStatus
	err := Database.QueryRow("SELECT * FROM \"ContractStatuses\" WHERE id=$1", id).Scan(&contractStatus.ID, &contractStatus.Name)
	if err != nil {
		fmt.Println(err)
		return contractStatus
	}
	return contractStatus
}
