package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("--insert")
		fmt.Println("\t-i")
		fmt.Println("\t\tThis flag inserts the string into the string passed as argument.")
		fmt.Println("--order")
		fmt.Println("\t-o")
		fmt.Println("\t\tThis flag will behave like a boolean, if it is called it will order the argument.")
		return
	}
	toinsert := ""
	order := false
	s := ""
	for _, v := range os.Args {
		if len(v) > 9 && v[:9] == "--insert=" {
			toinsert = v[9:]
		} else if len(v) > 2 && v[:3] == "-i=" {
			toinsert = v[3:]
		} else if v == "--order" || v == "-o" {
			order = true
		} else if v == "-h" || v == "--help" {
			fmt.Println("--insert")
			fmt.Println("\t-i")
			fmt.Println("\t\tThis flag inserts the string into the string passed as argument.")
			fmt.Println("--order")
			fmt.Println("\t-o")
			fmt.Println("\t\tThis flag will behave like a boolean, if it is called it will order the argument.")
			return
		} else {
			s = v
		}
	}
	s += toinsert
	if order {
		runes := []rune(s)
		for i := 0; i < len(runes); i++ {
			for j := 0; j < len(runes)-1-i; j++ {
				if runes[j] > runes[j+1] {
					runes[j], runes[j+1] = runes[j+1], runes[j]
				}
			}
		}
		s = string(runes)
	}
	fmt.Println(s)
}
