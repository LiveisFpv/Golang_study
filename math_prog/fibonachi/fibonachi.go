package fibonachi

import (
	"fmt"
	"math"
)

type Values struct {
	a   float64
	b   float64
	x1  float64
	x2  float64
	fx1 float64
	fx2 float64
	n   int
}

func (val *Values) New(a, b float64, n int) {
	val.a = a
	val.b = b
	val.n = n
}

func get_fibonachi(n int) float64 {
	var N float64
	N = float64(n)
	var res float64
	res = (math.Pow((1.0+math.Sqrt(5))/2, N) - math.Pow((1.0-math.Sqrt(5))/2, N)) / math.Sqrt(5)
	return res
}
func get_func_value(x float64) float64 {
	var res float64
	res = -math.Exp(-x) * math.Log(x)
	return res
}
func (val Values) Get_result() float64 {
	val.x1 = val.a + get_fibonachi(val.n)/get_fibonachi(val.n+2)*(val.b-val.a)
	val.x2 = val.a + get_fibonachi(val.n+1)/get_fibonachi(val.n+2)*(val.b-val.a)
	val.fx1 = get_func_value(val.x1)
	val.fx2 = get_func_value(val.x2)
	val.n -= 1
	fmt.Printf("Initial values: a:%f b:%f x1:%f x2:%f\n", val.a, val.b, val.x1, val.x2)
	for ; val.n > 0; val.n-- {
		if val.fx1 < val.fx2 {
			fmt.Println("f(x1)<f(x2) ---- ", val.fx1, "<", val.fx2)
			val.b = val.x2
			val.x2 = val.x1
			val.x1 = val.a + val.b - val.x2
			val.fx2 = val.fx1
			val.fx1 = get_func_value(val.x1)
			fmt.Printf("New values: a: %f b: %f x1: %f x2: %f\n", val.a, val.b, val.x1, val.x2)
		} else {
			fmt.Println("f(x1)>f(x2) ---- ", val.fx1, ">", val.fx2)
			val.a = val.x1
			val.x1 = val.x2
			val.x2 = val.a + val.b - val.x1
			val.fx1 = val.fx2
			val.fx2 = get_func_value(val.x2)
			fmt.Printf("New values: a: %f b: %f x1: %f x2: %f\n", val.a, val.b, val.x1, val.x2)
		}
	}
	return val.x1
}
