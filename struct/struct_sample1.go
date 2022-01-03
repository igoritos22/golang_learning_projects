//construir uma struct, utilizar os dados definidos nela.
//construir uma funcao que utiiza os dados da struct

package main

import "fmt"

//define a struct

type pessoa struct {
	nome  string
	idade int
}

//funcao recebe idade de duas pessoas, compara e retorna quem eh mais velho
func MaisVelho(p1, p2 pessoa) (pessoa, int) {
	if p1.idade > p2.idade {
		return p1, p1.idade - p2.idade
	} else {
		return p2, p2.idade - p1.idade
	}
}

func main() {
	var igor pessoa
	igor.nome, igor.idade = "Igor Luiz", 29

	tamara := pessoa{idade: 31, nome: "Tamara"}

	mais_velho, idade_dif := MaisVelho(igor, tamara)

	fmt.Printf("A pessoa mais velha eh: %v e a diferença de idades eh: %v anos", mais_velho, idade_dif)
}
