package huffman

import (
	"bytes"
	"fmt"
	"io"
	"text/scanner"
)

func Encode(source io.Reader) (code []byte, tree *HuffmanTree, err error) {
	codeSource := &bytes.Buffer{}
	treeSource := io.TeeReader(source, codeSource)
	tree = createHuffmanTree(treeSource)

	cache := make(map[rune][]byte, 0)
	var s scanner.Scanner
	s.Init(codeSource)

	for rune := s.Next(); rune != scanner.EOF; rune = s.Next() {
		if _, ok := cache[rune]; ok {
			code = append(code, cache[rune]...)
		} else {
			runeCode, err := tree.GetCode(rune)
			if err != nil {
				return nil, nil, fmt.Errorf("Unable to find '%s' in tree", string(rune))
			}
			code = append(code, runeCode...)
			cache[rune] = runeCode
		}
	}

	return code, tree, nil
}
