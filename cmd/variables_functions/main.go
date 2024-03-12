// Modules are combination of packages and their dependencies
// Packages are used to organize code into different files and directories

package main // -- Specify the package name - `main` is the default package name for the main file`
// -- This file will be executed first when the program is run

// Cannot have any unused imports and variables in GO

// The fmt package is used for formatting and printing
import (
	"errors"
	"fmt"
)

// This file explains about basic variables and types and also functions

func main() {
	printValue := "Hello, World!"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 0
	// It's a good practice to handle errors
	// -- Errors handling like this are usually a general design pattern in Go
	result, remainder, err := intDivision(numerator, denominator)
	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("Result: %d", result)
	} else {
		fmt.Printf("Result: %d, Remainder: %d", result, remainder)
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}

// DEFAULT VALUES
// Go will specify default values for each type if you don't specify a value
// For example, the default value for an int is 0
// The default value for a float is 0.0
// The default value for a string is an empty string
// The default value for a boolean is false
// The default value for a pointer is nil
// The default value for a function is nil

func Numbers() {
	// There are different types of integers in Go like int8, int16, int32, int64
	var intNum int = 32767
	fmt.Println(intNum)

	// There are different types of floating point numbers in Go like float32, float64
	var floatNum float64 = 3.14
	fmt.Println(floatNum)

	// You cannot do operations between different types of numbers
	// fmt.Println(intNum + floatNum) // Uncomment this and it will throw an error
	// You need to convert the numbers to the same type
	fmt.Println(intNum + int(floatNum)) // This will work
}

func Variables() {
	var myString string = "Hello, World!"
	fmt.Println(myString)

	// OR you can use the shorthand
	myString2 := "Hello, World!"
	fmt.Println(myString2)

	// You can also declare multiple variables at once
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)
}
