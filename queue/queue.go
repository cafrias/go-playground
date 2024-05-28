package queue

import "errors"

type item[T any] struct {
	value T
	next  *item[T]
}

type Queue[T any] struct {
	First *item[T]
	Last  *item[T]
}

func (q *Queue[T]) Enqueue(value T) {
	newItem := &item[T]{
		value: value,
	}

	if q.First == nil {
		q.First = newItem
		q.Last = newItem
		return
	}

	if q.Last == nil {
		q.Last = newItem
	}

	q.Last.next = newItem
	q.Last = newItem
}

var EmptyQueueErr = errors.New("The queue is empty")

func (q *Queue[T]) Dequeue() (T, error) {
	if q.First == nil {
		return *new(T), EmptyQueueErr
	}

	first := q.First

	if first == q.Last {
		q.First = nil
		q.Last = nil
	} else {
		q.First = q.First.next
	}

	return first.value, nil
}
