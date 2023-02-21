package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

var n = map[int]string{
	1: "hello world",
	2: "func go",
}

func main() {
	fmt.Println(m["Bell Labs"])
	fmt.Println(n[1])
}
