package graph

import "errors"

type Edge struct {
	To     int
	Weight uint
}

type AdjacencyList = [][]Edge

func NewAdjacencyList(n int) AdjacencyList {
	return make([][]Edge, n)
}

func AdjacencyListDFS(list AdjacencyList, source int, needle int) ([]int, error) {
	result := make([]int, 0, len(list))
	visited := make([]bool, len(list))
	return walkAdjacencyList(result, list, source, needle, visited)

}

func walkAdjacencyList(result []int, list AdjacencyList, currentNode int, needle int, visited []bool) ([]int, error) {
	if currentNode == needle {
		return result, nil
	}

	if visited[currentNode] {
		return result, errors.New("Node already visited")
	}

	visited[currentNode] = true
	result = append(result, currentNode)

	for _, edge := range list[currentNode] {
		result, err := walkAdjacencyList(result, list, edge.To, needle, visited)
		if err == nil {
			return result, nil
		}
	}

	return result, errors.New("Not found")
}
