package services

import (
	"context"
	"easycrm/models"
	"github.com/jackc/pgx/pgxpool"
	"log"
)

type CustomersSvc struct {
	pool *pgxpool.Pool
}

func NewCustomersSvc(pool *pgxpool.Pool) *CustomersSvc {
	return &CustomersSvc{pool: pool}
}

func (receiver *CustomersSvc) Add(customer models.Customer) (err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Println("can't get connection", err)
		return
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), `Insert into "customers"(id, name, tin, address, ceo, enabled, removed_at, created_at, updated_at, balance) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		customer.ID,
		customer.Name,
		customer.Tin,
		customer.Address,
		customer.CEO,
		customer.Enabled,
		customer.RemovedAt,
		customer.CreateAt,
		customer.UpdateAt,
		customer.Balance,
		)
	if err != nil {
		log.Println("can't add edit User StateDML = ", err)
		return
	}
	return
}
