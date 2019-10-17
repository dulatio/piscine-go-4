package main

import (
	"fmt"
	piscine ".."
)

func main(){
	str := "1"
	str1 := "0000012"
	str2 := "6431"
	fmt.Println(piscine.BasicAtoi(str))
	fmt.Println(piscine.BasicAtoi(str1))
	fmt.Println(piscine.BasicAtoi(str2))
}
