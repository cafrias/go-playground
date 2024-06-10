package graph

import (
	"reflect"
	"testing"
)

func TestMatrixBFS(t *testing.T) {
	m := NewAdjecencyMatrix(6)
	m[0][5] = 3
	m[1][0] = 6
	m[1][3] = 5
	m[2][0] = 1
	m[2][1] = 2
	m[2][4] = 4
	m[3][4] = 2
	m[4][2] = 2
	m[5][2] = 2
	m[5][4] = 1

	start := 0
	needle := 3
	expPath := []int{
		0, 5, 2, 1,
	}

	path, err := matrixBFS(m, start, needle)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		t.FailNow()
	}

	if !reflect.DeepEqual(path, expPath) {
		t.Errorf("Invalid path\ngot %v\nexpected %v", path, expPath)
		t.FailNow()
	}
}
