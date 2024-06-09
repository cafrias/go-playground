package huffman

import (
	"errors"
	"strings"
)

func Decode(encoded []byte, tree *HuffmanTree) (string, error) {
	builder := strings.Builder{}
	emptyChar := *new(rune)
	currentNode := tree
	for _, b := range encoded {
		if b == 0 {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}

		if currentNode == nil {
			return "", errors.New("Path not found in tree")
		}

		if currentNode.Character != emptyChar {
			builder.WriteRune(currentNode.Character)
			currentNode = tree
			continue
		}

	}

	return builder.String(), nil
}
