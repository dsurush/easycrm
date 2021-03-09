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

func (receiver *CustomersSvc) AddNewCustomer(customer models.Customer) (err error) {
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

func (receiver *CustomersSvc) GetAllCustomers() (customers []models.Customer, err error){
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Println("can't get connection", err)
		return
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), `SELECT id, name, tin, address, ceo, enabled, removed_at, created_at, updated_at, balance FROM public.customers`)
	if err != nil {
		log.Printf("can't read user rows %e", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		Customer := models.Customer{}
		err := rows.Scan(
			&Customer.ID,
			&Customer.Name,
			&Customer.Tin,
			&Customer.Address,
			&Customer.CEO,

			&Customer.Enabled,
			&Customer.RemovedAt,
			&Customer.CreateAt,
			&Customer.UpdateAt,
			&Customer.Balance,
			)
		if err != nil {
			log.Println("can't scan err is = ", err)
			continue
		}
		customers = append(customers, Customer)
	}
	return
}