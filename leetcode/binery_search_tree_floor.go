package main

import "fmt"

func main() {

}

type Node struct {
	Key   int
	Value int
	Left  *Node
	Right *Node
}

func floor(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.Key == key {
		return node
	}
	if key < node.Key {
		return floor(node.Left, key)
	}
	temp := floor(node.Right, key)
	if temp != nil {
		return temp
	}
	fmt.Println("")
	return node
}

func buildTree() *Node {
	head := new(Node)

	head.Key = 39
	head.Left = nil

	return nil
}
