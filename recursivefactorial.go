package piscine

func RecursiveFactorial(nb int) int {
	if nb < 2 {
		return 1
	}
	return nb * RecursiveFactorial(nb - 1)
}
