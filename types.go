package main

import (
	"fmt"
	"math/cmplx"
)

const (
	PI   = 3.14
	LOVE = true
)

var (
	ToBe      bool       = false
	MaxInt    uint64     = 1<<64 - 1
	z         complex128 = cmplx.Sqrt(-5 + 12i)
	intVar    int        = 12
	intVar32  int32      = 123
	stringVar string     = "hello world"
)

func main() {
	const x = 10

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Print("Can we cast from string to int", int(intVar32))
}
