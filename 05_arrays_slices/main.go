package main

import "fmt"

func main() {
	// Arrays
	// var fruitArr [2]string
	// asign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"
	// declare array and assign at the same time
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)

	// Slices
	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
