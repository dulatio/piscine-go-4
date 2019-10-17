package main

import (
	"fmt"
	piscine ".."
)

func main(){
	fmt.Println(piscine.Index("Hello!", "l"))
	fmt.Println(piscine.Index("汉字",  "字"))
	fmt.Println(piscine.Index("á, é, í, ó, ú, ü, ñ", "!!!"))
}
