package stack

type Stack struct {
	data []interface{}
}

func New() *Stack {
	return &Stack{
		data: make([]interface{}, 0),
	}
}

func (s *Stack) Pop() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

func (s *Stack) Push(v interface{}) {
	s.data = append(s.data, v)
}

func (s *Stack) Len() int {
	return len(s.data)
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	m := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := New()

	for _, c := range s {
		if v, ok := m[c]; ok {
			if rv, ok := stack.Pop().(rune); !ok || rv != v {
				return false
			}
		} else {
			stack.Push(c)
		}
	}
	// 注意最后栈中不能有数据剩余
	return stack.Len() == 0
}
