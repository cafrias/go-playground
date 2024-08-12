package problems

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs []string
		want [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			[]string{"a"},
			[][]string{
				{"a"},
			},
		},
		{
			[]string{""},
			[][]string{
				{""},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			got := GroupAnagrams(tt.strs)

			if !compareSlices(tt.want, got) {
				t.Fatalf("want=%v, got=%v", tt.want, got)
			}
		})
	}

}

func compareSlices(a [][]string, b [][]string) bool {
	aC := make([][]string, len(a))
	copy(aC, a)
	bC := make([][]string, len(b))
	copy(bC, b)

	slices.SortFunc(aC, func(a, b []string) int {
		return len(b) - len(a)
	})
	slices.SortFunc(bC, func(a, b []string) int {
		return len(b) - len(a)
	})

	for _, v := range aC {
		slices.Sort(v)
	}

	for _, v := range bC {
		slices.Sort(v)
	}

	fmt.Println(aC, bC)

	return reflect.DeepEqual(aC, bC)
}
