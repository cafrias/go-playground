package problems

import "testing"

func TestMinStack(t *testing.T) {
	stack := NewMinStack()

	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)

	minVal := stack.GetMin()
	if minVal != -3 {
		t.Fatalf("expected %v, got %v", -3, minVal)
	}

	stack.Pop()

	topVal := stack.Top()
	if topVal != 0 {
		t.Fatalf("expected %v, got %v", 0, topVal)
	}

	minVal = stack.GetMin()
	if minVal != -2 {
		t.Fatalf("expected %v, got %v", -2, minVal)
	}
}
