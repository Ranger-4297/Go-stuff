package main

import "fmt"

func main() {
	// Fixed length array
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Arrary is", a)

	// Any length array
	b := []int{2, 4, 6, 1, 9, 1}
	fmt.Println("Slice is", b)

	// Appended array
	c := append(b, 13)
	fmt.Println("Appended slice is", c)

	// Maps
	vertices := make(map[string]int)
	vertices["triangles"] = 3
	vertices["Squares"] = 4
	fmt.Println(vertices) 
}
