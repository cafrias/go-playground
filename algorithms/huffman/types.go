package huffman

type HuffmanTree struct {
	Frequency uint
	Character rune
	Left      *HuffmanTree
	Right     *HuffmanTree
}

func (t *HuffmanTree) Value() uint {
	return t.Frequency
}
