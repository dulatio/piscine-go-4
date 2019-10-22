package piscine

func AtoiBase(s string, base string) int {
	lenbase := 0
	for _, r1 := range base {
		flag := 0
		if r1 == '+' || r1 == '-' {
			return 0
		}
		for _, r2 := range base {
			if r1 == r2 {
				flag++
			}
			if flag == 2 {
				return 0
			}
		}
		lenbase++
	}
	if lenbase < 2 {
		return 0
	}
	lens := 0
	for _, r := range s {
		lens++
		_ = r
	}
	as := []rune(s)
	ab := []rune(base)
	for i, j := 0, lens-1; i < j; i, j = i+1, j-1 {
		as[i], as[j] = as[j], as[i]
	}
	res := 0
	for i := 0; i < lens; i++ {
		index := -1
		for j := 0; j < lenbase; j++ {
			if as[i] == ab[j] {
				index = j
			}
		}
		res += index * recursivePower(lenbase, i)
	}
	return res
}

func recursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	return nb * RecursivePower(nb, power-1)
}
