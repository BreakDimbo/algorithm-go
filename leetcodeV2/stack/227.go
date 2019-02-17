package stack

import "unicode"

/*
	solution 1: stack
*/

func calculate(s string) int {

	stack := make([]int, 0)
	num := 0
	preSign := '+'
	for i, ch := range s {
		if unicode.IsDigit(ch) {
			num = 10*num + int(ch-'0')
		}

		if !unicode.IsDigit(ch) && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, pop*num)
			case '/':
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, pop/num)
			}
			preSign = ch
			num = 0
		}
	}

	res := 0
	for _, v := range stack {
		res += v
	}
	return res
}
