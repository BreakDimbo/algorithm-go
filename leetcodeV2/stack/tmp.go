package stack

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	h := map[byte]byte{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	for i := range s {
		if s[i] == ' ' {
			continue
		}
		if _, ok := h[s[i]]; ok {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || h[stack[len(stack)-1]] != s[i] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}
