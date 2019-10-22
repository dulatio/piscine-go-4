package piscine

func Capitalize(s string) string {
	res := ""
	flag := true
	for _, r := range s {
		if flag {
			if r >= 'a' && r <= 'z' {
				res += string(r + ('A' - 'a'))
				flag = false
			} else if (r >= '0' && r <= '9') || (r >= 'A' && r <= 'Z') {
				res += string(r)
				flag = false
			} else {
				res += string(r)
				flag = true
			}
		} else {
			if r >= 'A' && r <= 'Z' {
				res += string(r - ('A' - 'a'))
			} else if (r >= '0' && r <= '9') || (r >= 'a' && r <= 'z') {
				res += string(r)
			} else {
				res += string(r)
				flag = true
			}
		}
	}
	return res
}
