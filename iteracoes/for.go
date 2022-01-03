package main

import "fmt"

func main() {
	sum := 0

	for index := 0; index < 6; index++ {
		sum += index
	}

	fmt.Println("sum eh igual a: ", sum)
}
