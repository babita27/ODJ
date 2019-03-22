package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	p := point{2, 4}
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", p.x)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", true)
}
