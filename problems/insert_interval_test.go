package problems

import (
	"reflect"
	"testing"
)

func TestInsertInterval(t *testing.T) {
	tests := []struct {
		name     string
		a        [][]int
		b        []int
		expected [][]int
	}{
		{
			"Case 1",
			[][]int{{1, 3}, {6, 9}},
			[]int{2, 5},
			[][]int{{1, 5}, {6, 9}},
		},
		{
			"Case 2",
			[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			[]int{4, 8},
			[][]int{{1, 2}, {3, 10}, {12, 16}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InsertInterval(tt.a, tt.b)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Fatalf("expected %v got %v", tt.expected, got)
			}
		})
	}
}
