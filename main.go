package main

import (
	"fmt"

	"github.com/JoVictorNunes/data-structures/lists"
	"github.com/JoVictorNunes/data-structures/trees"
)

func main() {
	fmt.Println("------------- LISTS -------------")
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
}
