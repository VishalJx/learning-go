package main

import "fmt"

func main() {

	fmt.Println(greet("Wick"))

	fmt.Println("-------callByValue.go--------")
}

func greet(name string) string {
	// return "Hello"
	return fmt.Sprintf("John, %s", name)
}
