package piscine

func IsLower(str string) bool {
	for _, r := range str {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}
