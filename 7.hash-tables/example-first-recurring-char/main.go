package main

import "fmt"

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}
type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	testHashTable := Init()
	testHashTable.Insert("Randy")
	fmt.Println(testHashTable)
}

func (b *bucket) insert(key string) {
	if !b.search(key) {
		newNode := &bucketNode{
			key: key,
		}
		newNode.next = b.head
		b.head = newNode

	} else {
		fmt.Println("Key already exists")
	}

}
func (b *bucket) search(key string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *bucket) delete(key string) {

	if b.head.key == key {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == key {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
	return
}

// Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)

}

func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)

}

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)

}

func hash(key string) int {
	total := 0
	for i := range key {
		total += int(rune(key[i]))
	}
	return total % ArraySize
}
