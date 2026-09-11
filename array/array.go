package main

/*
This program demonstrates various operations on arrays in Go.
It includes functions to print arrays, sum elements, generate Fibonacci numbers,
generate random arrays, and sort arrays using the bubble sort algorithm.

Note the difference between := and var in Go. The := syntax is used for short variable declaration and initialization, while var is used for declaring variables with an optional initial value.
*/
import (
	"fmt"
	"math/rand"
)


func printArray5(array [5]int) {
	for _, value := range array {
		fmt.Println("verdi", value)
	}

}

// Prints a heading with the given text
func heading(text string) {
	fmt.Println("********", text, "********")
}

// Sums the elements of an array of 5 integers and returns the result
func sumArray(array [5]int) int {
	sum := 0 // initialize sum to 0
	for _, value := range array {
		sum += value
	}
	return sum
}

// Generates the first 100 Fibonacci numbers and returns them as an array
func fabonacci() [100]int {
	var array [100]int
	array[0], array[1] = 0, 1
	for i := 2; i < 100; i++ {
		array[i] = array[i-1] + array[i-2]
	}
	return array
}

// Prints the elements of an array of 100 integers
func printArray100(fab [100]int ) {
	for _, value := range fab {
		fmt.Println("verdi", value)
	}
}

// Generates an array of 100 integers with random values between 0 and 99
func generateArray() [100]int {
	var array [100]int // not initialized with specific values, will be filled with random numbers
	for i := 0; i < 100; i++ {
		array[i] = rand.Intn(100)
	}
	return array
}

// Sorts an array of 100 integers using the bubble sort algorithm and returns the sorted array
func bubbleSort(array [100]int) [100]int {
	for i := 0; i < 100-1; i++ {
		for j := 0; j < 100-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func main(){
	var array [5]int
	array[4] = 13
	
	heading("Array Elements")
	printArray5(array)

	heading("Sum of Array Elements")	
	fmt.Println(sumArray(array))

	heading("Fibonacci Sequence")
	printArray100(fabonacci())

	heading("sort array")
	randomArray := generateArray()
	printArray100(randomArray)
	sortedArray := bubbleSort(randomArray)
	printArray100(sortedArray)

	heading("Multi dimensions arrays")
	var multiArray [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			multiArray[i][j] = i * j
			fmt.Printf("multiArray[%d][%d] = %d\n", i, j, multiArray[i][j])
		}
	}

}