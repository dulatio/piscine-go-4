package piscine

func BasicAtoi2(s string) int {
	var res int
	for _, r := range s {
		res *= 10
		if r < 48 || r > 57 {
			return 0
		}
		if r == 48 {
			res += 0
		} else if r == 49 {
			res += 1
		} else if r == 50 {
			res += 2
		} else if r == 51 {
			res += 3
		} else if r == 52 {
			res += 4
		} else if r == 53 {
			res += 5
		} else if r == 54 {
			res += 6
		} else if r == 55 {
			res += 7
		} else if r == 56 {
			res += 8
		} else if r == 57 {
			res += 9
		}
	}
	return res
}
