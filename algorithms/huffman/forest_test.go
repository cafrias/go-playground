package huffman

import (
	"errors"
	"go-playground/data-structures/minheap"
	"strings"
	"testing"
)

func TestCreateFrequencyMap(t *testing.T) {
	input := strings.NewReader("a, bb, ccc, dddd, 英英文版本版版版")
	res := createForest(input)

	exp := [][]HuffmanTree{
		{
			HuffmanTree{Frequency: 1, Character: 'a'},
			HuffmanTree{Frequency: 1, Character: '文'},
			HuffmanTree{Frequency: 1, Character: '本'},
		},
		{
			HuffmanTree{Frequency: 2, Character: 'b'},
			HuffmanTree{Frequency: 2, Character: '英'},
		},
		{
			HuffmanTree{Frequency: 3, Character: 'c'},
		},
		{
			HuffmanTree{Frequency: 4, Character: ','},
			HuffmanTree{Frequency: 4, Character: 'd'},
			HuffmanTree{Frequency: 4, Character: ' '},
			HuffmanTree{Frequency: 4, Character: '版'},
		},
	}

	idx := 0
	for i, err := res.Delete(); !errors.Is(err, minheap.ErrHeapEmpty); i, err = res.Delete() {
		var checkIdx int
		if idx < 3 {
			checkIdx = 0
		} else if idx < 5 {
			checkIdx = 1
		} else if idx < 6 {
			checkIdx = 2
		} else {
			checkIdx = 3
		}
		idx++

		var ok bool
		for _, v := range exp[checkIdx] {
			if v.Frequency == i.Frequency && v.Character == i.Character {
				ok = true
			}
		}

		if !ok {
			t.Errorf("Did not find %v in %v", *i, exp[checkIdx])
			t.FailNow()
		}
	}
}
