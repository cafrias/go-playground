package minheap

import (
	"reflect"
	"testing"
)

type testHeapable struct {
	v int
}

func (h testHeapable) Value() int {
	return h.v
}

func TestMinHeap(t *testing.T) {
	heap := MinHeap[int, testHeapable]{
		Items: make([]testHeapable, 0, 10),
	}

	heap.Insert(testHeapable{3})
	heap.Insert(testHeapable{5})
	heap.Insert(testHeapable{7})
	heap.Insert(testHeapable{9})
	heap.Insert(testHeapable{4})
	heap.Insert(testHeapable{2})

	expected := []testHeapable{
		{2},
		{4},
		{3},
		{9},
		{5},
		{7},
	}
	if !reflect.DeepEqual(heap.Items, expected) {
		t.Errorf("expected %v received %v", expected, heap.Items)
	}

	res, err := heap.Delete()
	if err != nil {
		t.Errorf("Unexpected error %v", err)
		return
	}

	if res.Value() != 2 {
		t.Errorf("Returned %v, expected %v", res.Value(), 2)
		return
	}

	expected = []testHeapable{
		{3},
		{4},
		{7},
		{9},
		{5},
	}
	if !reflect.DeepEqual(heap.Items, expected) {
		t.Errorf("expected %v received %v", expected, heap.Items)
	}

	_, _ = heap.Delete()
	_, _ = heap.Delete()
	expected = []testHeapable{
		{5},
		{9},
		{7},
	}
	if !reflect.DeepEqual(heap.Items, expected) {
		t.Errorf("expected %v received %v", expected, heap.Items)
	}
}
