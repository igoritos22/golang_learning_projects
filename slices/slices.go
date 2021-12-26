package main

import "fmt"

func main() {
	array:=[3]bool{true, true, false}
	fmt.Println(array)

	slice:=[]bool{true, true, false, false}
	fmt.Println(slice)
}