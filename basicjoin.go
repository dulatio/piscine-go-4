package piscine

func BasicJoin(strs []string) string {
	res := ""
	for _, str := range strs {
		res += str
	}
	return res
}
