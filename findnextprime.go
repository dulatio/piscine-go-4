package piscine

func FindNextPrime(nb int) int {
	if nb < 0 {
		return 2
	}
	if nb%2 == 0 && nb > 3 {
		nb++
	}
	for {
		if isPrime(nb) {
			return nb
		}
		if nb == 1 {
			nb++
		} else {
			nb += 2
		}
	}
	return -1
}

func isPrime(nb int) bool {
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
