package huffman

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestCreateHuffmanTree(t *testing.T) {
	tests := []struct {
		input    io.Reader
		expected *HuffmanTree
	}{
		{
			input: strings.NewReader("abbcccdddd英英英英英"),
			expected: &HuffmanTree{
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
			},
		},
		{
			input:    strings.NewReader(""),
			expected: nil,
		},
		{
			input: strings.NewReader("aaaaaaa"),
			expected: &HuffmanTree{
				Frequency: 7,
				Left: &HuffmanTree{
					Frequency: 7,
					Character: 'a',
				},
			},
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			tree := createHuffmanTree(tt.input)

			tRes := make([]struct {
				uint
				string
			}, 0)
			tRes = dfs(tree, tRes)
			tExp := make([]struct {
				uint
				string
			}, 0)
			tExp = dfs(tt.expected, tExp)

			if !reflect.DeepEqual(tRes, tExp) {
				t.Errorf("expected %v, got %v", tExp, tRes)
				t.FailNow()
			}

		})
	}

}

func dfs(t *HuffmanTree, inResult []struct {
	uint
	string
}) []struct {
	uint
	string
} {
	if t == nil {
		return inResult
	}

	outResult := dfs(t.Left, inResult)
	outResult = append(outResult, struct {
		uint
		string
	}{
		t.Frequency,
		string(t.Character),
	})
	outResult = dfs(t.Right, outResult)

	return outResult
}
