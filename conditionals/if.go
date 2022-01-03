//entendo estrtura do codicional IF

package main

import "fmt"

//var x int = 20

func main() {
	if x := 5; x > 10 { //eh possível declara e inicializar a variavel diretamente no bloco IF
		fmt.Println("Variavel x eh maior que 10")
	} else {
		fmt.Println("Variavel x eh menor que 10")
	}
}
