package models

type Debtor struct {
	ID  int        `json:"id"`
	E   Enterprise `json:"e"`
	C   Contract   `json:"c"`
	Sum float64    `json:"sum"`
}
