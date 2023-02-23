package main

import (
	"fmt"
)

/* Go is a statically typed programming language not a dinamically typed language
because we cannot assign an type to an already existing type of an variable.
*/

var ajj string

var gkk string = "the gkk is something else\n"

var rawliteral = `the string, \n "ajith i\t ""^s a good senshey"  ^.^`

func main() {

	ajj = "hey da"
	fmt.Println(ajj)
	fmt.Printf("%T\n", ajj)
	fmt.Printf(gkk)
	fmt.Printf("%T\n", gkk)
	fmt.Printf(rawliteral)
	fmt.Printf("\n%T", rawliteral)
}
