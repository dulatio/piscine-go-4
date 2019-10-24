package piscine

func AppendRange(min, max int) []int {
	a := []int(nil)
	for i := min; i < max; i++ {
		a = append(a, i)
	}
	return a
}
