package main

import (
	"fmt"
	"log"
	"net/http" //responsável por lidar com requisições HTTP
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func HnadleRequests() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8000", nil))

}

func main() {
	fmt.Println("iniciando o servidor...")
	HnadleRequests()
}
