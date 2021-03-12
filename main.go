package main

import (
	"context"
	"easycrm/cmd/app"
	_ "easycrm/loginit"
	"easycrm/pkg/core/services"
	"easycrm/token"
	"fmt"
	"github.com/jackc/pgx/pgxpool"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func main() {

	//pool, err := pgxpool.Connect(context.Background(), `postgres://dsurush:dsurush@localhost:5432/bbunique?sslmode=disable`)
	pool, err := pgxpool.Connect(context.Background(), `postgres://dsurush:dsurush@localhost:5000/bbunique?sslmode=disable`)
	if err != nil {log.Fatalf("Can't Connection to DB %e", err)
	time.Sleep(time.Second*5)}
	log.Println("CONNECTION TO DB IS SUCCESS")
	fmt.Println("CONNECTION TO DB IS SUCCESS")

	router := httprouter.New()
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			w.Header().Set("Content-Type", "application/json, text/html")
			//			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, RefreshToken")
			w.Header().Set("Accept", "*/*")
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
	tokenSvc := token.NewTokenSvc([]byte("My Secret Key"), pool)
	userSvc := services.NewUserSvc(pool)
	customersSvc := services.NewCustomersSvc(pool)
	server := app.NewMainServer(router, pool, userSvc, tokenSvc, customersSvc)
	server.Start()
}
