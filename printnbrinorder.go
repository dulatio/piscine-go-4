package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var ar []int
	count := 0
	for ; n > 0; n /= 10 {
		ar = append(ar, n%10)
		count++
	}
	for i := 1; i < count; i++ {
		for j := 0; j < count-i; j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
	}
	for i := 0; i < count; i++ {
		z01.PrintRune(rune(ar[i]) + '0')
	}
}
