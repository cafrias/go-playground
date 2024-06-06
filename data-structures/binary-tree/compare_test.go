package binarytree

import (
	"fmt"
	"testing"
)

func TestCompare(t *testing.T) {
	var tests = []struct {
		a, b BinaryTree[int]
		want bool
	}{
		{
			a:    createTree(),
			b:    createTree(),
			want: true,
		},
		{
			a: createTree(),
			b: BinaryTree[int]{
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
							Value: 900,
						},
					},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			result := tt.a.Equals(&tt.b)
			if result != tt.want {
				t.Errorf("wanted %v got %v", tt.want, result)
			}
		})
	}
}
