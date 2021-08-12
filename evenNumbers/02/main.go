package main

import "fmt"

func main() {

	var evnum, i int

	fmt.Print("Enter the Number: ")
	fmt.Scanln(&evnum)

	for i = 0;i<=evnum; i++ {
		if i%2 == 0{
			fmt.Println("Number's:",i)
			i++
		}
	}
}
