package main

import (
	"fmt"
	"log"
)

const ArraySize = 7

type HashTable struct {
	data [ArraySize][][]interface{}
}

func main() {
	testHashTable := &HashTable{}
	testHashTable.Insert(1934, "Uruguay")
	testHashTable.Insert(1938, "Brazil")
	testHashTable.Insert(1942, "Argentina")
	testHashTable.Insert(1946, "France")
	testHashTable.Insert(1950, "Brazil")
	testHashTable.Insert(1954, "Germany")
	testHashTable.Insert(1958, "Italy")

	fmt.Println(testHashTable)
	fmt.Println(testHashTable.Keys())

	fmt.Println(testHashTable.Get(1960))
	fmt.Println(testHashTable.Get(1942))

	testHashTable.Delete(1938)
	testHashTable.Delete(1942)

	fmt.Println(testHashTable)
	fmt.Println(testHashTable.Keys())

}

func (h *HashTable) Insert(key, value interface{}) {

	for i := range h.data[hash(key)] {
		if h.data[hash(key)][i][0] == key {
			log.Println("Key already exists")
			return
		}
	}
	h.data[hash(key)] = append(h.data[hash(key)], []interface{}{key, value})
}

func (h *HashTable) Get(key interface{}) interface{} {

	for i := range h.data[hash(key)] {
		if h.data[hash(key)][i][0] == key {
			return h.data[hash(key)][i][1]
		}
	}
	return nil
}
func (h *HashTable) Delete(key interface{}) {
	for i := range h.data[hash(key)] {
		if h.data[hash(key)][i][0] == key {
			h.data[hash(key)] = append(h.data[hash(key)][:i], h.data[hash(key)][i+1:]...)
			return
		}
	}

}
func (h *HashTable) Keys() []interface{} {

	keys := []interface{}{}
	for i := range h.data {
		if h.data[i] != nil {
			for j := range h.data[i] {
				if h.data[i][j] != nil {
					keys = append(keys, h.data[i][j][0])
				}
			}
		}
	}
	return keys

}
func hash(key interface{}) int {
	keyStr := fmt.Sprintf("%v", key)
	total := 0
	for i := range keyStr {
		total += int(rune(keyStr[i]))
	}
	return total % ArraySize
}
