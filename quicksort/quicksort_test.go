package quicksort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{9, 3, 7, 4, 69, 420, 42},
			expected: []int{3, 4, 7, 9, 42, 69, 420},
		},
		{
			input:    []int{0, 1, 2, 3, 4, 5},
			expected: []int{0, 1, 2, 3, 4, 5},
		},
		{
			input:    []int{5, 4, 3, 2, 1, 0},
			expected: []int{0, 1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.input)
		t.Run(testname, func(t *testing.T) {
			QuickSort(tt.input)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("not sorted: got %v, want %v", tt.input, tt.expected)
			}
		})
	}
}

func TestPartition(t *testing.T) {
	var tests = []struct {
		arr       []int
		low       int
		high      int
		pivotIdx  int
		resultArr []int
	}{
		{
			[]int{9, 3, 7, 4, 69, 420, 42},
			0,
			6,
			4,
			[]int{9, 3, 7, 4, 42, 420, 69},
		},
		{
			[]int{9, 3, 7, 4, 69, 420, 900},
			0,
			6,
			6,
			[]int{9, 3, 7, 4, 69, 420, 900},
		},
		{
			[]int{9, 3, 7, 4, 69, 420, 2},
			0,
			6,
			0,
			[]int{2, 3, 7, 4, 69, 420, 9},
		},
		{
			[]int{9, 3, 7, 4, 5, 420, 42},
			1,
			4,
			3,
			[]int{9, 3, 4, 5, 7, 420, 42},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v,%v", tt.arr, tt.low, tt.high)
		t.Run(testname, func(t *testing.T) {
			result := partition(tt.arr, tt.low, tt.high)
			if result != tt.pivotIdx {
				t.Errorf("pivot index: got %d, want %d", result, tt.pivotIdx)
			}

			if !reflect.DeepEqual(tt.arr, tt.resultArr) {
				t.Errorf("not sorted: got %v, want %v", tt.arr, tt.resultArr)
			}
		})
	}
}
