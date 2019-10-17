package piscine

func StrRev(s string) string {
	olds := []byte(s)
	news := []byte(s)
	counter := 0
	for i := range olds {
		counter++
		_ = i
	}
	for i := 0; i < counter; i++ {
		news[i] = olds[counter-1-i]
	}
	return string(news)
}
