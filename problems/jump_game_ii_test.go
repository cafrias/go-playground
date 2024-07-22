package problems

import (
	"fmt"
	"testing"
)

func TestJumpGameII(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			[]int{
				2, 3, 1, 1, 4,
			},
			2,
		},
		{
			[]int{
				2, 3, 0, 1, 4,
			},
			2,
		},
		{
			[]int{
				1, 1, 1, 1, 1,
			},
			4,
		},
		{
			[]int{
				1, 2, 1, 1, 1,
			},
			3,
		},
		{
			[]int{
				1, 1, 6, 1, 1,
			},
			3,
		},
		{
			[]int{
				5,
			},
			0,
		},
	}

	for i, tt := range tests {
		t.Run(
			fmt.Sprintf("%d", i),
			func(t *testing.T) {
				got := JumpGameII(tt.nums)
				if got != tt.want {
					t.Fatalf("expected=%d, got=%d", tt.want, got)
				}
			},
		)
	}
}
