package piscine

func Map(f func(int) bool, arr []int) []bool {
	counter := 0
	for i := range arr {
		_ = i
		counter++
	}
	a := make([]bool, counter)
	for i, v := range arr {
		a[i] = f(v)
	}
	return a
}
