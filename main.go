package main

import "fmt"

type node struct {
	data int
	next *node
}

type singlyLinkedList struct {
	head   *node
	length int
}

func (l *singlyLinkedList) prepend(n *node) {
	newHead := l.head
	l.head = n
	l.head.next = newHead
	l.length++
}

func (l singlyLinkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}

	fmt.Printf("\n")
}

func (l *singlyLinkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head

	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}

		previousToDelete = previousToDelete.next

	}

	previousToDelete.next = previousToDelete.next.next
	l.length--

}

func main() {
	myList := singlyLinkedList{}

	node1 := &node{data: 48}
	node2 := &node{data: 32}
	node3 := &node{data: 16}
	node4 := &node{data: 8}
	node5 := &node{data: 4}
	node6 := &node{data: 2}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)

	myList.printListData()

	myList.deleteWithValue(16)

	myList.printListData()
}
