package quicksort

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkQuickSort(b *testing.B) {
	g := rand.New(rand.NewSource(123))
	lengths := [...]int{1000, 10_000, 100_000, 1_000_000}
	data := make([][]int, 0, len(lengths))
	for _, l := range lengths {
		newArr := make([]int, l)
		for i := 0; i < l; i++ {
			newArr[i] = g.Int()
		}
		data = append(data, newArr)
	}

	for i := 0; i < len(data); i++ {
		input := data[i]
		name := fmt.Sprintf("with %d items", len(input))
		b.Run(name, func(b *testing.B) {
			inp := make([]int, 0, len(input))
			copy(inp, input)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				QuickSort(inp)
			}
		})
	}
}
