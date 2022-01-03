package main

import "fmt"

func soma1(a int) int {
	a = a + 1
	return a
}

func main() {
	x := 3

	fmt.Print("o valor atual de x eh: \n", x)

	x1 := soma1(x) // passa X para a funcao soma1()
	fmt.Println("\n x + 1 eh: ", x1)

	//valor novo de x
	fmt.Print("\n o novo valor de x eh: \n", x)

}

//mesmo passando X para funco soma1() o valor permanece o valor da declaração da funcao main
//isso acontece pq a função soma1() recebeu uma copia da variavel X
//nesse caso se faz necessário o uso de ponteiros
//veja esse mesmo exemplo no arquivo ponteiros_2.go na raiz desta pasta
