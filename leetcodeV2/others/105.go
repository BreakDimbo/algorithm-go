package others

import (
	"strconv"
)

/*
	solution 1: use stack
*/

func evalRPN(tokens []string) int {
	operators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}

	stack := make([]string, 0)
	for _, token := range tokens {
		if !operators[token] {
			stack = append(stack, token)
		} else {
			op1Str, op2Str := stack[len(stack)-2], stack[len(stack)-1]
			op1, _ := strconv.Atoi(op1Str)
			op2, _ := strconv.Atoi(op2Str)
			stack = stack[:len(stack)-2]
			var res string
			switch token {
			case "+":
				res = strconv.Itoa(op1 + op2)
			case "-":
				res = strconv.Itoa(op1 - op2)
			case "*":
				res = strconv.Itoa(op1 * op2)
			case "/":
				res = strconv.Itoa(op1 / op2)
			}
			stack = append(stack, res)
		}
	}
	result, _ := strconv.Atoi(stack[len(stack)-1])
	return result
}
