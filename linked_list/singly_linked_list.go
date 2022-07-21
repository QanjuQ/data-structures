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
	head *Node
}

func (l *LinkedList) isEmpty() bool {
	return l.head == nil
}

func (l *LinkedList) InsertAt(position, value int) {
	new := &Node{value: value}
	if l.isEmpty() {
		l.head = new
		return
	}
	curr := l.head
	for index := 0; index < position; index++ {
		if curr.next == nil {
			curr.next = new
			return
		}
		if curr == nil {
			break
		}
		curr = curr.next
	}

	temp := &Node{next: curr.next, value: curr.value}
	curr.value = new.value
	curr.next = temp
}

func (l *LinkedList) DeleteAt(position int) {
	var prev *Node
	curr := l.head
	for index := 0; index < position; index++ {
		prev = curr
		if curr != nil {
			curr = curr.next
		}
		if curr.next == nil {
			prev.next = nil
			return
		}
	}
	if prev != nil {
		prev.next = curr.next
	} else {
		l.head = curr.next
	}
}

func (l *LinkedList) PrintAll() {
	currNode := l.head
	if l.isEmpty() {
		fmt.Println("empty linked list")
		return
	}
	for currNode != nil {
		fmt.Printf("%d ", currNode.value)
		currNode = currNode.next
	}
	fmt.Println("")
}

func (l *LinkedList) Exists(nodeValue int) (bool, int) {
	currNode := l.head
	position := 0
	for currNode != nil {
		if currNode.value == nodeValue {
			return true, position
		}
		currNode = currNode.next
		position++
	}
	return false, -1
}

func (l *LinkedList) Sort() {
	current := l.head
	for current != nil {
		index := current.next
		for index != nil {
			if current.value > index.value {
				temp := current.value
				current.value = index.value
				index.value = temp
			} else {
				index = index.next
			}
		}
		current = current.next
	}
}

func (l *LinkedList) Reverse() *LinkedList {
	reversed := &LinkedList{}
	curr := l.head
	for curr != nil {
		reversed.head = &Node{value: curr.value, next: reversed.head}
		curr = curr.next
	}
	return reversed
}
