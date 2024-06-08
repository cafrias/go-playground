package huffman

import (
	"go-playground/data-structures/minheap"
	"io"
	"text/scanner"
)

func createHuffmanTree(file io.Reader) *HuffmanTree {
	var s scanner.Scanner
	s.Init(file)

	freq := make(map[rune]uint, 1024)
	for tok := s.Next(); tok != scanner.EOF; tok = s.Next() {
		if _, ok := freq[tok]; !ok {
			freq[tok] = 1
		} else {
			freq[tok] += 1
		}
	}
	if len(freq) == 0 {
		return nil
	}

	forest := minheap.MinHeap[uint, *HuffmanTree]{
		Items: make([]*HuffmanTree, 0, len(freq)),
	}
	for c, f := range freq {
		forest.Insert(
			&HuffmanTree{
				Frequency: f,
				Character: c,
			},
		)
	}

	if forest.Length() < 2 {
		item, _ := forest.Delete()
		root := HuffmanTree{
			Frequency: item.Frequency,
			Left:      item,
		}
		return &root
	}

	for {
		if forest.Length() <= 1 {
			break
		}

		l, _ := forest.Delete()
		r, _ := forest.Delete()

		n := HuffmanTree{
			Frequency: l.Value() + r.Value(),
			Left:      l,
			Right:     r,
		}

		forest.Insert(&n)
	}

	result, _ := forest.Delete()

	return result
}
