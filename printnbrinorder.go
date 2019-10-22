package piscine

func PrintNbrInOrder(n int) {
	for ; n > 0; n /= 10 {
		z01.PrintRune(rune(n%10))
	}
}
