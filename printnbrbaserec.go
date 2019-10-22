package piscine

import "github.com/01-edu/z01"

func PrintNbrBaseRec(n int, runes []rune, len int) {
	if n/len != 0 {
		PrintNbrBaseRec(n/len, runes, len)
	}
	z01.PrintRune(runes[n%len])
}