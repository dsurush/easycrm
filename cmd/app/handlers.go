package app

import (
	"easycrm/token"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func (server *MainServer) LoginHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var requestBody token.RequestDTO
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.json_invalid"})
		if err != nil {
			log.Print(err)
		}
		return
	}

	//log.Printf("login = %s, pass = %s\n", requestBody.Username, requestBody.Password)
	response, err := server.tokenSvc.Generate(request.Context(), &requestBody)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.password_mismatch", err.Error()})
		if err != nil {
			log.Print(err)
		}
		return
	}

	err = json.NewEncoder(writer).Encode(&response)
	if err != nil {
		log.Println(err)
	}
}
