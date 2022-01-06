# golang_learning_projects

## Testes e anotações do projeto go-rest-api

### Inicializar projetos:

go mod init <nome-do-seu-proejto>

### Importar modulos/pacotes exeternos
go mod init <nome-do-projeto>
go get <path-do-modulo> Ex: go get github.com/gorilla/mux

### Testar POST Json com Curl
```bash
curl -v -X POST http://localhost:3000/evento \
-H 'Content-Type: application-json' \
-d '{"ID":"4","Titulo":"SP open BJJ21","Descricao":"SP open de JiuJitsu","Preco":"250"}'
```
### Testar pegar um evento por ID
```bash
curl -v http://localhost:3000/buscaeventos/4
```
