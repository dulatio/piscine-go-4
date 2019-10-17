package piscine

func BasicAtoi(s string) int {
	var res int
	for _, r := range s {
		res *= 10
		res += int(r) - 48
	}
	return res
}
