package piscine

func ToLower(s string) string {
	res := ""
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			res += string(r - ('A' - 'a'))
		} else {
			res += string(r)
		}
	}
	return res
}
