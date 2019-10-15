package piscine

import "fmt"

func PrintComb2(){
	for i1 := 0; i1 <= 9; i1++ {
		for i2 := 0; i2 <= 9; i2++ {
			for j1 := i1; j1 <= 9; j1++{
				for j2 := 0; j2 <= 9; j2++{
					if(i1 == j1 && i2 >= j2){
					} else {
						fmt.Printf("%d%d %d%d", i1, i2, j1, j2)
						if(i1 != 9 || i2 != 8 || j1 != 9 || j2 != 9){
							fmt.Printf(", ")
						}
					}
				}
			}
		}
	}
	fmt.Println()
}
