package piscine

func Index(s string, toFind string) int {
	ar1 := []rune(s)
	ar2 := []rune(toFind)
	count1 := 0
	for _, r := range ar1 {
		count1++
		_ = r
	}
	count2 := 0
	for _, r := range ar2 {
		count2++
		_ = r
	}
	for i := 0; i < count1; i++ {
		j, k := i, 0
		for ; j < count1 && k < count2; {
			if ar1[j] != ar2[k] {
				break
			}
			j++
			k++
		}
		if k == count2 {
			return i
		}
	}
	return -1
}
