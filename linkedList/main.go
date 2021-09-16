package main

import "fmt"

type node struct {
	next *node
	data int
}

type linkedList struct {
	head   *node
	lenght int
}

func (l *linkedList) add(value int) {
	newNode := node{
		data: value,
	}
	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
		l.lenght++
	} else {
		l.head = &newNode
		l.lenght++
	}
	return
}

func (l *linkedList) printList() {
	if l.head == nil {
		return
	}
	currentNode := l.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

func (l *linkedList) deleteWithValue(value int) {
	if l.lenght == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.lenght--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.lenght--
}


func main() {
	myList := linkedList{}
	myList.add(21)
	myList.add(78)
	myList.add(64)
	myList.add(57)
	myList.printList()
	fmt.Println("Length:", myList.lenght)
	fmt.Println("After deleting:")
	myList.deleteWithValue(78)
	myList.printList()
	fmt.Println("Length:", myList.lenght)


}
