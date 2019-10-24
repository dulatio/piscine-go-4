package main

import (
	"fmt"
	"os"
)

func main() {
	counter := 0
	astr := ""
	bstr := ""
	operator := ""
	for i, v := range os.Args {
		if i == 1 {
			astr = v
		}
		if i == 2 {
			operator = v
		}
		if i == 3 {
			bstr = v
		}
		counter++
	}
	if counter != 4 {
		return
	}
	if !isNumeric(astr) || !isNumeric(bstr) {
		fmt.Println(0)
		return
	}
	a := atoi(astr)
	b := atoi(bstr)
	if operator == "+" {
		fmt.Println(a+b)
	} else if operator == "-" {
		fmt.Println(a-b)
	} else if operator == "*" {
		fmt.Println(a*b)
	} else if operator == "/" {
		if b == 0 {
			fmt.Println("No division by 0")
			return
		}
		fmt.Println(a/b)
	} else if operator == "%" {
		if b == 0 {
			fmt.Println("No Modulo by 0")
			return
		}
		fmt.Println(a%b)
	} else if operator == "&" {
		fmt.Println(a&b)
	} else if operator == "|" {
		fmt.Println(a|b)
	} else if operator == "^" {
		fmt.Println(a^b)
	} else if operator == "&^" {
		fmt.Println(a&^b)
	} else {
		fmt.Println(0)
	}
}

func isNumeric(str string) bool {
	for _, r := range str {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func atoi(s string) int {
	var res int
	flag := false
	for i, r := range s {
		if r == 43 && i == 0 {
			continue
		}
		if r == 45 && i == 0 {
			flag = true
			continue
		}
		res *= 10
		if r < 48 || r > 57 {
			return 0
		}
		if r == 48 {
			res += 0
		} else if r == 49 {
			res += 1
		} else if r == 50 {
			res += 2
		} else if r == 51 {
			res += 3
		} else if r == 52 {
			res += 4
		} else if r == 53 {
			res += 5
		} else if r == 54 {
			res += 6
		} else if r == 55 {
			res += 7
		} else if r == 56 {
			res += 8
		} else if r == 57 {
			res += 9
		}
	}
	if flag {
		return -res
	}
	return res
}
