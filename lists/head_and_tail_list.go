package lists

type (
	HeadAndTailListNode[T Value] struct {
		value T
		next  *HeadAndTailListNode[T]
	}

	SHeadAndTailList[T Value] struct {
		head *HeadAndTailListNode[T]
		tail *HeadAndTailListNode[T]
	}

	IHeadAndTailList[T Value] interface {
		Unshift(value T)
		Push(value T)
		Print()
		IsEmpty() bool
		Remove(value T)
	}
)

func HeadAndTailList[T Value]() *SHeadAndTailList[T] {
	tail := HeadAndTailListNode[T]{0, nil}
	head := HeadAndTailListNode[T]{0, &tail}
	return &SHeadAndTailList[T]{&head, &tail}
}

func (list *SHeadAndTailList[T]) Unshift(value T) {}

func (list *SHeadAndTailList[T]) Push(value T) {}

func (list *SHeadAndTailList[T]) Print() {}

func (list *SHeadAndTailList[T]) IsEmpty() bool {
	return true
}

func (list *SHeadAndTailList[T]) Remove(value T) {}
