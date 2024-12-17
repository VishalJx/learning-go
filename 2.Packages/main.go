package main

import (
	"fmt"
	x "go-a-to-z/2.Packages/extra_folder"
)

func main() {

	var num3 = AddNumbers(5, 3)
	var num4 = x.AddNumbersExternal(6, 4)

	//both in root folder
	fmt.Printf("Imported sum from add.go is: %d\n", num3)

	//both in separate folders
	fmt.Printf("Imported sum from insidefolder_add.go is: %d", num4)

}
