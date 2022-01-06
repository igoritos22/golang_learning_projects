package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type evento struct {
	ID          string `json:"ID"`
	Titulo      string `json:"Titulo"`
	Descricacao string `json:"Descricacao"`
	Preco       string `json:"price,string"`
}

type TodosEventos []evento

var eventos = TodosEventos{
	{
		ID:          "1",
		Titulo:      "Sao Paulo Open 2022",
		Descricacao: "Evento de JiuJitsu da federação CBJJ",
		Preco:       "200",
	},
}

func criarEvento(w http.ResponseWriter, r *http.Request) {
	//variavel para receber o novo evento
	var novoEvento evento
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Informe o titulo, descrição e preco do novo evento: \n")
	}

	json.Unmarshal(reqBody, &novoEvento)  //desorquestra os dados enviados via usuario
	eventos = append(eventos, novoEvento) //realiza a mescla/concatena o slice eventos com o novoEvento adicionado
	w.WriteHeader(http.StatusCreated)     //retorna o status 201 com a criação do novo Evento

	json.NewEncoder(w).Encode(novoEvento)

}

func pegarEventoPorId(w http.ResponseWriter, r *http.Request) {
	//variavel para receber o ID informado peo user
	IdEvento := mux.Vars(r)["id"] //mux.Vars recebe as vaiaveis passadas via requisição (r http.Request)

	for _, eventoUnico := range eventos {
		if eventoUnico.ID == IdEvento {
			json.NewEncoder(w).Encode(eventoUnico)
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo a home page")
}

func main() {
	router := mux.NewRouter().StrictSlash(true) //StrictSlash(true) caso o path for /path e o user digitar /path/ a aplicação redireciona para o path correto
	router.HandleFunc("/", homePage)
	router.HandleFunc("/evento", criarEvento).Methods("POST")
	router.HandleFunc("/buscaeventos/{id}", pegarEventoPorId).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}
