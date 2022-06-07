package main

import "fmt"

type LinkedList struct {
	head *Node
	tail *Node
}

func newLinkedList() *LinkedList {
	return &LinkedList{}
}

type Node struct {
	value    interface{}
	previous *Node
	next     *Node
}

func createNode(value interface{}) *Node {
	return &Node{value: value}
}

func main() {
	myLinkedList := newLinkedList()
	myLinkedList.append(100)
	myLinkedList.append(200)
	myLinkedList.append(300)

	myLinkedList.prepend("a")
	myLinkedList.prepend("b")
	myLinkedList.prepend("c")

	myLinkedList.insert("insertion element", 4)
	myLinkedList.print()

	myLinkedList.remove(2)
	myLinkedList.print()
	myLinkedList.remove(6)
	myLinkedList.print()
}

func (l *LinkedList) print() {
	slice := []Node{}
	currentNode := l.head
	for currentNode != nil {
		slice = append(slice, *currentNode)
		currentNode = currentNode.next
	}
	fmt.Println(slice)
}
func (l *LinkedList) insert(value interface{}, index int) {
	newNode := createNode(value)
	currentNode := l.head
	for i := 1; i < index; i++ {
		if currentNode.next != nil {
			currentNode = currentNode.next
		} else {
			l.append(currentNode)
			fmt.Println("The index is out of range")
			return
		}
	}
	currentNode.next.previous = newNode
	newNode.next = currentNode.next
	currentNode.next = newNode
	newNode.previous = currentNode

}

func (l *LinkedList) append(value interface{}) {
	newNode := createNode(value)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.previous = l.tail
		l.tail.next = newNode
		l.tail = newNode
	}
}

func (l *LinkedList) prepend(value interface{}) {
	newNode := createNode(value)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.head.previous = newNode
		newNode.next = l.head
		l.head = newNode
	}

}

func (l *LinkedList) remove(index int) {
	currentNode := l.head
	for i := 1; i < index; i++ {
		if currentNode.next != nil {
			currentNode = currentNode.next
		} else {
			fmt.Println("Index out of range")
			return
		}
	}

	currentNode.next.next.previous = currentNode
	currentNode.next = currentNode.next.next
}
