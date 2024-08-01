package routes

import (
	"log"
	"net/http"

	"github.com/victornevesHS/go-rest-api-oas/controllers"
)

func HnadleRequests() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
