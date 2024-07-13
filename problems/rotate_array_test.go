package problems

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotateArray(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			[]int{1, 2, 3, 4, 5},
			8,
			[]int{3, 4, 5, 1, 2},
		},
		{
			[]int{-1, -100, 3, 99},
			2,
			[]int{3, 99, -1, -100},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			inp := make([]int, len(tt.nums))
			copy(inp, tt.nums)
			RotateArray(inp, tt.k)
			if !reflect.DeepEqual(inp, tt.want) {
				t.Fatalf("expected %v, got %v", tt.want, inp)
			}
		})
	}
}
