package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	/*In this for loop, first argument represents index(key)
	and second represents value*/

	//Using Value of elements
	// for _, n := range nums {
	// 	total = total + n
	// }

	//Using Index of elements
	for i, _ := range nums {
		total = total + nums[i]
	}

	//Traditional 'for' loop
	// for i := 0; i < len(nums); i++ {
	// 	total = total + nums[i]
	// }

	return total
}

func main() {
	fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))
	fmt.Println("Sum of no numbers:", sum())
}

//Basically we pass array of arguments
//Variadic functions can only have one variadic parameter,
//and it must be the last parameter
