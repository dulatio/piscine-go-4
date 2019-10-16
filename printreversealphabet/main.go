package main

import "github.com/01-edu/z01"

func main() {
	//for c := 'z'; c >= 'a'; c-- {
	//	z01.PrintRune(c)
	//}
	//z01.PrintRune('\n');
	//z01.PrintRune('\n');
	alph := "zyxwvutsrqponmlkjihgfedcba\n"
	for i := 0; i < 27; i++ {
		z01.PrintRune(rune(alph[i]))
	}
}
