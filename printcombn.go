package piscine

import "fmt"

func PrintCombN(n int) {
	a := make([]int, n)

	for i, _ := range a {
		a[i] = i
	}

	// index := n - 1

	for ; ; {
		for _,v := range a {
			fmt.Printf("%d", v)
		}
	}
}
