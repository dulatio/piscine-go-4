package main

import (
	"fmt"
	piscine ".."
)

func main(){
	s := []int{5,4,3,2,1,0}
	piscine.ASortIntegerTable(s)
	fmt.Println(s)
	s1 := []int{-12,-15,4,0,2,1,0}
	piscine.ASortIntegerTable(s1)
	fmt.Println(s1)
	s2 := []int{}
	piscine.ASortIntegerTable(s2)
	fmt.Println(s2)
	s3 := []int{1}
	piscine.ASortIntegerTable(s3)
	fmt.Println(s3)
	s4 := []int{0,1,2,3,4,5,6,7,8,9,10}
	piscine.ASortIntegerTable(s4)
	fmt.Println(s4)
}
