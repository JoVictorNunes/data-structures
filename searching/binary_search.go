package searching

type (
	BinarySearchValue interface {
		~float32 | ~float64
	}

	BinarySearchNotFoundError struct {
		msg string
	}
)

func (o *BinarySearchNotFoundError) Error() string {
	return o.msg
}

func BinarySearch[V BinarySearchValue](value V, list []V) (element V, err error) {
	left, right := 0, len(list)-1

	for left <= right {
		half := (left + right) / 2

		if value < list[half] {
			right = half - 1
			continue
		}
		if value > list[half] {
			left = half + 1
			continue
		}

		return list[half], nil
	}

	return -1, &BinarySearchNotFoundError{"Element not found"}
}
