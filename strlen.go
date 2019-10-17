package piscine

func StrLen(str string) int {
	counter := 0
	for i := range str {
		counter = i
	}
	if counter == 0 {
		return 0
	} else {
		return counter + 1
	}
}
