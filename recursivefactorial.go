package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 20{
		return 0
	}
	if nb < 2 {
		return 1
	}
	res := nb * RecursiveFactorial(nb-1)
	return res
}
