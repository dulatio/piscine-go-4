package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	last := 0
	for i, v := range tab {
		if i != 0 {
			if f(last, v) > 0 {
				return false
			}
		}
		last = v
	}
	return true
}
