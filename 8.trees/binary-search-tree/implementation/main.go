package main

import (
	"fmt"
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

// func showNode(node Node) NodeDeref {
// 	var nodeDeref NodeDeref
// 	if node.right != nil && node.left != nil {
// 		nodeDeref = NodeDeref{value: node.value, right: *node.right, left: *node.left}
// 	} else if node.right != nil && node.left == nil {
// 		nodeDeref = NodeDeref{value: node.value, right: *node.right}
// 	} else if node.left != nil && node.right == nil {
// 		nodeDeref = NodeDeref{value: node.value, left: *node.left}
// 	} else {
// 		nodeDeref = NodeDeref{value: node.value}
// 	}

// 	return nodeDeref

// }

func main() {
	myBinarySearchTree := newBinarySearchTree()
	myBinarySearchTree.insert(9)
	myBinarySearchTree.insert(4)
	myBinarySearchTree.insert(6)
	myBinarySearchTree.insert(20)
	myBinarySearchTree.insert(170)
	myBinarySearchTree.insert(15)
	myBinarySearchTree.insert(1)

	// BFS := myBinarySearchTree.breadthFirstSearch()
	// fmt.Println(BFS)
	// BFSR := myBinarySearchTree.breadthFirstSearchRecursive([]int{}, []*Node{myBinarySearchTree.root})
	// fmt.Println(BFSR)

	DFSInOrder := myBinarySearchTree.DFSInOrder()
	fmt.Println(DFSInOrder)

	// DFSPreOrder := myBinarySearchTree.DFSPreOrder()
	// fmt.Println(DFSPreOrder)

	// DFSPostOrder := myBinarySearchTree.DFSPostOrder()
	// fmt.Println(DFSPostOrder)
	// fmt.Println(*myBinarySearchTree)
	// fmt.Println(*myBinarySearchTree.root)
	// fmt.Println(*myBinarySearchTree.root.right)
	// fmt.Println(*myBinarySearchTree.root.left)
	// fmt.Println(*myBinarySearchTree.lookup(4).right)
	// fmt.Println(*myBinarySearchTree.lookup(4).left)
	// fmt.Println(myBinarySearchTree.lookup(20).right)
	// fmt.Println(myBinarySearchTree.lookup(20).left)
	// fmt.Println(myBinarySearchTree.lookup(6))
	// fmt.Println(myBinarySearchTree.lookup(170))
	// fmt.Println(myBinarySearchTree.lookup(15))
	// fmt.Println(myBinarySearchTree.lookup(1))

	// fmt.Println(preOrderTraverse(*myBinarySearchTree.root))

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

func (b *BinarySearchTree) remove(value int) {

	currentNode := b.root

	// for currentNode.right.value == value || currentNode.left.value == value {
	for currentNode != nil {
		if value > currentNode.value {
			if currentNode.right.value == value {
				if currentNode.right.right == nil && currentNode.right.left == nil {
					currentNode.right = nil
				} else if currentNode.right.left != nil && currentNode.right.right == nil {
					currentNode.right = currentNode.right.left
				} else if currentNode.right.right != nil && currentNode.right.left == nil {
					currentNode.right = currentNode.right.right
				} else {
					// TODO
				}
				return

			} else {
				currentNode = currentNode.right
			}
		} else if value < currentNode.value {
			if currentNode.left.value == value {
				if currentNode.left.right == nil && currentNode.left.left == nil {
					currentNode.left = nil
				} else if currentNode.left.left != nil && currentNode.left.right == nil {
					currentNode.left = currentNode.left.left
				} else if currentNode.left.right != nil && currentNode.left.left == nil {
					currentNode.left = currentNode.left.right
				} else {
					// TODO
				}
				return
			} else {
				currentNode = currentNode.left
			}

		}
	}
	return

}

func (b *BinarySearchTree) breadthFirstSearch() []int {
	currentNode := b.root
	list := []int{}
	queue := []*Node{}
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

func (b *BinarySearchTree) breadthFirstSearchRecursive(list []int, queue []*Node) []int {

	if len(queue) == 0 {
		return list
	}
	currentNode := queue[0]
	queue = queue[1:]
	list = append(list, currentNode.value)
	if currentNode.left != nil {
		queue = append(queue, currentNode.left)
	}
	if currentNode.right != nil {
		queue = append(queue, currentNode.right)
	}
	return b.breadthFirstSearchRecursive(list, queue)
}

// func (b *BinarySearchTree) DFSInOrder() []int {
// 	return traverseInOrder(b.root, []int{})
// }

// func traverseInOrder(node *Node, list []int) []int {
// 	log.Println(node.value)

// 	if node.left != nil {
// 		list = append(list, traverseInOrder(node.left, list)...)

// 	}
// 	list = append(list, node.value)
// 	if node.right != nil {
// 		list = append(list, traverseInOrder(node.right, list)...)

// 	}
// 	// list = append(list, node.value)
// 	return list
// }

func (b *BinarySearchTree) DFSInOrder() []int {
	return b.root.traverseInOrder()
}

func (node *Node) traverseInOrder() []int {

	var list []int
	if node.left != nil {
		fmt.Println(list)
		list = append(list, node.left.traverseInOrder()...)

	}
	list = append(list, node.value)

	if node.right != nil {
		list = append(list, node.right.traverseInOrder()...)

	}

	return list
}

func (b *BinarySearchTree) DFSPreOrder() []int {
	return b.root.traversePreOrder()
}

func (node *Node) traversePreOrder() []int {

	var list []int
	list = append(list, node.value)
	if node.left != nil {
		list = append(list, node.left.traversePreOrder()...)
	}

	if node.right != nil {
		list = append(list, node.right.traversePreOrder()...)
	}

	return list
}

func (b *BinarySearchTree) DFSPostOrder() []int {
	return b.root.traversePostOrder()
}

func (node *Node) traversePostOrder() []int {

	var list []int

	if node.left != nil {
		list = append(list, node.left.traversePostOrder()...)
	}

	if node.right != nil {
		list = append(list, node.right.traversePostOrder()...)
	}
	list = append(list, node.value)
	return list
}
