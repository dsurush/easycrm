package app

import (
	"fmt"
	"log"
	"net/http"
)

func (server *MainServer) InitRoutes(){
	fmt.Println("Init routes")
	//test(server)
	server.router.POST("/api/login", server.LoginHandler)
	log.Println(http.ListenAndServe(":8888", server))
}
