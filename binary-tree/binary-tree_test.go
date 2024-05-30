package binarytree

import (
	"reflect"
	"testing"
)

func createTree() BinaryTree[int] {
	return BinaryTree[int]{
		Value: 5,
		Left: &BinaryTree[int]{
			Value: 4,
			Left: &BinaryTree[int]{
				Value: 2,
				Left: &BinaryTree[int]{
					Value: 1,
				},
				Right: &BinaryTree[int]{
					Value: 3,
				},
			},
		},
		Right: &BinaryTree[int]{
			Value: 6,
			Left: &BinaryTree[int]{
				Value: 7,
			},
			Right: &BinaryTree[int]{
				Value: 8,
				Right: &BinaryTree[int]{
					Value: 9,
				},
			},
		},
	}

}

func TestPreOrder(t *testing.T) {
	tree := createTree()
	expected := []int{
		5, 4, 2, 1, 3, 6, 7, 8, 9,
	}

	res := tree.PreOrder()
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v got %v", expected, res)
	}
}

func TestInOrder(t *testing.T) {
	tree := createTree()
	expected := []int{
		1, 2, 3, 4, 5, 7, 6, 8, 9,
	}

	res := tree.InOrder()
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v got %v", expected, res)
	}
}

func TestPostOrder(t *testing.T) {
	tree := createTree()
	expected := []int{
		1, 3, 2, 4, 7, 9, 8, 6, 5,
	}

	res := tree.PostOrder()
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %v got %v", expected, res)
	}
}
