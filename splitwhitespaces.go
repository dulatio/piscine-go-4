package piscine

func SplitWhiteSpaces(str string) []string {
	counter := 0
	for _, r := range str {
		if r == ' ' || r == '\t' || r == '\n' {
			counter++
		}
	}
	a := make([]string, counter + 2)
	word := ""
	index := 0
	for _, r := range str {
		if r == ' ' || r == '\t' || r == '\n' {
			if word != "" {
				a[index] = word
				index++
				word = ""
			}
		} else {
			word += string(r)
		}
	}
	if word != "" {
		a[index] = word
		index++
	}
	b := make([]string, index)
	for i := 0; i < index; i++ {
		b[i] = a[i]
	}
	return b
}