package main

import "fmt"

func main() {
	var x1, y1, x2, y2, x, y int
	fmt.Scanf("%d%d%d%d%d%d", &x1, &y1, &x2, &y2, &x, &y)
	res := (x-x1)*(y2-y1) - (y-y1)*(x2-x1)
	if res == 0 {
		fmt.Println("On line")
	}
	if res > 0 {
		fmt.Println("On right")
	}
	if res < 0 {
		fmt.Println("On left")
	}
}
