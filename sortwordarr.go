package piscine

func SortWordArr(array []string) {
	counter := 0
	for i := range array {
		counter++
		_ = i
	}
	for i := 0; i < counter; i++ {
		for j := 0; j < counter-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}