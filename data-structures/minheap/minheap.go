package minheap

import (
	"cmp"
	"errors"
)

type Heapable[T cmp.Ordered] interface {
	Value() T
}

type MinHeap[V cmp.Ordered, H Heapable[V]] struct {
	Items []H
}

func (m *MinHeap[T, H]) Insert(value H) {
	m.Items = append(m.Items, value)
	if len(m.Items) == 1 {
		return
	}

	for i := len(m.Items) - 1; i > 0; i = (i - 1) / 2 {
		pIdx := (i - 1) / 2
		p := m.Items[pIdx]
		if m.Items[i].Value() < p.Value() {
			m.Items[pIdx] = m.Items[i]
			m.Items[i] = p
		} else {
			break
		}

	}
}

var ErrHeapEmpty = errors.New("Heap is empty")

func (m *MinHeap[V, H]) Delete() (H, error) {
	if len(m.Items) == 0 {
		return *new(H), ErrHeapEmpty
	}

	lastIdx := len(m.Items) - 1
	popped := m.Items[0]
	m.Items[0] = m.Items[lastIdx]
	m.Items = m.Items[:lastIdx]

	cIdx := 0
	cVal := m.Items[0].Value()
	for {
		lIdx := 2*cIdx + 1
		if lIdx >= len(m.Items) {
			break
		}
		lVal := m.Items[lIdx].Value()

		rIdx := 2*cIdx + 2
		var rVal V
		if rIdx >= len(m.Items) {
			rVal = lVal
		} else {
			rVal = m.Items[rIdx].Value()
		}

		if cVal <= lVal && cVal <= rVal {
			break
		}

		if lVal <= rVal {
			aux := m.Items[cIdx]
			m.Items[cIdx] = m.Items[lIdx]
			m.Items[lIdx] = aux
			cIdx = lIdx
		} else {
			aux := m.Items[cIdx]
			m.Items[cIdx] = m.Items[rIdx]
			m.Items[rIdx] = aux
			cIdx = rIdx
		}
	}

	return popped, nil
}
