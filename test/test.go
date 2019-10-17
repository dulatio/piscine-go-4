package main

import (
	"fmt"
	piscine ".."
)

func main(){
	s := []int{-12,-124,24,4}
	piscine.SortIntegerTable(s)
	fmt.Println(s)
}
