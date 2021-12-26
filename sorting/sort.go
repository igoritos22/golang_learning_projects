package main

import (
	"fmt"
	"sort"
)

func main() {
	ints:=[]int{2, 4, 6, 8, 10}

	sort.Ints(ints)
	fmt.Println("Inteiros:  ", ints)
}