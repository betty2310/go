package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	s[2] = 12

	// var nilSlice []int
	// if nilSlice == nil {
	// 	fmt.Println("nil", len(nilSlice), cap(nilSlice))
	// }
	// nilSlice = make([]int, 5, 5)
	// fmt.Println(len(nilSlice), cap(nilSlice))

	for i, v := range s {
		fmt.Println(i, v*2)
	}
	//skip index or value using _
	for _, v := range s {
		fmt.Println(v)
	}
}
