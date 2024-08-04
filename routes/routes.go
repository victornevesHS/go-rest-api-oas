package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/victornevesHS/go-rest-api-oas/controllers"
)

func HnadleRequests() {
	r := mux.NewRouter() //criando instância do gorilla mux
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidades).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
