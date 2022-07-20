package linkedlist

import (
	"fmt"
)

// Traversal - access each element of the linked list
// Insertion - adds a new element to the linked list
// Deletion - removes the existing elements
// Search - find a node in the linked list
// Sort - sort the nodes of the linked list
// Insert At - insert at a position
// Delete At - delete at a position

type LinkedList struct {
	head   *Node
	Length int
}

func (l *LinkedList) InsertAt(position, value int) {
	if position > l.Length {
		fmt.Println("position can't be greater than length")
		return
	}
	temp := &Node{value: value}
	if l.head == nil {
		l.head = temp
		l.Length += 1
		return
	}

	currNode := l.head
	for index := 0; index < position-1; index++ {
		if currNode == nil {
			return
		}
		currNode = currNode.next
	}
	temp.next = currNode.next
	currNode.next = temp
	l.Length += 1
	l.PrintAll()
}

func (l *LinkedList) DeleteAt(position int) {
	if l.head == nil {
		fmt.Println("empty list")
		return
	}
	if position > l.Length {
		fmt.Println("position can't be greater than length")
		return
	}
	currNode := l.head
	for index := 0; index < position-2; index++ {
		if currNode == nil {
			return
		}
		currNode = currNode.next
	}
	currNode.value = currNode.next.value
	currNode.next = currNode.next.next
	l.Length -= 1
}

func (l *LinkedList) PrintAll() {
	currNode := l.head
	if l.head == nil {
		fmt.Println("empty linked list")
		return
	}
	for currNode != nil {
		fmt.Printf("%d ", currNode.value)
		currNode = currNode.next
	}
	fmt.Println("")
}
