package main

import (
	"context"
	_ "easycrm/loginit"
	"easycrm/models"
	"fmt"
	"github.com/jackc/pgx/pgxpool"
	"log"
)

func main() {
	fmt.Println("Hello I am new easyCRM-ADMIN")
	pool, err := pgxpool.Connect(context.Background(), `postgres://dsurush:dsurush@localhost:5432/bbunique?sslmode=disable`)
	if err != nil {
		log.Printf("Owibka - %e", err)
		log.Fatal("Can't Connection to DB")
	} else {
		log.Println("CONNECTION TO DB IS SUCCESS")
	}
	admin := models.Admin{
		FirstName: "surush",
		LastName:  "surush",
		UserName:  "surush",
		Password:  "surush",
	}
	admin.AddNew(pool)
}
