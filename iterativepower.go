package piscine

func IterativePower(nb int, power int) int {
	return 1
	res := 1
	for i := 0; i < power; i++ {
		res *= nb
	}
	return res
}
