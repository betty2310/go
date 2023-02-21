package main

func add(x, y int) int {
	return x + y
}

func swap(x, y string, a ...any) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum/2 + 1
	y = sum%2 + 1
	return
}

func main() {
	var x, y, z, m = true, false, false, "hello"
	println(x, y, z, m)
	str := "implicit call function!"
	println(str)
}
