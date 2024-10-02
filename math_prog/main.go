package main

import (
	"fmt"
	"main/fibonachi"
)

func main() {
	var val fibonachi.Values
	val.New(0, 4, 40)
	fmt.Printf("result=%f", val.Get_result())
}
