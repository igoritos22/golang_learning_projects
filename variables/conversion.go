//convertendo variavel do tip string numca variavel slice

package main

import "fmt"

func main() {
	s := "hello"
	c := []byte(s) //converte string para o tipo byte

	c[0] = 'i'
	c[1] = 'g'
	c[2] = 'o'
	c[3] = 'r'

	igor := string(c) //converte o slice de volta numa string porem com os indices alterados

	a := s + igor //concatenando variaveis

	fmt.Printf("%s\n", igor)
	fmt.Printf("juntando as duas variaveis sai: %s", a)
}
