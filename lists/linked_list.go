package lists

import "fmt"

type (
	LinkedListNode struct {
		value int
		next  *LinkedListNode
	}

	SLinkedList struct {
		start *LinkedListNode
	}

	ILinkedList interface {
		IsEmpty() bool
		Unshift(value int, sorted bool)
		SortedPush(value int)
		Search(value int) *LinkedListNode
		Remove(value int)
		Print()
	}
)

func LinkedList() *SLinkedList {
	return &SLinkedList{nil}
}

func (list *SLinkedList) Unshift(value int) {
	newNode := LinkedListNode{value, list.start}
	list.start = &newNode
}

func (list *SLinkedList) SortedPush(value int) {
	var currentNode, previousNode *LinkedListNode
	previousNode = nil

	for currentNode = list.start; currentNode != nil && currentNode.value < value; currentNode = currentNode.next {
		previousNode = currentNode
	}

	newNode := LinkedListNode{value, nil}

	if previousNode == nil {
		newNode.next = list.start
		list.start = &newNode
		return
	}

	newNode.next = currentNode
	previousNode.next = &newNode
}

func (list *SLinkedList) Print() {
	currentNode := list.start

	for currentNode != nil {
		fmt.Print(currentNode.value)
		fmt.Print(" ")
		currentNode = currentNode.next
	}

	fmt.Println()
}

func (list *SLinkedList) IsEmpty() bool {
	return list.start == nil
}

func (list *SLinkedList) Search(value int) *LinkedListNode {
	if list.IsEmpty() {
		return nil
	}

	currentNode := list.start

	for currentNode != nil {
		if currentNode.value == value {
			return currentNode
		}
	}

	return nil
}

func (list *SLinkedList) Remove(value int) {
	if list.IsEmpty() {
		return
	}

	var currentNode, previousNode *LinkedListNode

	for currentNode = list.start; currentNode != nil && currentNode.value != value; currentNode = currentNode.next {
		previousNode = currentNode
	}

	if currentNode == nil {
		return
	}

	if previousNode == nil {
		list.start = currentNode.next
		return
	}

	previousNode.next = currentNode.next
}
