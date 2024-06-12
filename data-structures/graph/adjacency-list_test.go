package graph

import (
	"fmt"
	"reflect"
	"testing"
)

func TestListDFS(t *testing.T) {
	tests := []struct {
		list     AdjacencyList
		start    int
		needle   int
		expected []int
	}{
		{
			list:   createList1(),
			start:  0,
			needle: 3,
			expected: []int{
				0, 5, 2, 1,
			},
		},
		{
			list:   createList2(),
			start:  0,
			needle: 6,
			expected: []int{
				0,
				1,
				4,
				5,
			},
		},
		{
			list:   createList3(),
			start:  0,
			needle: 5,
			expected: []int{
				0,
				3,
			},
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", idx), func(t *testing.T) {
			path, err := AdjacencyListDFS(tt.list, tt.start, tt.needle)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				t.FailNow()
			}

			if !reflect.DeepEqual(path, tt.expected) {
				t.Errorf("Invalid path\ngot %v\nexpected %v", path, tt.expected)
				t.FailNow()
			}

		})
	}

}

func createList1() AdjacencyList {
	m := NewAdjacencyList(6)
	m[0] = []Edge{
		{To: 5, Weight: 3},
	}
	m[1] = []Edge{
		{To: 0, Weight: 6},
		{To: 3, Weight: 5},
	}
	m[2] = []Edge{
		{To: 0, Weight: 1},
		{To: 1, Weight: 2},
		{To: 4, Weight: 4},
	}
	m[3] = []Edge{
		{To: 4, Weight: 2},
	}
	m[4] = []Edge{
		{To: 2, Weight: 2},
	}
	m[5] = []Edge{
		{To: 2, Weight: 2},
		{To: 4, Weight: 1},
	}

	return m
}

func createList2() AdjacencyList {
	l := NewAdjacencyList(7)

	//     >(1)<--->(4) ---->(5)
	//    /          |       /|
	// (0)     ------|------- |
	//    \   v      v        v
	//     >(2) --> (3) <----(6)
	l[0] = []Edge{
		{To: 1, Weight: 3},
		{To: 2, Weight: 1},
	}
	l[1] = []Edge{{To: 4, Weight: 1}}
	l[2] = []Edge{{To: 3, Weight: 7}}
	l[3] = []Edge{}
	l[4] = []Edge{{To: 1, Weight: 1},
		{To: 3, Weight: 5},
		{To: 5, Weight: 2}}
	l[5] = []Edge{{To: 2, Weight: 18},
		{To: 6, Weight: 1}}
	l[6] = []Edge{{To: 3, Weight: 1}}

	return l
}

func createList3() AdjacencyList {
	l := NewAdjacencyList(6)

	l[0] = []Edge{
		{To: 1, Weight: 1},
		{To: 2, Weight: 1},
		{To: 3, Weight: 1},
	}
	l[1] = []Edge{}
	l[2] = []Edge{
		{To: 4, Weight: 1},
	}
	l[3] = []Edge{
		{To: 5, Weight: 1},
	}
	l[4] = []Edge{}
	l[5] = []Edge{}

	return l
}
