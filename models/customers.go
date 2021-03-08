package models

import (
	"github.com/google/uuid"
	"time"
)

type Customer struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Tin       string    `json:"tin"`
	Address   string    `json:"address"`
	CEO       string    `json:"ceo"`
	Enabled   bool      `json:"enabled"`
	RemovedAt time.Time `json:"removed_at"`
	CreateAt  time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
	Balance   int64     `json:"balance"`
}
