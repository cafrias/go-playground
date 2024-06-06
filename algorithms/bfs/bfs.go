package bfs

import (
	"fmt"
	"go-playground/data-structures/queue"
)

func BFS(graph map[string][]string, start string) ([]string, error) {
	result := make([]string, 0, len(graph))
	visited := make(map[string]bool, len(graph))

	q := queue.Queue[string]{}
	currentNode := start

	for {
		neighbors, ok := graph[currentNode]
		if !ok {
			return nil, fmt.Errorf("Node %s not found in graph", start)
		}

		if _, isVisited := visited[currentNode]; !isVisited {
			for _, neighbor := range neighbors {
				q.Enqueue(neighbor)
			}

			result = append(result, currentNode)
			visited[currentNode] = true
		}

		var err error
		currentNode, err = q.Dequeue()
		if err == queue.EmptyQueueErr {
			return result, nil
		}
	}
}
