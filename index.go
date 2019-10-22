package piscine

func Index(s string, toFind string) int {
	index := -1
	rtofind := firstRune(toFind)
	for _, r := range s {
		index++
		if r == rtofind {
			return index
		}
	}
	return -1
}

func firstRune(s string) rune {
	for _, r := range s {
		return r
	}
	return '\x00'
}
