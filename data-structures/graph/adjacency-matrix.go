package graph

import (
	"errors"
	"fmt"
	"go-playground/data-structures/queue"
	"slices"
)

type AdjacencyMatrix = [][]int

func NewAdjecencyMatrix(n uint) AdjacencyMatrix {
	matrix := make(AdjacencyMatrix, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	return matrix
}

var ErrNotFound = errors.New("Not Found")

func matrixBFS(matrix AdjacencyMatrix, start int, needle int) ([]int, error) {
	prev := make([]int, len(matrix))
	for i := range prev {
		prev[i] = -1
	}
	q := queue.Queue[int]{}
	q.Enqueue(start)
	v := make([]bool, len(matrix))
	v[start] = true

	for {
		curr, err := q.Dequeue()
		if err != nil {
			break
		}

		if curr == needle {
			path := make([]int, 0, len(matrix))
			for i := needle; prev[i] != -1; i = prev[i] {
				path = append(path, prev[i])
			}
			slices.Reverse(path)
			return path, nil
		}

		for node, weight := range matrix[curr] {
			if weight > 0 && !v[node] {
				q.Enqueue(node)
				v[node] = true
				prev[node] = curr
			}
		}
	}

	return nil, fmt.Errorf("Error searching for '%v' starting in '%v': %w", needle, start, ErrNotFound)
}
