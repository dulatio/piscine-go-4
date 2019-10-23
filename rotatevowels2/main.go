package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	var sliceVowels []int
    arguments := os.Args
	k := 0
	nV := 0
	for index := range arguments {
		index = index
		k++
	}
	if k == 1 {
		z01.PrintRune('\n')
	} else {
		stroka := arguments[1]
		for i := 2; i < k; i++ {
			stroka = stroka + " " + arguments[i]
		}
		slice := []rune(stroka)
		slice0 := []rune(stroka)

		for index, letter := range slice {
			if isVowel(letter) == true {
				sliceVowels = append(sliceVowels, index)
				nV++
			}
		}
		if nV > 0 {
			for i := 1; i <= nV; i++ {
				slice[sliceVowels[i-1]] = slice0[sliceVowels[nV-i]]
			}
		}
		strokaNew := string(slice)
		for _, letter := range strokaNew {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}

func isVowel(r rune) bool{
	if r == 65 {
		return true
	}
	if r == 69 {
		return true
	}
	if r == 73 {
		return true
	}
	if r == 79 {
		return true
	}
	if r == 85 {
		return true
	}
	if r == 97 {
		return true
	}
	if r == 101 {
		return true
	}
	if r == 105 {
		return true
	}
	if r == 111 {
		return true
	}
	if r == 117 {
		return true
	}
	return false
}