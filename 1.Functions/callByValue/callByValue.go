//Call by Value

package main

import "fmt"

func multiply(num1, num2 int) int {
	num1 = num1 + 1 //Mutation
	return num1 * num2
}

func main() {
	x := 5
	y := 6

	//Before
	fmt.Printf("Before: x=%d; y=%d\n", x, y)

	result := multiply(x, y)

	//Mutated
	fmt.Printf("Multiplication: %d\n", result)

	//After
	fmt.Printf("After: x=%d; y=%d\n", x, y)

}

// Conclusion is that the value remains the same