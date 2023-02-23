package main

import "fmt"

type ak int

var a ak

var b int

func main() {

	a = 30
	b = 15
	fmt.Println(a)
	fmt.Println(b)

	b = int(a)
	fmt.Println(b)
}
