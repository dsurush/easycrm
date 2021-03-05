package services

import (
	"errors"
	"github.com/jackc/pgx/pgxpool"
)

type UserSvc struct {
	*pgxpool.Pool
}

func NewUserSvc(pool *pgxpool.Pool) *UserSvc {
	if pool == nil {
		panic(errors.New("pool can't be nil")) // <- be accurate
	}
	return &UserSvc{Pool: pool}
}
