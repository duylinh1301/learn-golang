package main

import "fmt"

func main() {
	var list *LinkedList

	list = &LinkedList{}

	element1 := LinkedList{
		data: "item 1",
		next: nil,
	}

	element2 := LinkedList{
		data: "item 2",
		next: nil,
	}

	list.add(&element1)

	list.add(&element2)

	show(list)

}

// LinkedList danh sach lien ket
type LinkedList struct {
	data string
	next *LinkedList
}

func (list *LinkedList) add(item *LinkedList) bool {
	if list == nil {
		list = item
	} else {
		for list.next != nil {
			list = list.next
		}
		list.next = item
	}

	return true
}

func show(list *LinkedList) bool {
	var index *LinkedList
	index = list

	for index != nil {
		fmt.Println(index.data)
		index = index.next
	}
	return true
}
