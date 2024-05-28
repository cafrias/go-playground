package bfs

import (
	"reflect"
	"testing"
)

func TestBFS(t *testing.T) {
	// Create a graph
	graph := make(map[string][]string)
	graph["A"] = []string{"B", "C"}
	graph["B"] = []string{"A", "D", "E"}
	graph["C"] = []string{"A", "F"}
	graph["D"] = []string{"B"}
	graph["E"] = []string{"B", "F"}
	graph["F"] = []string{"C", "E"}

	// Run BFS
	result, err := BFS(graph, "A")
	if err != nil {
		t.Errorf("Expected no error but received %v", err)
	}

	expected := []string{
		"A", "B", "C", "D", "E", "F",
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v received %v", expected, result)
	}
}
