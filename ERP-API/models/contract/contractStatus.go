package contract

import (
	"ERP-API/models"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type ContractStatus struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (c *ContractStatus) ReadAll() []ContractStatus {
	var statuses []ContractStatus

	rows, err := models.Database.Query("SELECT id, name FROM \"ContractStatuses\"")
	if err != nil {
		fmt.Println(err)
		return statuses
	}

	for rows.Next() {
		var status ContractStatus
		if err := rows.Scan(&status.Id, &status.Name); err != nil {
			fmt.Println(err)
			return statuses
		}
		statuses = append(statuses, status)
	}

	return statuses
}

func (c *ContractStatus) Read(id string) ContractStatus {
	var contractStatus ContractStatus
	err := models.Database.QueryRow("SELECT * FROM \"ContractStatuses\" WHERE id=$1", id).Scan(&contractStatus.Id, &contractStatus.Name)
	if err != nil {
		fmt.Println(err)
		return contractStatus
	}
	return contractStatus
}

func (c *ContractStatus) Write() error {
	_, err := models.Database.Exec("INSERT INTO \"ContractStatuses\" (id, name) VALUES ($1, $2)", uuid.New(), &c.Name)
	if err != nil {
		return errors.New("Failed to create a new contract status")
	}
	return nil
}
