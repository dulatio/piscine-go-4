package piscine

func Sqrt(nb int) int {
	for i := 1; ; i++ {
		if i * i == nb {
			return i
		}
		if i * i > nb {
			return 0
		}
	}
}