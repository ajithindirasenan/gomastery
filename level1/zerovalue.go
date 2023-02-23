package main

import (
	"fmt"
)

var z int
var x string
var c bool
var v float64

func main() {
	fmt.Println("print the value of z", z)
	fmt.Printf("%T\n", z)
	fmt.Println("print the value of x", x)
	fmt.Printf("%T\n", x)
	fmt.Println("print the value of c", c)
	fmt.Printf("%T\n", c)
	fmt.Println("print the value of z", v)
	fmt.Printf("%T\n", v)
}
