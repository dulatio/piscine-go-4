package main

import (
	"bufio"
	"github.com/01-edu/z01"
	"io/ioutil"
	"os"
)

func main() {
	counter := 0
	for i, v := range os.Args {
		counter++
		if i != 0 {
			data, err := ioutil.ReadFile(v)
			if err != nil {
				printString(err.Error())
			}
			printString(string(data))
			z01.PrintRune(10)
		}
	}
	if counter == 1 {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		printString(text)
	}
}

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
