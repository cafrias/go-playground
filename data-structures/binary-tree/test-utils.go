package binarytree

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
