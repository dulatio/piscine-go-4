package main

import (
	"github.com/01-edu/z01"
	piscine ".."
)

func main() {
	piscine.PrintNbrBase(125, "012345678")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(125, "0123456789ABCDEFghij")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(-125, "ch")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(125, "awert-q")
	z01.PrintRune('\n')
	//piscine.PrintNbrBase2(-9223372036854775808, "0123456789")
	//z01.PrintRune('\n')
	piscine.PrintNbrBase(-9223372036854775808, "0123456789")
	z01.PrintRune('\n')
}
