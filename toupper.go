package piscine

func ToUpper(s string) string {
	res := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			res += string(r + ('A' - 'a'))
		} else {
			res += string(r)
		}
	}
	return res
}
