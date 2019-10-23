package main

import (
	"fmt"
	piscine ".."
)

func main() {
	result := piscine.ConvertBase("9223372036854775807", "0123456789", "01")
	fmt.Println(result)
}
