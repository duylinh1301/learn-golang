package main

import (
	"fmt"
)

func main() {

	linkedList := new(LinkedList)

	linkedList.addTailItem(10)

	linkedList.addHeadItem(100)

	linkedList.addTailItem(25)

	linkedList.addHeadItem(125)

	linkedList.addTailItem(40)

	linkedList.addTailItem(5)

	linkedList.addTailItem(0)

	linkedList.addTailItem(15)

	linkedList.addHeadItem(275)

	linkedList.addHeadItem(150)

	linkedList.addItemWithPosition(500, 20)

	linkedList.show()

}

// LinkedList is a is a list store many node
type LinkedList struct {
	name string

	head *Node
}

// addTailItem function for add item to the tail of the list
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

// addTailItem function for add item to the head of the list
func (l *LinkedList) addHeadItem(data int) bool {
	node := Node{
		data: data,
	}

	if l == nil {
		fmt.Println("Linked list is empty!")
		return false
	}

	if l.head == nil {
		l.head = &node
	} else {
		node.next = l.head

		l.head = &node

	}

	return true
}

// addTailItem function for add item to the position of the list
func (l *LinkedList) addItemWithPosition(data int, position int) bool {
	var index *Node

	node := Node{
		data: data,
	}

	index = l.head

	count := 1

	if position == 0 || l.head == nil {
		node.next = l.head

		l.head = &node

		return true
	}

	for index.next != nil {
		if count == position || position == 0 {
			break
		}

		index = index.next

		count++
	}

	node.next = index.next

	index.next = &node

	return true
}

// show function print data of the list
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
