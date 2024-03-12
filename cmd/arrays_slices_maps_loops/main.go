package main

import "fmt"

// This file explains about arrays

func main() {
	// ----------------- Arrays -----------------
	// Arrays are fixed size and cannot be resized
	var arr [3]int32 = [3]int32{1, 2, 3} // This is how you declare an array
	// OR arr := [3]int32{1, 2, 3}
	// This is how you declare an array without specifying the size - This will be inferred from the number of elements
	//  arr := [...]int32{1, 2, 3}
	fmt.Println(arr[:2]) // Indexing and slicing works similar to Python

	// ----------------- Slices -----------------
	// Slices are dynamic arrays
	// When you don't specify the size of the array, it automatically becomes a slice
	slice := []int32{1, 2, 3}
	newSlice := append(slice, 4)
	fmt.Println("newSlice", newSlice)
	// You can use the method `cap(newSlice)` to get the capacity of the slice
	// You can also append a slice to another slice
	slice1 := []int32{1, 2, 3}
	slice2 := []int32{4, 5, 6}
	slice1 = append(slice1, slice2...) // The spread operator is behind (opp to JS)
	fmt.Println("New slice1", slice1)
	// We can also make new slices using the `make(...) method`
	orangeSlices := make([]int32, 3, 5) // 3 is the length and 5 is the capacity
	fmt.Println("orangeSlices", orangeSlices)

	// ----------------- Maps -----------------
	// Maps are like objects in JS (key value pairs)
	userObj := make(map[string]uint8)
	fmt.Println("userObj", userObj)
	// OR declate the object directly
	ageObj := map[string]uint8{"Alpha": 21, "Beta": 35}
	fmt.Println("ageObj", ageObj)
	// You can access the map using the key -- ageObj["Alpha"]
	// If you try accessing a key that is not present, it will return the default value of the type
	// -- In this case - uint8 will return a default value of 0
	var age, ok = ageObj["Alpha"]     // ok will be true if the key is present, and false if otherwise
	fmt.Println("age", age, "ok", ok) // age 21 ok true
	// You can also remove a key from the map using the `delete` method
	delete(ageObj, "Alpha")

	// ----------------- Loops -----------------
}
