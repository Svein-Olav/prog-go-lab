package main

import "fmt"

func printArray(array [5]int) {
	for _, value := range array {
		fmt.Println("verdi", value)
	}

}

func main(){
	var array [5]int
	array[4] = 13
	printArray(array)
}