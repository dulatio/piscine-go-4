package main

import (
	"fmt"
	raid1 "../raid1"
)

func main(){
	raid1.Raid1e(5,3)
	raid1.Raid1e(5,1)
	raid1.Raid1e(1,5)
	raid1.Raid1e(0,3)
	raid1.Raid1e(1,1)
	raid1.Raid1e(2,2)
	fmt.Println()
}
