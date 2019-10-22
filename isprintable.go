package piscine

func IsPrintable(str string) bool {
	for _, r := range str {
		if r >= 0 && r <= 31 {
			return false
		}
	}
	return true
}
