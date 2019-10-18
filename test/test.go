package main

import (
	"fmt"
	raid1 "../raid1"
)

func main(){
	raid1.Raid1c(5,3)
	raid1.Raid1c(5,1)
	raid1.Raid1c(1,5)
	raid1.Raid1c(0,3)
	raid1.Raid1c(1,1)
	raid1.Raid1c(2,2)
	fmt.Println()
}
