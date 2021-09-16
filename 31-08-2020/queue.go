package main

import "fmt"

type queue struct {
	item []int
}

func (q *queue) enqueue(i int) {
	q.item = append(q.item, i)
}

func (q *queue) dequeue() int {
	toRemove := q.item[0]
	q.item = q.item[1:]
	return toRemove
}
func main() {
	myqueue := queue{}
	fmt.Println(myqueue)

	myqueue.enqueue(1)
	myqueue.enqueue(2)
	myqueue.enqueue(3)
	myqueue.enqueue(4)
	fmt.Println("Enqueue:", myqueue)

	myqueue.dequeue()
	fmt.Println("dequeue:", myqueue)
}
