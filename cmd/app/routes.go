package app

import (
	"easycrm/models"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

func (server *MainServer) InitRoutes(){
	fmt.Println("Init routes")
	test(server)
	//ЛОГИН
	server.router.POST("/api/login", server.LoginHandler)

	//КЛИЕНТЫ
	server.router.POST("/api/customers/add", server.AddCustomerHandler)
	server.router.GET("/api/customers", server.GetAllCustomersHandler)


	log.Println(http.ListenAndServe(":8888", server))
}

func test(server *MainServer){
	a := models.Customer{
		Name:      "Name",
		Tin:       "Name",
		Address:   "Name",
		CEO:       "Name",
		Enabled:   false,
		RemovedAt: time.Now(),
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
		Balance:   0,
	}
	a.ID = uuid.New()
	err := server.CustomerSvc.AddNewCustomer(a)
	if err != nil {
		fmt.Println("err ", err)
		return
	}
}