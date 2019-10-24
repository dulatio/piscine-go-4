package piscine

func Split(str, charset string) []string {
	i, j := 0, 0
	index := 0
	res := ""
	words := make([]string, StrLen(str))
	for i < StrLen(str) {
		if j == StrLen(charset) {
			words[index] = res[:StrLen(res)-StrLen(charset)]
			index++
			res = ""
			j = 0
		}
		if str[i] == charset[j] {
			j++
		} else {
			j = 0
		}
		res += string(str[i])
		i++
	}
	if StrLen(res) > 0 {
		words[index] = res
		index++
	}
	words = words[:index]
	return words
}
