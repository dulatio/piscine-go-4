package main

import (
	"fmt"
	piscine ".."
)

func main(){
	str := "Hello!"
	str1 := "W"
	str2 := "Héllo!"
	fmt.Println(piscine.StrLen(str))
	fmt.Println(piscine.StrLen(str1))
	fmt.Println(piscine.StrLen(str2))
}
