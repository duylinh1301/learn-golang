package main

import (
	"fmt"
)

func main() {
	list := new(DoublyLinkedList)

	list.addLast("Linh", 24)

	list.addLast("Duy", 18)

	list.addFirst("Nguyen", 20)

	list.addLast("Trong", 22)

	list.show()
}

// Node data structure
type Node struct {
	name string
	age  int
	head *Node
	tail *Node
}

// DoublyLinkedList list store linkedlist
type DoublyLinkedList struct {
	name string
	node *Node
}

func (list *DoublyLinkedList) addNilList(node Node) bool {
	if list.node == nil {
		list.node = &node
		return true
	}
	return false
}

func (list *DoublyLinkedList) addFirst(name string, age int) bool {

	node := Node{
		name: name,
		age:  age,
	}

	if list.addNilList(node) {
		return true
	}

	node.tail = list.node
	list.node.head = &node
	list.node = &node

	return true
}

func (list *DoublyLinkedList) addLast(name string, age int) bool {

	node := Node{
		name: name,
		age:  age,
	}

	if list.addNilList(node) {
		return true
	}

	index := list.node
	for index != nil {

		if index.tail == nil {

			// add last
			index.tail = &node
			node.head = index

			break
		}

		index = index.tail
	}

	return true
}

func (list *DoublyLinkedList) show() {

	if list == nil || list.node == nil {
		fmt.Printf("list empty!")

		return
	}

	index := list.node

	for index != nil {

		fmt.Printf("***---------***\n")

		fmt.Printf("name: %s \n", index.name)

		fmt.Printf("age: %d\n", index.age)

		if index.tail == nil {
			break
		}

		index = index.tail
	}

	return
}
