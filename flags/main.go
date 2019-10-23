package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("--insert")
		fmt.Println("\t-i")
		fmt.Println("\t\tThis flag inserts the string into the string passed as argument.")
		fmt.Println("--order")
		fmt.Println("\t-o")
		fmt.Println("\t\tThis flag will behave like a boolean, if it is called it will order the argument.")
	}
	toinsert := ""
	order := false
	for i, v := range os.Args {
		if (len(v) > 9 && v[:9] == "--insert=") || (len) {
			toinsert = v[9:]
		} else if len(v) == 7 && v == "--order" {
			order = true
		}
	}
}
