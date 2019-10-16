package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	// using int32 instead of int is more convenient because rune and int 32 can be used 'interchangeably'
	var n32 int32
	
	// create array with size 10 since max n is 9
	a := [...]rune{48, 48, 48, 48, 48, 48, 48, 48, 48, 48}

	// bruteforcing because we cannot use casting
	if n == 1 {
		n32 = 1
	} else if n == 2 {
		n32 = 2
	} else if n == 3 {
		n32 = 3
	} else if n == 3 {
		n32 = 3
	} else if n == 4 {
		n32 = 4
	} else if n == 5 {
		n32 = 5
	} else if n == 6 {
		n32 = 6
	} else if n == 7 {
		n32 = 7
	} else if n == 8 {
		n32 = 8
	} else {
		n32 = 9
	}

	// i is going to be temporary index for coming loops (we need it to be int32 instead of int)
	var i int32
	for i = 0; i < n32; i++ {
		a[i] = i + 48
	}

	// index is going to be the main index to navigate through the array (we need it to be int32 instead of int)
	var index int32
	// setting index to the last digit
	index = n32 - 1
	
	// upper limit for each index: 10 - n + index

	// starting an unconditional loop 
	for ; ; {
		// printing the combination
		for i = 0; i < n32; i++ {
			z01.PrintRune(a[i])
		}

		// if we reached the final combination stop the main loop
		if a[0] == 48 + 10 - n32 {
			break
		}

		// if we did not reach the final combination append ', '
		z01.PrintRune(',')
		z01.PrintRune(' ')

		// here THE MAGIC happens
		// it is triggered by the last digit being 9
		a[index]++;
		if a[index] == (48 + 10 - n32 + index) + 1 {
			// shift index to the left		
			for ; index >= 0 && a[index] >= (48 + 10 - n32 + index); index-- {}
			a[index]++;

			// shift index to the right
			for index = index + 1; index < n32; index++ {
				a[index] = a[index - 1] + 1
			}
			// setting index to the last digit again 
			index = n32 - 1
		}
	}

	// putting newline at the end of the function
	// because examples in the task seem to have it
	z01.PrintRune(10)
}