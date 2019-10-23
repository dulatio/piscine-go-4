package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	res := ""
	for i, v := range os.Args {
		if i == 0 {
			continue
		}
		if i != 1 {
			res += " "
		}
		res += v
	}
	runes := []rune(res)
	i := 0
	j := runesLen(runes) - 1
	for i < j {
		if isVowel(runes[i]) && isVowel(runes[j]) {
			runes[i], runes[j] = runes[j], runes[i]
			i++
			j--
		} else if isVowel(runes[i]) {
			j--
		} else if isVowel(runes[j]) {
			i++
		} else {
			i++
			j--
		}
	}
	res = string(runes)
	for _, r := range res {
		z01.PrintRune(r)
	}
	z01.PrintRune(10)
}

func runesLen(a []rune) int {
	counter := 0
	for _, r := range a {
		counter++
		_ = r
	}
	return counter
}

func isVowel(r rune) bool {
	if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' {
		return true
	}
	return false
}
