package huffman

import (
	"fmt"
	"go-playground/data-structures/minheap"
	"io"
	"text/scanner"
)

func createForest(file io.Reader) minheap.MinHeap[uint, *HuffmanTree] {
	var s scanner.Scanner
	s.Init(file)

	m := make(map[rune]uint, 1024)

	for tok := s.Next(); tok != scanner.EOF; tok = s.Next() {
		if _, ok := m[tok]; !ok {
			m[tok] = 1
		} else {
			m[tok] += 1
		}
	}

	fmt.Println(m)
	res := minheap.MinHeap[uint, *HuffmanTree]{
		Items: make([]*HuffmanTree, 0, len(m)),
	}

	for c, f := range m {
		res.Insert(
			&HuffmanTree{
				Frequency: f,
				Character: c,
			},
		)
	}

	return res
}
