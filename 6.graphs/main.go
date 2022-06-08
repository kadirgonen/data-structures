package main

import "fmt"

type Graph struct {
	numberOfNodes int
	adjacentList  map[Node][]int
}
type Node struct {
	value int
}

func newGraph() *Graph {
	return &Graph{adjacentList: make(map[Node][]int)}
}
func createNode(value int) *Node {
	return &Node{value: value}
}

func main() {
	myGraph := newGraph()
	myGraph.addVertex(*createNode(0))
	myGraph.addVertex(*createNode(1))
	myGraph.addVertex(*createNode(2))
	myGraph.addVertex(*createNode(3))
	myGraph.addVertex(*createNode(4))
	myGraph.addVertex(*createNode(5))
	myGraph.addVertex(*createNode(6))

	myGraph.addEdge(Node{value: 3}, Node{value: 1})
	myGraph.addEdge(Node{value: 3}, Node{value: 4})
	myGraph.addEdge(Node{value: 4}, Node{value: 2})
	myGraph.addEdge(Node{value: 4}, Node{value: 5})
	myGraph.addEdge(Node{value: 1}, Node{value: 2})
	myGraph.addEdge(Node{value: 1}, Node{value: 0})
	myGraph.addEdge(Node{value: 0}, Node{value: 2})
	myGraph.addEdge(Node{value: 6}, Node{value: 5})
	myGraph.addEdge(Node{value: 7}, Node{value: 8})
	myGraph.addVertex(*createNode(6))
	fmt.Println(myGraph)
	myGraph.showConnections()
}

func (g *Graph) addVertex(n Node) bool {

	if g.adjacentList[n] != nil {
		fmt.Println("the node already exists")
		return false
	}
	g.adjacentList[n] = nil
	g.numberOfNodes++
	return true
}
func (g *Graph) addEdge(n1, n2 Node) {
	g.adjacentList[n1] = append(g.adjacentList[n1], n2.value)
	g.adjacentList[n2] = append(g.adjacentList[n2], n1.value)

}

func (g *Graph) showConnections() {
	for k, v := range g.adjacentList {
		fmt.Println(k, "--->", v)
	}
}
