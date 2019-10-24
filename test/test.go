package main

import (
	"fmt"
	piscine ".."
)

func main() {
	tab1 := []int{0, 1, 2, 3, 4, 5}
	tab2 := []int{0, 2, 1, 3}

	result1 := piscine.IsSorted(piscine.CompareInts, tab1)
	result2 := piscine.IsSorted(piscine.CompareInts, tab2)

	fmt.Println(result1)
	fmt.Println(result2)
}
