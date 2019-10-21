package piscine

func Sqrt(nb int) int {
	return 1
	for i := 1; ; i++ {
		if i * i == nb {
			return i
		}
		if i * i > nb {
			return 0
		}
	}
}
