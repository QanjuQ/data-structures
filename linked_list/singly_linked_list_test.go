package linkedlist_test

import (
	ll "data-structures/linked_list"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestSinglyLinkedList struct {
	suite.Suite
}

func TestSinglyLinkedList_InsertAt_AtStartWhenListIsEmpty(t *testing.T) {
	list := ll.LinkedList{}
	list.InsertAt(1, 10)
	expected := ll.LinkedList{Head: &ll.Node{Value: 10}, Length: 1}
	assert.Exactly(t, expected, list)
}

func TestSinglyLinkedList_InsertAt_AtStartWhenListHasOneElement(t *testing.T) {
	list := ll.LinkedList{Head: &ll.Node{Value: 10}, Length: 1}
	list.InsertAt(1, 20)
	expected := ll.LinkedList{Head: &ll.Node{Value: 20, Next: &ll.Node{Value: 10}}, Length: 2}
	assert.Exactly(t, expected, list)
}

func TestSinglyLinkedList_InsertAt_AtStartWhenListHasTwoElement(t *testing.T) {
	list := ll.LinkedList{Head: &ll.Node{Value: 10, Next: &ll.Node{Value: 30}}, Length: 2}
	list.InsertAt(1, 20)
	expected := ll.LinkedList{Head: &ll.Node{Value: 20, Next: &ll.Node{Value: 10, Next: &ll.Node{Value: 30}}}, Length: 3}
	assert.Exactly(t, expected, list)
}

func TestSinglyLinkedList_InsertAt_AtEndWhenListHas2Elements(t *testing.T) {
	list := ll.LinkedList{Head: &ll.Node{Value: 20, Next: &ll.Node{Value: 10}}, Length: 2}
	list.InsertAt(3, 30)
	expected := ll.LinkedList{Head: &ll.Node{Value: 20,
		Next: &ll.Node{Value: 10, Next: &ll.Node{Value: 30}}}, Length: 3}
	assert.Exactly(t, expected, list)
}

func TestSinglyLinkedList_InsertAt_AtLengthWhenPositionEqualToLength(t *testing.T) {
	list := ll.LinkedList{Head: &ll.Node{Value: 20, Next: &ll.Node{Value: 10}}, Length: 2}
	list.InsertAt(2, 30)
	expected := ll.LinkedList{Head: &ll.Node{Value: 20,
		Next: &ll.Node{Value: 30, Next: &ll.Node{Value: 10}}}, Length: 3}
	assert.Exactly(t, expected, list)
}

func TestSinglyLinkedList_InsertAt_InBetweenWhenListHasElements(t *testing.T) {
	list := ll.LinkedList{Head: &ll.Node{Value: 10, Next: &ll.Node{Value: 20, Next: &ll.Node{Value: 30}}}, Length: 3}
	list.InsertAt(2, 16)
	expected := ll.LinkedList{Head: &ll.Node{Value: 10, Next: &ll.Node{Value: 16, Next: &ll.Node{Value: 20, Next: &ll.Node{Value: 30}}}}, Length: 4}
	assert.Exactly(t, expected, list)
}

func TestSinglyLinkedList_InsertAt_ErrorWhenPositionIsNegative(t *testing.T) {
	list := ll.LinkedList{Head: &ll.Node{Value: 10, Next: &ll.Node{Value: 20, Next: &ll.Node{Value: 30}}}, Length: 3}
	err := list.InsertAt(-1, 16)
	assert.Equal(t, err.Error(), "invalid position")
	err = list.InsertAt(0, 16)
	assert.Equal(t, err.Error(), "invalid position")
	expected := ll.LinkedList{Head: &ll.Node{Value: 10, Next: &ll.Node{Value: 20, Next: &ll.Node{Value: 30}}}, Length: 3}
	assert.Exactly(t, expected, list)
}

func TestSinglyLinkedList_InsertAt_ErrorWhenPositionIsGreaterThanLengthPlus1(t *testing.T) {
	list := ll.LinkedList{Head: &ll.Node{Value: 10, Next: &ll.Node{Value: 20, Next: &ll.Node{Value: 30}}}, Length: 3}
	err := list.InsertAt(5, 16)
	assert.Equal(t, err.Error(), "invalid position")
	expected := ll.LinkedList{Head: &ll.Node{Value: 10, Next: &ll.Node{Value: 20, Next: &ll.Node{Value: 30}}}, Length: 3}
	assert.Exactly(t, expected, list)
}

func TestSinglyLinkedList_DeleteAt_WhenListIsEmpty(t *testing.T) {
	list := ll.LinkedList{}
	err := list.DeleteAt(1)
	assert.Equal(t, err.Error(), "invalid position")
	assert.Exactly(t, list, ll.LinkedList{})
}

func TestSinglyLinkedList_DeleteAt_StartWhenListHasOneElement(t *testing.T) {
	list := ll.LinkedList{Head: &ll.Node{Value: 10}, Length: 1}
	list.DeleteAt(1)
	assert.Exactly(t, ll.LinkedList{}, list)
}

func TestSinglyLinkedList_DeleteAt_AtEndWhenListHas2Elements(t *testing.T) {
	list := ll.LinkedList{Head: &ll.Node{Value: 20, Next: &ll.Node{Value: 10}}, Length: 2}
	list.DeleteAt(2)
	list.PrintAll()
	expected := ll.LinkedList{Head: &ll.Node{Value: 20}, Length: 1}
	assert.Exactly(t, expected, list)
}

func TestSinglyLinkedList_DeleteAt_InBetweenWhenListHasElements(t *testing.T) {
	list := ll.LinkedList{Head: &ll.Node{Value: 10, Next: &ll.Node{Value: 20, Next: &ll.Node{Value: 30}}}, Length: 3}
	list.DeleteAt(2)
	expected := ll.LinkedList{Head: &ll.Node{Value: 10, Next: &ll.Node{Value: 30}}, Length: 2}
	assert.Exactly(t, expected, list)
}

func TestSinglyLinkedList_DeleteAt_ErrorWhenPositionIsNegative(t *testing.T) {
	list := ll.LinkedList{Head: &ll.Node{Value: 10, Next: &ll.Node{Value: 20, Next: &ll.Node{Value: 30}}}, Length: 3}
	err := list.DeleteAt(-1)
	assert.Equal(t, err.Error(), "invalid position")
	err = list.InsertAt(0, 16)
	assert.Equal(t, err.Error(), "invalid position")
	expected := ll.LinkedList{Head: &ll.Node{Value: 10, Next: &ll.Node{Value: 20, Next: &ll.Node{Value: 30}}}, Length: 3}
	assert.Exactly(t, expected, list)
}

func TestSinglyLinkedList_DeleteAt_ErrorWhenPositionIsGreaterThanLength(t *testing.T) {
	list := ll.LinkedList{Head: &ll.Node{Value: 10, Next: &ll.Node{Value: 20, Next: &ll.Node{Value: 30}}}, Length: 3}
	err := list.DeleteAt(4)
	assert.Equal(t, err.Error(), "invalid position")
	expected := ll.LinkedList{Head: &ll.Node{Value: 10, Next: &ll.Node{Value: 20, Next: &ll.Node{Value: 30}}}, Length: 3}
	assert.Exactly(t, expected, list)
}
