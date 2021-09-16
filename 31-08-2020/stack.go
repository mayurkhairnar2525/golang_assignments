package main

import "fmt"

type stack struct {
	item []int
}

//Push will add value at end
func (s *stack) push(i int) {
	s.item = append(s.item, i)
}

//Pop will remove value from the end
func (s *stack) pop() int {
	l := len(s.item) - 1
	toRemove := s.item[l]
	s.item = s.item[:l]
	return toRemove
}

func main() {
	mystack := stack{}
	fmt.Println(mystack)

	mystack.push(1)
	mystack.push(2)
	mystack.push(3)
	mystack.push(4)
	fmt.Println("Push:", mystack)

	mystack.pop()
	fmt.Println("Pop:", mystack)
}

