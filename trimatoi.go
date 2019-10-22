package piscine

func TrimAtoi(s string) int {
	beforesign := true
	neg := false
	res := 0
	for _, r := range s {
		if r == '-' && beforesign {
			neg = true
		}
		if r >= '0' && r <= '9' {
			res *= 10
			res += int(r) - '0'
			beforesign = false
		}
	}
	if neg == true {
		return -res
	}
	return res
}
