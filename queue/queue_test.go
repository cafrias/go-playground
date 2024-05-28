package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := Queue[int]{}

	q.Enqueue(1)
	if q.First.value != 1 || q.Last.value != 1 {
		t.Error()
	}

	q.Enqueue(2)
	if q.First.next.value != 2 || q.Last.value != 2 {
		t.Errorf("Expected last to be 2 but it received %v", q.Last.value)
	}

	q.Enqueue(3)
	if q.First.next.value != 2 || q.Last.value != 3 {
		t.Errorf("Expected last to be 3 but it received %v", q.Last.value)
	}

	v, err := q.Dequeue()
	if err != nil {
		t.Errorf("Expected no error when dequeuing but received one")
	}
	if v != 1 {
		t.Errorf("Expected dequeued value to be 1, but received %v", v)
	}
	if q.First.value != 2 {
		t.Errorf("Expected deque to remove 1 from the queue, but it didn't")
	}

	q.Dequeue()
	v, _ = q.Dequeue()
	if v != 3 {
		t.Errorf("Expected dequeued value to be 3, but received %v", v)
	}
	if q.First != nil || q.Last != nil {
		t.Errorf("Expected queue to be empty")
	}

	v, err = q.Dequeue()
	if v != 0 {
		t.Errorf("Expected empty value to return from empty queue, instead received %v", v)
	}
	if err != EmptyQueueErr {
		t.Errorf("Expected emtpy queue error to be returned, instead received %v", err)
	}
}
