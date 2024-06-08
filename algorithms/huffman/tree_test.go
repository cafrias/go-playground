package huffman

import (
	"reflect"
	"testing"
)

func TestGetCode(t *testing.T) {
	tree := HuffmanTree{
		Frequency: 15,
		Left: &HuffmanTree{
			Frequency: 6,
			Left: &HuffmanTree{
				Frequency: 3,
				Character: 'c',
			},
			Right: &HuffmanTree{
				Frequency: 3,
				Left: &HuffmanTree{
					Frequency: 1,
					Character: 'a',
				},
				Right: &HuffmanTree{
					Frequency: 2,
					Character: 'b',
				},
			},
		},
		Right: &HuffmanTree{
			Frequency: 9,
			Left: &HuffmanTree{
				Frequency: 4,
				Character: 'd',
			},
			Right: &HuffmanTree{
				Frequency: 5,
				Character: '英',
			},
		},
	}

	tests := map[rune][]byte{
		'a': {0, 1, 0},
		'b': {0, 1, 1},
		'c': {0, 0},
		'd': {1, 0},
		'英': {1, 1},
	}

	for char, expected := range tests {
		code, err := tree.GetCode(char)
		if err != nil {
			t.Errorf("unexpected error %v", err)
			t.FailNow()
		}
		if !reflect.DeepEqual(code, expected) {
			t.Errorf("expected for '%s' %v, got %v", "a", expected, code)
			t.FailNow()
		}
	}
}
