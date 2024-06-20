package enterprise

import "github.com/google/uuid"

type Debtor struct {
	Id           *uuid.UUID `json:"id"`
	EnterpriseId *uuid.UUID `json:"enterprise_id"`
	ContractId   *uuid.UUID `json:"contract_id"`
	Sum          float64    `json:"sum"`
}

func (d *Debtor) ReadByEnterprise(enterpriseId uuid.UUID) []Debtor {
	var debtors []Debtor

	return debtors
}

func (d *Debtor) Read(id uuid.UUID) Debtor {
	var debtor Debtor

	return debtor
}

func (d *Debtor) Write() error {

	return nil
}
