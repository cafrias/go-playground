package binarytree

type BinaryTree[T any] struct {
	Value T
	Left  *BinaryTree[T]
	Right *BinaryTree[T]
}

func (t *BinaryTree[T]) preTraverse(result []T) []T {
	if t == nil {
		return result
	}

	result = append(result, t.Value)
	result = t.Left.preTraverse(result)
	result = t.Right.preTraverse(result)

	return result
}

func (t *BinaryTree[T]) PreOrder() []T {
	result := make([]T, 0)
	result = t.preTraverse(result)
	return result
}

func (t *BinaryTree[T]) inTraverse(result []T) []T {
	if t == nil {
		return result
	}

	result = t.Left.inTraverse(result)
	result = append(result, t.Value)
	result = t.Right.inTraverse(result)

	return result
}

func (t *BinaryTree[T]) InOrder() []T {
	result := make([]T, 0)
	result = t.inTraverse(result)
	return result
}

func (t *BinaryTree[T]) postTraverse(result []T) []T {
	if t == nil {
		return result
	}

	result = t.Left.postTraverse(result)
	result = t.Right.postTraverse(result)
	result = append(result, t.Value)

	return result
}

func (t *BinaryTree[T]) PostOrder() []T {
	result := make([]T, 0)
	result = t.postTraverse(result)
	return result
}
