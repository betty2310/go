package main

import "fmt"

type stack interface {
	push(value int)
	pop()
	peek() int
	count() int
	isEmpty() bool
}

type Int []int

func (x *Int) push(value int) {
	var t []int = *x
	t = append(t, value)
	*x = t
}

func (x *Int) pop() {
	var t []int = *x
	t = t[0 : len(t)-1]
	*x = t
}

func (x *Int) peek() int {
	return (*x)[x.count()-1]
}

func (x *Int) count() int {
	return len(*x)
}

func (x *Int) isEmpty() bool {
	if x.count() == 0 {
		return false
	}
	return true
}

func main() {
	var value = Int{1, 2, 3, 5, 12}
	var s stack = &value
	// s.push(2)
	fmt.Println(s.count(), s.peek())
	s.pop()
	fmt.Println(s.count(), s.peek())
	s.push(23)
	fmt.Println(s.peek())
}
