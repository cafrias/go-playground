package huffman

import (
	"errors"
)

type HuffmanTree struct {
	Frequency uint
	Character rune
	Left      *HuffmanTree
	Right     *HuffmanTree
}

func (t *HuffmanTree) Value() uint {
	return t.Frequency
}

func (t *HuffmanTree) GetCode(char rune) ([]byte, error) {
	code := make([]byte, 0)
	return t.searchCode(char, code)
}

func (t *HuffmanTree) searchCode(char rune, code []byte) ([]byte, error) {
	if t == nil {
		return code, errors.New("not found")
	}

	code = append(code, 0)
	code, err := t.Left.searchCode(char, code)
	if err == nil {
		return code, nil
	} else {
		code = code[:len(code)-1]
	}

	code = append(code, 1)
	code, err = t.Right.searchCode(char, code)
	if err == nil {
		return code, nil
	} else {
		code = code[:len(code)-1]
	}

	if t.Character == *new(rune) {
		return code, err
	}

	if char == t.Character {
		return code, nil
	} else {
		return code, errors.New("not found")
	}
}
