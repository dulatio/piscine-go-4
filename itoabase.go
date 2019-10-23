package piscine

func ItoaBase(nbr int, base string) string {
	counter := 0
	for _, r := range base {
		counter++
		_ = r
	}
	res := ""
	ra := []rune(base)
	for ; nbr != 0; nbr /= counter {
		index := nbr % counter
		res = string(ra[index]) + res
	}
	return res
}
