package main

import "fmt"

type Stack struct {
	top    *Node
	bottom *Node
	length int
}

type Node struct {
	value interface{}
	next  *Node
}

func newStack() *Stack {
	return &Stack{}
}

func createNode(value interface{}) *Node {
	return &Node{value: value}
}

func main() {

	myStack := newStack()
	myStack.push("google")
	myStack.push("twitter")
	myStack.push("udemy")
	myStack.push("stackoverflow")

	myStack.print()

	fmt.Println(myStack.peek())

	myStack.pop()

	myStack.print()
}

func (s *Stack) print() {
	stackSlice := []Node{}
	currentNode := s.top
	for currentNode != nil {
		stackSlice = append(stackSlice, *currentNode)
		currentNode = currentNode.next
	}
	fmt.Println(s)
	fmt.Println(stackSlice)

}

func (s *Stack) isEmpty() bool {
	return s.length == 0
}

// pop removes the last item
func (s *Stack) pop() {
	if s.isEmpty() {
		fmt.Println("The stack is empty")
		return
	}
	if s.top == s.bottom {
		s.bottom = nil
	}
	newTop := s.top.next
	s.top.next = nil
	s.top = newTop
	// s.top = s.top.next

	s.length--
}

// push adds item to the top
func (s *Stack) push(value interface{}) {
	newNode := createNode(value)
	if s.isEmpty() {
		s.top = newNode
		s.bottom = newNode
	} else {
		newNode.next = s.top
		s.top = newNode
	}
	s.length++
}

// peek shows the top item
func (s *Stack) peek() Node {
	return *s.top
}
