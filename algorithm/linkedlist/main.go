package main

import (
	"fmt"
)

func main() {

	linkedList := new(LinkedList)

	linkedList.addTailItem(10)

	linkedList.addTailItem(25)

	linkedList.addTailItem(40)

	linkedList.addTailItem(5)

	linkedList.addTailItem(0)

	linkedList.addTailItem(15)

	linkedList.show()

}

// LinkedList is a is a list store many node
type LinkedList struct {
	name string

	head *Node
}

func (l *LinkedList) addTailItem(data int) bool {
	node := new(Node)

	node.data = data

	if l.head == nil {
		l.head = node
	} else {
		var index *Node

		index = l.head

		for index.next != nil {
			index = index.next
		}

		index.next = node
	}

	return true
}

func (l *LinkedList) show() bool {

	if l == nil {
		fmt.Println("List empty!")
	}

	var index *Node

	index = l.head

	for {
		fmt.Println(index.data)

		if index.next == nil {
			break
		}

		index = index.next

	}

	return true
}

// Node is a struct type for one element of linked list
type Node struct {
	data int

	next *Node
}
