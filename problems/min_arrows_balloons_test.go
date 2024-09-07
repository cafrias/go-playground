package problems

import (
	"testing"
)

func TestMinArrowsBalloons(t *testing.T) {
	tests := []struct {
		name     string
		a        [][]int
		expected int
	}{
		{
			"Case 1",
			[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
			2,
		},
		{
			"Case 2",
			[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			4,
		},
		{
			"Case 3",
			[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinArrowsBalloons(tt.a)
			if tt.expected != got {
				t.Fatalf("expected %v got %v", tt.expected, got)
			}
		})
	}
}
