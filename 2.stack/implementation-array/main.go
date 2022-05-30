package main

import "fmt"

type Stack struct {
	top    interface{}
	bottom interface{}
	length int
}

func newStack() *Stack {
	return &Stack{}
}

var stackArr []interface{}

func main() {

	myStack := newStack()
	myStack.push("google")
	myStack.push("twitter")
	myStack.push("udemy")
	myStack.push("discord")

	myStack.print()

	myStack.pop()
	myStack.print()
}

func (s *Stack) isEmpty() bool {
	return s.length == 0
}
func (s *Stack) print() {
	fmt.Println(s)
	fmt.Println(stackArr)
}

func (s *Stack) push(value interface{}) {
	// if s.isEmpty() {
	// 	s.top=value
	// 	s.bottom=value
	// }
	stackArr = append(stackArr, value)
	s.top = stackArr[len(stackArr)-1]
	s.bottom = stackArr[0]
	s.length = len(stackArr)
}
func (s *Stack) pop() {

	if s.isEmpty() {
		fmt.Println("Stack is currently empty")
		return
	}
	stackArr = stackArr[:len(stackArr)-1]
	s.top = stackArr[len(stackArr)-1]
	s.length = len(stackArr)
}

func (s *Stack) peek() interface{} {
	return stackArr[len(stackArr)-1]
}
