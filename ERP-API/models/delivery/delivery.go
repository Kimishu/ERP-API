package delivery

import (
	"github.com/google/uuid"
	"time"
)

type Delivery struct {
	Id         uuid.UUID `json:"id"`
	ContractId uuid.UUID `json:"contract_id"`
	Quantity   uint      `json:"quantity"`
	date       time.Time `json:"date"`
	statusId   uuid.UUID `json:"status_id"`
}
