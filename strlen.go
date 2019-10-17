package piscine

func StrLen(str string) int {
	counter := 0
	for c := range str {
		counter++
	}
	if counter == 0 {
		return 0
	} else {
		return counter + 1
	}
}
