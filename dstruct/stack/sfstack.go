package stack

// LIFO

type SimpleFixedStack struct {
	items []string
	count int
}

func NewSFStack(cap int) *SimpleFixedStack {
	return &SimpleFixedStack{
		items: make([]string, cap),
	}
}

func (s *SimpleFixedStack) Push(item string) {
	s.count++
	s.items[s.count] = item
}

func (s *SimpleFixedStack) Pop() string {
	item := s.items[s.count]
	s.count--
	return item
}

func (s *SimpleFixedStack) IsEmpty() bool {
	return s.count == 0
}

func (s *SimpleFixedStack) Size() int {
	return s.count
}
