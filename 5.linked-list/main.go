package main

import "fmt"

type LinkedList struct {
	head *Node
	tail *Node
}
type Node struct {
	value interface{}
	next  *Node
}

func newLinkedList() *LinkedList {
	return &LinkedList{}
}
func main() {
	myLinkedList := newLinkedList()
	myLinkedList.append(10)
	myLinkedList.append(20)
	myLinkedList.append(30)

	// myLinkedList.prepend("x")
	// myLinkedList.prepend("y")
	// myLinkedList.prepend("z")

	myLinkedList.insert("insertion element", 7)
	myLinkedList.print()

	myLinkedList.reverse()
	myLinkedList.print()
	// myLinkedList.remove(8)
	// fmt.Println(*myLinkedList.head)
	// fmt.Println(*myLinkedList.tail)
	// myLinkedList.print()

}

func createNode(value interface{}) *Node {
	return &Node{value: value}
}

// print to see all the elements to check
func (l *LinkedList) print() {
	listSlice := []interface{}{}
	currentNode := l.head
	for currentNode != nil {
		listSlice = append(listSlice, *currentNode)
		currentNode = currentNode.next
	}
	fmt.Println(listSlice)
}

func (l *LinkedList) insert(value interface{}, index int) {

	newNode := createNode(value)

	currentNode := l.head
	for i := 1; i < index; i++ {
		if currentNode.next != nil {
			currentNode = currentNode.next
		} else {
			fmt.Println("List length is below given index")
			l.append(value)
			return
		}

	}
	newNode.next = currentNode.next
	currentNode.next = newNode
}

func (l *LinkedList) append(value interface{}) {

	newNode := createNode(value)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
}

func (l *LinkedList) prepend(value interface{}) {

	newNode := createNode(value)

	if l.head == nil {
		l.head = newNode
	} else {
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
			fmt.Println("List length is below given index")
			return
		}
	}
	currentNode.next = currentNode.next.next
}

func (l *LinkedList) reverse() {

	if l.head.next == nil {
		return
	}

	currentNode := l.head
	l.tail = l.head
	nextNode := currentNode.next

	for nextNode != nil {
		temp := nextNode.next
		nextNode.next = currentNode
		currentNode = nextNode
		nextNode = temp

	}
	l.head.next = nil
	l.head = currentNode

}

// func (l *LinkedList) reverse() {
// 	list := []Node{}
// 	currentNode := l.head
// 	for currentNode != nil {
// 		list = append(list, *currentNode)
// 		currentNode = currentNode.next
// 	}
// 	// fmt.Println(list)
// 	for i := len(list) - 1; i > 0; i-- {
// 		// ptr := &(list[i-1])
// 		fmt.Printf("%p \n", list[i].next)
// 		list[i].next = &list[i-1]
// 		fmt.Printf("%p \n", list[i].next)
// 	}
// 	list[0].next = nil
// 	l.head = &list[len(list)-1]
// 	l.tail = &list[0]
// }
