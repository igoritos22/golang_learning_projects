package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
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
		crutime := time.Now().Unix()
		hash := md5.New()

		io.WriteString(hash, strconv.FormatInt(crutime, 10))

		token := fmt.Sprintf("%x", hash.Sum(nil))

		t, _ := template.ParseFiles("webserver/login.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")

		if token != "" {
			//
		} else {
			//
		}

		fmt.Println("username:", r.Form.Get("username"))
		fmt.Println("password:", r.Form.Get("password"))
		//fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		//fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		fmt.Println("POST content: ", r.Form.Encode())
		//template.HTMLEscape(w, []byte(r.Form.Get("username")))
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
