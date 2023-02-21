package main

func main() {
	var p *int
	i := 10
	p = &i
	println(*p)
	*p = 12
	println(i)
}
