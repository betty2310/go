package main

import "fmt"

func buzz(n int) bool {
	if n%3 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Print(buzz(12))
}
