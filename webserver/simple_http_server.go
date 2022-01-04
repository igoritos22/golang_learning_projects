package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func apresentarDadosNaWeb(wr http.ResponseWriter, req *http.Request) { //parametros wr para retornar dados do lado do server e req para receber os dados das requisições
	req.ParseForm()                                           //Parsear os argumentos da requisição
	fmt.Println(req.Form)                                     //Apresenta o formulario enviado pelo cliente na requisção no lado do server
	fmt.Println("o PATH da solicitacao eh", req.URL.Path)     //apresena o path no lado do server
	fmt.Println("o HOST da solicitacao eh", req.URL.Host)     //apresena o host da solicitação no lado do server
	fmt.Println("o SCHEMA da solicitacao eh", req.URL.Scheme) //apresena o schema da solicitação no lado do server
	fmt.Println(req.Form["url_long"])

	for k, v := range req.Form {
		fmt.Println("Chave: ", k)
		fmt.Println("Valor", strings.Join(v, ""))
	}

	fmt.Fprintf(wr, "\n Ola cliente.. essa eh uma simples app em GOlangs!") //Dados que serao enviados para o cliente
}
func main() {
	http.HandleFunc("/", apresentarDadosNaWeb)
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe:  ", err)
	}
}
