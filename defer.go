package main

import "fmt"

func main() {
	fmt.Println("0")
	fmt.Println("1")
	defer fmt.Println("defer")
	fmt.Println("2")
	fmt.Println("3")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

}
