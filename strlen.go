package piscine

func StrLen(str string) int {
	counter := 0
	for _, r := range str {
		counter++
		_ = r
	}
	return counter
}
