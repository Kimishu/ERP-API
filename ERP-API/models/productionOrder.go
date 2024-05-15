package models

import (
	"github.com/google/uuid"
	"time"
)

type ProductionOrder struct {
	Id        uuid.UUID `json:"id"`
	StatusId  uuid.UUID `json:"status"`
	DateStart time.Time `json:"date_start"`
	DateEnd   time.Time `json:"date_end"`
	ProductId uuid.UUID `json:"product"`
}
