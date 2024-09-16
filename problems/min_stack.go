package problems

// Minstack solves LeetCode problem 155. Min Stack
type Minstack struct {
	mins []int
	v    []int
}

func NewMinStack() Minstack {
	return Minstack{
		mins: make([]int, 0),
		v:    make([]int, 0),
	}
}

// Push pushes element `val` onto the stack
func (s *Minstack) Push(val int) {
	s.v = append(s.v, val)

	if len(s.mins) == 0 || val <= s.GetMin() {
		s.mins = append(s.mins, val)
	}
}

// Pop removes the top element of the stack
func (s *Minstack) Pop() {
	val := s.Top()
	s.v = s.v[:len(s.v)-1]

	if val == s.GetMin() {
		s.mins = s.mins[:len(s.mins)-1]
	}
}

// Top gets the top element of the stack
func (s *Minstack) Top() int {
	return s.v[len(s.v)-1]
}

func (s *Minstack) GetMin() int {
	return s.mins[len(s.mins)-1]
}
