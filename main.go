package main

import (
	"fmt"
	"practice-mod/helper"     // Importing the helper package
	"practice-mod/helpermain" // Importing the helper package
)

func main() {
	message := helper.Greet() // Using the Greet() function from the helper package
	fmt.Println(message)
	helpermain.HelperMainPackage()
}
