package main

import (
	"fmt"
)

func main() {

	linkedList := new(LinkedList)

	// linkedList.deleteHeadItem()

	linkedList.addTailItem(40)

	linkedList.addTailItem(15)

	linkedList.addTailItem(25)

	linkedList.addTailItem(30)

	linkedList.addTailItem(50)

	linkedList.deleteItemWithValue(40)
	linkedList.deleteItemWithValue(15)
	linkedList.deleteItemWithValue(25)

	// linkedList.deleteTailItem()

	linkedList.addHeadItem(5)

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

// deleteHeadItem function for delete first item of the list
func (l *LinkedList) deleteHeadItem() bool {

	if l.head == nil {
		return true
	}

	l.head = l.head.next

	return true
}

// deleteHeadItem function for delete last item of the list
func (l *LinkedList) deleteTailItem() bool {

	index := l.head
	var preIndex *Node

	for index.next != nil {
		preIndex = index
		index = index.next
	}

	index = nil
	preIndex.next = nil

	return true
}

func (l *LinkedList) deleteItemWithValue(value int) bool {
	if l.head == nil {
		return true
	}

	doDelete := false

	index := l.head

	var preIndex *Node

	for {
		if index.data == value {
			doDelete = true

			break
		}

		if index.next == nil {
			break
		}

		preIndex = index

		index = index.next
	}

	if doDelete {

		if preIndex == nil {
			l.head = index.next

			index = nil
		} else {
			preIndex.next = index.next

			index.next = nil
		}

	}

	return true
}

// show function print data of the list
func (l *LinkedList) show() bool {

	if l == nil || l.head == nil {
		fmt.Println("List empty!")
		return false
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
