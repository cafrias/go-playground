package problems

import (
	"fmt"
	"testing"
)

func TestHIndex(t *testing.T) {
	tests := []struct {
		input []int
		exp   int
	}{
		{
			[]int{3, 0, 6, 1, 5},
			3,
		},
		{
			[]int{1, 3, 1},
			1,
		},
		{
			[]int{1},
			1,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			res := Hindex(tt.input)
			if res != tt.exp {
				t.Fatalf("expected %d, got %d", tt.exp, res)
			}
		})
	}

}
