package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//Definicao da struct de produtos
type Produto struct {
	UUID   string  `json:"uuid"`
	Name   string  `json:"name"`
	Detail string  `json:"detail"`
	Price  float64 `json:"price,string"`
}

//Definicao da minha colecao de produtos
type Produtos struct {
	Produtos []Produto
}

//Carrega os produtos do arquivo produto.json
func loadData() []byte {
	jsonFile, err := os.Open("produtos.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	return data
}

//Recebe os produtos da funcao loadData e responde atraves das requisicoes na rota /produtos
func ListaProduto(writer http.ResponseWriter, req *http.Request) {
	produtos := loadData()
	writer.Write([]byte(produtos))
}

func ProdutosPorId(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	data := loadData()

	var produtos Produtos
	json.Unmarshal(data, &produtos)

	for _, valor := range produtos.Produtos {

		if valor.UUID == vars["id"] {
			produto, _ := json.Marshal(valor)
			writer.Write([]byte(produto))
		}
	}
}

func main() {

	rota := mux.NewRouter()
	rota.HandleFunc("/produtos", ListaProduto)
	rota.HandleFunc("/produto/{id}", ProdutosPorId)
	http.ListenAndServe(":8081", rota)

	fmt.Println(string(loadData()))
}
