package trees

import "fmt"

type (
	Value interface {
		int | float64 | float32
	}

	BinarySearchTree[T Value] struct {
		value T
		left  *BinarySearchTree[T]
		right *BinarySearchTree[T]
	}

	IBinarySearchTree[T Value] interface {
		Search(value T) *BinarySearchTree[T]
		Insert(value T) *BinarySearchTree[T]
		Remove(value T) *BinarySearchTree[T]
		IsEmpty() bool
		Print()
	}
)

func (tree *BinarySearchTree[T]) IsEmpty() bool {
	return tree == nil
}

func (tree *BinarySearchTree[T]) Search(value T) *BinarySearchTree[T] {
	if tree.IsEmpty() {
		return nil
	}
	if value == tree.value {
		return tree
	}
	if value > tree.value {
		return tree.right.Search(value)
	}
	return tree.left.Search(value)
}

func (tree *BinarySearchTree[T]) Insert(value T) *BinarySearchTree[T] {
	if tree.IsEmpty() {
		tree = &BinarySearchTree[T]{value, nil, nil}
	} else if value < tree.value {
		tree.left = tree.left.Insert(value)
	} else if value > tree.value {
		tree.right = tree.right.Insert(value)
	}
	return tree
}

func (tree *BinarySearchTree[T]) Print() {
	if tree.IsEmpty() {
		return
	}
	tree.left.Print()
	fmt.Printf("%v ", tree.value)
	tree.right.Print()
	fmt.Println()
}

func (tree *BinarySearchTree[T]) Remove(value T) *BinarySearchTree[T] {
	if tree.IsEmpty() {
		return nil
	}
	if value < tree.value {
		tree.left = tree.left.Remove(value)
	} else if value > tree.value {
		tree.right = tree.right.Remove(value)
	} else {
		if tree.left.IsEmpty() && tree.right.IsEmpty() {
			tree = nil
		} else if tree.left.IsEmpty() {
			tree = tree.right
		} else if tree.right.IsEmpty() {
			tree = tree.left
		} else {
			temp := tree.left
			for temp.right != nil {
				temp = temp.right
			}
			tree.value = temp.value
			temp.value = value
			tree.left = tree.left.Remove(value)
		}
	}
	return tree
}
