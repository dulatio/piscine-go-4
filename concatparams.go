package piscine

func ConcatParams(args []string) string {
	res := ""
	for i, v := range args {
		if i != 0 {
			res += "\n"
		}
		res += v
	}
	return res
}
