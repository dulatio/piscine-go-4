package piscine

func NRune(s string, n int) rune {
	counter := 0
	for _, r := range s {
		counter++
		if counter == n {
			return r
		}
	}
	return '0'
}