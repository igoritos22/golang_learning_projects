package main

import (
	"fmt"
	"mathigor"
)
func main() {
	resultado := Soma(10, 5)
	fmt.Printf("Funcao interna - O valor da soma eh: %v  \n", resultado)

	resultFuncDiff := math.Diff(10, 5)
	fmt.Printf("Funcao externa - O valor da subtracao eh: %v  \n", resultFuncDiff)

	resultFuncDiv := math.Div(10, 5)
	fmt.Printf("Funcao externa - O valor da divisao eh: %v  \n", resultFuncDiv)

	resultFuncProd := math.Product(10, 5)
	fmt.Printf("Funcao externa - O valor da multiplicacao eh: %v  \n", resultFuncProd)
}
func Soma(a int, b int) int {
	return a + b
}