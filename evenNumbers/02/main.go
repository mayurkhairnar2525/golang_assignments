package main

import "fmt"

func main() {
	var number int = 0
	var evnum int
	fmt.Print("Enter the Number: ")
	fmt.Scanln(&evnum)
	for i :=evnum; number<10; i++ {
		if i%2 == 0 {
			fmt.Println("Number's:", i)
			number ++
		}
	}
	fmt.Println("You are out of the scope")
}
