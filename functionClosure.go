package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		c := a + b
		a = b
		b = c
		return a
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 4; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
	fibonacci_number := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci_number())
	}

}
