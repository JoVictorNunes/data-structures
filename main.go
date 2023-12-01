package main

import (
	"fmt"

	"github.com/JoVictorNunes/data-structures/lists"
	"github.com/JoVictorNunes/data-structures/searching"
	"github.com/JoVictorNunes/data-structures/trees"
)

func main() {
	fmt.Println("------------- LINKED LIST -------------")
	linear_list := lists.LinkedList()
	linear_list.Unshift(1)
	linear_list.Unshift(2)
	linear_list.Unshift(3)
	linear_list.Unshift(4)
	linear_list.Print()
	linear_list.Remove(3)
	linear_list.Print()
	linear_list.Unshift(3)
	linear_list.Print()

	fmt.Println("------------- TREES -------------")
	var bst *trees.BinarySearchTree[int] = nil
	bst = bst.Insert(1)
	bst = bst.Insert(2)
	bst = bst.Insert(3)
	bst.Print()
	bst = bst.Remove(2)
	bst.Print()

	fmt.Println("------------- DESCRIPTOR LIST -------------")
	descriptor_list := lists.DescriptorList[int]()
	descriptor_list.Push(1)
	descriptor_list.Push(2)
	descriptor_list.Unshift(3)
	fmt.Println(descriptor_list.NumberOfElements())
	descriptor_list.Print()
	descriptor_list.Remove(2)
	descriptor_list.Print()

	fmt.Println("------------- BINARY SEARCH -------------")
	list := []float32{1, 2, 3, 4, 5, 6, 7}
	elem, err := searching.BinarySearch[float32](9, list)
	fmt.Println(elem, err)
}
