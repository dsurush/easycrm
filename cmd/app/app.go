package app

import (
	"easycrm/pkg/core/services"
	"easycrm/token"
	"github.com/jackc/pgx/pgxpool"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type MainServer struct {

	router *httprouter.Router
	pool *pgxpool.Pool
	svc *services.UserSvc
	tokenSvc *token.TokenSvc
}

func NewMainServer(router *httprouter.Router, pool *pgxpool.Pool, svc *services.UserSvc, tokenSvc *token.TokenSvc) *MainServer {
	return &MainServer{router: router, pool: pool, svc: svc, tokenSvc: tokenSvc}
}

func (server *MainServer) Start() {
	server.InitRoutes()

}

func (server *MainServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	server.router.ServeHTTP(writer, request)
}
