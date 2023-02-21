package main

import "fmt"

func function_value(fn func(x, y float64) float64, x int) int {
	return int(fn(1, 2)) * x
}

func main() {
	fun := func(x, y float64) float64 {
		return x*2 + y*3
	}
	fmt.Printf("%T\n", fun)
	println(int(fun(12, 23)), fun(3, 4))
	println(function_value(fun, 12))
}
