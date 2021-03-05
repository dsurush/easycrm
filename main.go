package main

import (
	_ "ccs/loginit"
	"context"
	"fmt"
	"github.com/jackc/pgx/pgxpool"
	"log"
)
func main() {
	fmt.Println("Hello I am new easyCRM-ADMIN")
	pool, err := pgxpool.Connect(context.Background(), `postgres://dsurush:dsurush@localhost:5432/ccs?sslmode=disable`)
	if err != nil {
		log.Printf("Owibka - %e", err)
		log.Fatal("Can't Connection to DB")
	} else {
		log.Println("CONNECTION TO DB IS SUCCESS")
	}
	
}
