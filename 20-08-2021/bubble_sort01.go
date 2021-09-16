package main

import "fmt"

func BubbleSort(array *[]int)  {
	for i := 0; i < len(*array)-1; i++ {
		for j := 0; j < len(*array)-i-1; j++ {
			if (*array)[j] > (*array)[j+1] {
				(*array)[j], (*array)[j+1] = (*array)[j+1], (*array)[j]
			}
		}
	}
}
func main() {
	array := []int{11, 14, 3, 8, 18, 17, 43}
	BubbleSort(&array)
	fmt.Println(array)

}
