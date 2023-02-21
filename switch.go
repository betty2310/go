package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("Happy try MacOS from Thinkpad!")
	default:
		fmt.Println("F**K!")
	}
}
