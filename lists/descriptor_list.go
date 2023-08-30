package lists

import "fmt"

type (
	DescriptorListNode[T Value] struct {
		value T
		next  *DescriptorListNode[T]
	}

	Descriptor[T Value] struct {
		first            *DescriptorListNode[T]
		numberOfElements int
		last             *DescriptorListNode[T]
	}

	IDescriptorList[T Value] interface {
		Unshift(value T)
		Push(value T)
		IsEmpty() bool
		Remove(value T)
		Print()
		NumberOfElements() int
	}
)

func DescriptorList[T Value]() *Descriptor[T] {
	return &Descriptor[T]{
		first:            nil,
		numberOfElements: 0,
		last:             nil,
	}
}

func (list *Descriptor[T]) Unshift(value T) {
	newNode := DescriptorListNode[T]{value, nil}
	if list.IsEmpty() {
		list.first = &newNode
		list.last = &newNode
	} else {
		newNode.next = list.first
		list.first = &newNode
	}
	list.numberOfElements++
}

func (list *Descriptor[T]) Push(value T) {
	newNode := DescriptorListNode[T]{value, nil}
	if list.IsEmpty() {
		list.first = &newNode
		list.last = &newNode
	} else {
		list.last.next = &newNode
		list.last = &newNode
	}
	list.numberOfElements++
}

func (list *Descriptor[T]) IsEmpty() bool {
	return list.numberOfElements == 0
}

func (list *Descriptor[T]) Print() {
	currentNode := list.first
	for currentNode != nil {
		fmt.Print(currentNode.value)
		fmt.Print(" ")
		currentNode = currentNode.next
	}
	fmt.Println()
}

func (list *Descriptor[T]) NumberOfElements() int {
	return list.numberOfElements
}

func (list *Descriptor[T]) Remove(value T) {
	if list.IsEmpty() {
		return
	}
	var currentNode, previousNode *DescriptorListNode[T]
	for currentNode = list.first; currentNode != nil && currentNode.value != value; currentNode = currentNode.next {
		previousNode = currentNode
	}
	if currentNode == nil {
		return
	}
	if previousNode == nil {
		list.first = currentNode.next
		return
	}
	previousNode.next = currentNode.next
}
