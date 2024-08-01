package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/victornevesHS/go-rest-api-oas/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)

}
