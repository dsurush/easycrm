package main

import (
	"context"
	"easycrm/cmd/app"
	_ "easycrm/loginit"
	"easycrm/pkg/core/services"
	"easycrm/token"
	"github.com/jackc/pgx/pgxpool"
	"github.com/julienschmidt/httprouter"
	"log"
)

func main() {

	pool, err := pgxpool.Connect(context.Background(), `postgres://dsurush:dsurush@localhost:5432/bbunique?sslmode=disable`)
	if err != nil {log.Fatalf("Can't Connection to DB %e", err)}
	log.Println("CONNECTION TO DB IS SUCCESS")

	router := httprouter.New()
	tokenSvc := token.NewTokenSvc([]byte("My Secret Key"), pool)
	userSvc := services.NewUserSvc(pool)
	customersSvc := services.NewCustomersSvc(pool)
	server := app.NewMainServer(router, pool, userSvc, tokenSvc, customersSvc)
	server.Start()
}
