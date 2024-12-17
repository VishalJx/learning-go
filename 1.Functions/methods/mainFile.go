//This file serves as entry point for other two files
//i.e. pointerReceiverMethod.go and valueReceiverMethod.go

package main

import "fmt"

func main() {
	// Passing by value
	radius, area := ValueReceiverMethodExample()
	fmt.Println("Radius: ", radius)
	fmt.Println("Area: ", area)

	fmt.Println("---------------------------------")

	//Passing by reference
	ogLength, ogBreadth, modLength, modBreadth, rArea := PointerReceiverMethodExample()

	fmt.Printf("Original length: %f & breadth: %f \n", ogLength, ogBreadth)
	fmt.Printf("Modified length: %f & breadth: %f \n", modLength, modBreadth)
	fmt.Println("Area of the rectangle:", rArea)
}

/* SUMMARY
- Value Receiver: Use it when you don’t need to modify the object.
- Pointer Receiver: Use it when you need to modify the object or avoid copying large objects.
*/
