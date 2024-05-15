package models

import "github.com/google/uuid"

type Debtor struct {
	Id           uuid.UUID `json:"id"`
	EnterpriseId uuid.UUID `json:"enterprise_id"`
	ContractId   uuid.UUID `json:"contract_id"`
	Sum          float64   `json:"sum"`
}

func (d *Debtor) ReadByEnterprise(enterpriseId string) []Debtor {
	var debtors []Debtor

	return debtors
}
