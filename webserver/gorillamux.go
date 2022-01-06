package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Meu app Go com Gorilla mux!")

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", homePage)
	log.Println("API esta rodando")
	http.ListenAndServe(":10000", router)
}
