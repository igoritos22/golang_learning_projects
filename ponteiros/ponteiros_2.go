package main

import "fmt"

func soma1(a *int) int { //recebe um ponteiro no argumento
	*a = *a + 1
	return *a
}

func main() {
	x := 3

	fmt.Print("o valor atual de x eh: \n", x)

	x1 := soma1(&x) // passa X para a funcao soma1(). O cracter & indica a referencia da posição da memoria dessa vaiavel
	fmt.Println("\n x + 1 eh: ", x1)

	//valor novo de x
	fmt.Print("\n o novo valor de x eh: \n", x)

}
