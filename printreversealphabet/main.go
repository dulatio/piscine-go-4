package main

import "fmt"

func main() {
	for c := 'z'; c >= 'a'; c-- {
		fmt.Printf("%c", c)
	}
	fmt.Println();
}
