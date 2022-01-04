package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

//Função para tratar o formulário de requisição enviado pelo user
func TratarDadosDoForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //Tratar parametros passados via URL. Caso não sejam tratados os dados nao pderao ser obtidos via Go
	fmt.Println(r.Form) //Apresentar informações do formuçario na saída do servidor
	fmt.Println("O PATH da requisicao eh: ", r.URL.Path)
	fmt.Println("O SCHEMA da requiscao eh:", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("Chave: ", k)
		fmt.Println("Valor: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "\n Ola cliente.. essa eh uma simples app em GOlangs!") //Dados que serao enviados para o cliente
}

//Função para tratar e receber os parametros enviados via pagina de Login
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("metodo da requisicao: ", r.Method) //apresentar o metodo da requição
	if r.Method == "GET" {
		t, _ := template.ParseFiles("webserver/login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.FormValue("username"))
		fmt.Println("password:", r.FormValue("password"))
		fmt.Println("POST content: ", r.Form.Encode())
	}
}

//Função principal
func main() {
	http.HandleFunc("/", TratarDadosDoForm)
	http.HandleFunc("/login", Login)
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

//testar a req na pagina de login
//curl -d "password=mypassword&username=adminmster" -X POST http://localhost:9090/login
