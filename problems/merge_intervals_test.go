package problems

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		name     string
		a        [][]int
		expected [][]int
	}{
		{
			"Case 1",
			[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			"Case 2",
			[][]int{{1, 4}, {4, 5}},
			[][]int{{1, 5}},
		},
		{
			"Must Sort before",
			[][]int{{2, 6}, {8, 10}, {1, 3}, {15, 18}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			"Next interval is included in previous",
			[][]int{{1, 4}, {2, 3}},
			[][]int{{1, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeIntervals(tt.a)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Fatalf("expected %v got %v", tt.expected, got)
			}
		})
	}
}
