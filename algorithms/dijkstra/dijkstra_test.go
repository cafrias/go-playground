package dijkstra

import (
	"go-playground/data-structures/graph"
	"reflect"
	"testing"
)

func TestFindShortPath(t *testing.T) {
	list := createList()

	path, cost, err := FindShortPath(list, 0, 6)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		t.FailNow()
	}

	exp := []int{0, 1, 4, 5, 6}
	if !reflect.DeepEqual(path, exp) {
		t.Errorf("Invalid path\nexpected: %v\ngot: %v", exp, path)
		t.FailNow()
	}

	var expCost uint = 20
	if cost != expCost {
		t.Errorf("Invalid cost\nexpected: %v\ngot: %v", expCost, cost)
		t.FailNow()
	}
}

func createList() graph.AdjacencyList {
	l := graph.NewAdjacencyList(7)

	//     >(1)<--->(4) ---->(5)
	//    /          |       /|
	// (0)     ------|------- |
	//    \   v      v        v
	//     >(2) --> (3) <----(6)
	l[0] = []graph.Edge{
		{To: 1, Weight: 3},
		{To: 2, Weight: 1},
	}
	l[1] = []graph.Edge{
		{To: 0, Weight: 3},
		{To: 2, Weight: 4},
		{To: 4, Weight: 1},
	}
	l[2] = []graph.Edge{
		{To: 1, Weight: 4},
		{To: 3, Weight: 7},
		{To: 0, Weight: 1},
	}
	l[3] = []graph.Edge{
		{To: 2, Weight: 7},
		{To: 4, Weight: 5},
		{To: 6, Weight: 1},
	}
	l[4] = []graph.Edge{
		{To: 1, Weight: 1},
		{To: 3, Weight: 5},
		{To: 5, Weight: 2},
	}
	l[5] = []graph.Edge{
		{To: 6, Weight: 1},
		{To: 4, Weight: 2},
		{To: 2, Weight: 18},
	}
	l[6] = []graph.Edge{
		{To: 3, Weight: 1},
		{To: 5, Weight: 1},
	}

	return l
}
