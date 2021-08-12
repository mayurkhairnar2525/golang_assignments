package main

import (
	"fmt"
	"os"
)

func main() {
	// Write a program in golang to print even numbers within a range
	var min, max int

	fmt.Print("Enter the Minimum number:")
	fmt.Scanln(&min)

	fmt.Print("Enter the Maximum number:")

	fmt.Scanln(&max)
	if (max <min){
		fmt.Println("Invalid range given from user side")
		os.Exit(1)
	}

	if min%2 != 0 {
		min++
	}
	for i := min; i <= max; i = i+2 {
		fmt.Println(i)
	}
}
