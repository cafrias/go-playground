package huffman

import "testing"

func TestDecode(t *testing.T) {
	input := []byte{0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	tree := &HuffmanTree{
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

	result, err := Decode(input, tree)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		t.FailNow()
	}

	expected := "abbcccdddd英英英英英"
	if result != expected {
		t.Errorf("expected %s got %s", expected, result)
		t.FailNow()
	}
}
