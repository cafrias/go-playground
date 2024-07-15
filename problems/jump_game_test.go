package problems

import (
	"fmt"
	"testing"
)

var tests = []struct {
	nums []int
	want bool
}{
	{
		[]int{2, 3, 1, 1, 4},
		true,
	},
	{
		[]int{3, 2, 1, 0, 4},
		false,
	},
	{
		[]int{2, 3, 0, 4, 1},
		true,
	},
	{
		[]int{8, 0, 0, 0, 0, 1},
		true,
	},
}

func TestJumpGame(t *testing.T) {
	for i, tt := range tests {
		t.Run(
			fmt.Sprintf("%d", i),
			func(t *testing.T) {
				got := JumpGame(tt.nums)
				if got != tt.want {
					t.Fatalf("want=%v, got=%v", tt.want, got)
				}
			},
		)
	}
}

func TestJumpGameSolution2(t *testing.T) {
	for i, tt := range tests {
		t.Run(
			fmt.Sprintf("%d", i),
			func(t *testing.T) {
				got := JumpGameSolution2(tt.nums)
				if got != tt.want {
					t.Fatalf("want=%v, got=%v", tt.want, got)
				}
			},
		)
	}
}
