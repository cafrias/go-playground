package problems

import (
	"fmt"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{
			[]int{100, 4, 200, 1, 3, 2},
			4,
		},
		{
			[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			9,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{},
			0,
		},
		{
			[]int{10, 9, 8, 5, 6, 7, 3, 2, 1, 0},
			6,
		},
	}

	for i, tt := range tests {
		t.Run(
			fmt.Sprintf("%d", i+1),
			func(t *testing.T) {
				got := LongestConsecutive2(tt.input)
				if got != tt.want {
					t.Fatalf("want=%d, got=%d", tt.want, got)
				}
			},
		)
	}
}
