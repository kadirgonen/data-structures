package main

import (
	"fmt"
	"log"
)

type BinarySearchTree struct {
	root *Node
}
type Node struct {
	value int
	right *Node
	left  *Node
}

type NodeDeref struct {
	value int
	right Node
	left  Node
}

func newBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}
func createNode(value int) *Node {
	return &Node{value: value}
}

func main() {
	myBinarySearchTree := newBinarySearchTree()
	myBinarySearchTree.insert(9)
	myBinarySearchTree.insert(4)
	myBinarySearchTree.insert(6)
	myBinarySearchTree.insert(20)
	myBinarySearchTree.insert(170)
	myBinarySearchTree.insert(15)
	myBinarySearchTree.insert(1)
	// fmt.Println(myBinarySearchTree.breadthFirstSearch())
	fmt.Println(myBinarySearchTree.validate())
}

func (b *BinarySearchTree) insert(value int) {
	newNode := createNode(value)
	if b.root == nil {
		b.root = newNode
		return
	}

	currentNode := b.root
	for {
		if newNode.value > currentNode.value {
			if currentNode.right == nil {
				currentNode.right = newNode
				break
			}
			currentNode = currentNode.right

		} else {
			if currentNode.left == nil {
				currentNode.left = newNode
				break
			}
			currentNode = currentNode.left

		}
	}

}

func (b *BinarySearchTree) lookup(value int) *Node {
	currentNode := b.root
	for currentNode != nil {
		if value > currentNode.value {
			currentNode = currentNode.right
		} else if value < currentNode.value {
			currentNode = currentNode.left
		} else {
			return currentNode
		}
	}
	return nil
}

func (b *BinarySearchTree) breadthFirstSearch() []int {
	currentNode := b.root
	var queue []*Node
	var list []int
	queue = append(queue, currentNode)

	for len(queue) > 0 {
		currentNode = queue[0]
		queue = queue[1:]
		list = append(list, currentNode.value)
		if currentNode.left != nil {
			queue = append(queue, currentNode.left)
		}
		if currentNode.right != nil {
			queue = append(queue, currentNode.right)
		}

	}
	return list
}

func (b *BinarySearchTree) validate() bool {
	list := b.breadthFirstSearch()

	for i := range list {
		if i*2+2 < len(list) {
			if list[i] < list[i*2+1] || list[i] > list[i*2+2] {
				log.Println(list[i])
				return false
			}
		}
	}
	return true
}
