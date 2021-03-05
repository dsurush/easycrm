package models

import "github.com/jackc/pgtype"

type Admin struct {
	ID pgtype.UUID `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
