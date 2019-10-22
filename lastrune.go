package piscine

func LastRune(s string) rune {
	var lr rune
	for _, r := range s {
		lr = r
	}
	return lr
}
