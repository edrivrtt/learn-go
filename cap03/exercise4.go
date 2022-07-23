package main

import "fmt"

type whatever int

var x whatever

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x := 42
	fmt.Println(x)

}
