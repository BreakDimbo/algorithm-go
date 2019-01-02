package stackqueue

type Stack []string

func NewStack() *Stack {
	var stack []string
	return (*Stack)(&stack)
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() string {
	if len(*s) == 0 {
		return ""
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func isValid(s string) bool {
	left := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	stack := NewStack()
	for _, v := range s {
		e := string(v)
		if left[e] != "" {
			stack.Push(e)
		} else {
			enclose := stack.Pop()
			if left[enclose] != e {
				return false
			}
		}
	}

	return len(*stack) == 0
}
