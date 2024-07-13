package problems

import (
	"fmt"
	"testing"
)

func TestBestStock(t *testing.T) {
	tests := []struct {
		inp  []int
		want int
	}{
		{
			[]int{7, 1, 5, 3, 6, 4},
			7,
		},
		{
			[]int{1, 2, 3, 4, 5},
			4,
		},
		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := BestTimeStock(tt.inp)
			if got != tt.want {
				t.Fatalf("got=%d, want=%d", got, tt.want)
			}
		})
	}
}
