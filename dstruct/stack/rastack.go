package stack

// LIFO

type ResizingArrayStack struct {
	items []string
	count int
}

func NewRAStack(cap int) *ResizingArrayStack {
	return &ResizingArrayStack{
		items: make([]string, cap),
	}
}

func (s *ResizingArrayStack) Push(item string) {
	s.count++
	s.items[s.count] = item
}

func (s *ResizingArrayStack) Pop() string {
	item := s.items[s.count]
	s.count--
	return item
}

func (s *ResizingArrayStack) IsEmpty() bool {
	return s.count == 0
}

func (s *ResizingArrayStack) Size() int {
	return s.count
}
