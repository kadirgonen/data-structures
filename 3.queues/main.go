package main

import "fmt"

type Queue struct {
	first  *Node
	last   *Node
	length int
}

type Node struct {
	value interface{}
	next  *Node
}

func newQueue() *Queue {
	return &Queue{}
}
func createNode(value interface{}) *Node {
	return &Node{value: value}
}
func main() {

	myQueue := newQueue()
	myQueue.enqueue("Joy")
	myQueue.enqueue("Matt")
	myQueue.enqueue("Pavel")
	myQueue.enqueue("Samir")

	myQueue.print()

	fmt.Println(myQueue.peek())
	myQueue.dequeue()
	myQueue.dequeue()
	myQueue.dequeue()
	myQueue.dequeue()

	myQueue.print()

}

func (q *Queue) print() {
	queueSlice := []Node{}
	currentNode := q.first
	for currentNode != nil {
		queueSlice = append(queueSlice, *currentNode)
		currentNode = currentNode.next
	}
	fmt.Println(q)
	fmt.Println(queueSlice)
}
func (q *Queue) isEmpty() bool {
	return q.length == 0
}
func (q *Queue) enqueue(value interface{}) {
	newNode := createNode(value)
	if q.isEmpty() {
		q.first = newNode
		q.last = newNode
	} else {
		q.last.next = newNode
		q.last = newNode
	}
	q.length++

}

func (q *Queue) dequeue() {
	if q.isEmpty() {
		fmt.Println("The queue is currently empty")
		return
	}

	q.first = q.first.next

	q.length--
}

func (q *Queue) peek() *Node {
	return q.first
}
