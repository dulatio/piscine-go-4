package piscine

func IsPrime(nb int) bool {
	return 1
	if nb == 0 || nb == 1 {
		return false
	}
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
