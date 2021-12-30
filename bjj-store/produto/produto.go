package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

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

func main() {

	rota := mux.NewRouter()
	rota.HandleFunc("/produtos", ListaProduto)
	http.ListenAndServe(":8081", rota)

	fmt.Println(string(loadData()))
}
