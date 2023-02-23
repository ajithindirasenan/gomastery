package main

import "fmt"

var ak int
var gk int

// The value of the variable declared is zero = 0,
// the value of the pointers, functions, interfaces, slices, channels, and maps are nil

func main() {

	ak = 20
	var h = 35
	fmt.Println(ak, h, gk)
	something()
}

func something() {
	fmt.Println(ak)
}
