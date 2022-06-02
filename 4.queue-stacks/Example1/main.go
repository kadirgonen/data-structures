package main

import "fmt"

func main() {
	myQueue := Constructor()
	fmt.Println(myQueue.stack1)
	myQueue.Push(1)
	fmt.Println(myQueue.stack1)
	fmt.Println(myQueue.Pop())
	fmt.Println(myQueue.Empty())
}

type MyQueue struct {
	length int
	stack1 Stack
	stack2 Stack
}

func Constructor() MyQueue {
	return MyQueue{stack1: *newStack(), stack2: *newStack()}
}

func (this *MyQueue) Push(x int) {

	if this.Empty() {
		this.stack1.push(x)
	} else {
		for !this.stack1.isEmpty() {
			this.stack2.push(this.stack1.peek())
			this.stack1.pop()
		}
		this.stack1.push(x)
		for !this.stack2.isEmpty() {
			this.stack1.push(this.stack2.peek())
			this.stack2.pop()
		}
	}
	this.length = this.stack1.size()
}

func (this *MyQueue) Pop() int {
	result := this.stack1.peek()
	this.stack1.pop()
	this.length = this.stack1.size()
	return result
}

func (this *MyQueue) Peek() int {
	return this.stack1.peek()
}

func (this *MyQueue) Empty() bool {
	return this.length == 0
}

type Stack struct {
	data []int
}

func newStack() *Stack {
	return &Stack{}
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) push(value int) {
	s.data = append(s.data, value)
}

func (s *Stack) pop() {
	if s.isEmpty() {
		fmt.Println("Stack is empty")
		return
	}
	s.data = s.data[:len(s.data)-1]
}

func (s *Stack) peek() int {
	if s.isEmpty() {
		fmt.Println("Stack is empty")
		return -1
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) size() int {
	return len(s.data)
}
