package main

import "fmt"

func main() {
	var min, max int

	fmt.Print("Enter the Minimum number:")
	fmt.Scanln(&min)

	fmt.Print("Enter the Maximum number:")

	fmt.Scanln(&max)

	if min%2 != 0 {
		min++
	}
	for i := min; i <= max; i = i+2 {
		fmt.Println(i)
	}
}
