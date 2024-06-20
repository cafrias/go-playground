package dijkstra

import (
	"go-playground/data-structures/graph"
	"go-playground/data-structures/minheap"
	"math"
	"slices"
)

type weightedNode struct {
	id     int
	weight uint
}

func (w weightedNode) Value() uint {
	return w.weight
}

func FindShortPath(list graph.AdjacencyList, source int, target int) (path []int, cost uint, err error) {
	seen := make([]bool, len(list))
	dists := make([]uint, len(list))
	for i := range dists {
		dists[i] = math.MaxUint
	}
	dists[source] = 0
	prev := make([]int, len(list))
	for i := range prev {
		prev[i] = -1
	}

	h := minheap.MinHeap[uint, weightedNode]{}
	h.Insert(weightedNode{source, 0})

	for {
		w, err := h.Delete()
		if err != nil {
			break
		}
		nodeId := w.id

		if seen[nodeId] {
			continue
		}

		seen[nodeId] = true

		nodeDist := dists[nodeId]
		for _, adj := range list[nodeId] {
			adjId := adj.To
			dist := adj.Weight + nodeDist
			if dist < dists[adjId] {
				dists[adjId] = dist
				prev[adjId] = nodeId
			}

			if !seen[adjId] {
				h.Insert(
					weightedNode{
						id:     adjId,
						weight: dist,
					},
				)
			}
		}
	}

	path = append(path, target)
	for i := target; prev[i] != -1; i = prev[i] {
		path = append(path, prev[i])
		cost += dists[i]
	}
	slices.Reverse(path)

	return path, cost, nil
}
