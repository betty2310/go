package main

type vector struct {
	x int
	y int
}

func (v vector) double(y int) int {
	return v.x + y
}

type myInt int

func (x myInt) add(y int) int {
	return int(x) + y
}

func main() {
	v := vector{3, 4}
	println(v.double(12))
	var x myInt
	x = 12
	println(x.add(1))
}
