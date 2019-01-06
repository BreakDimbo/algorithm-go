package tree

func generateParenthesis(n int) []string {
	p := &Par{make([]string, 0)}
	p.gen(0, 0, n, "")
	return p.listR
}

type Par struct {
	listR []string
}

func (p *Par) gen(left, right, n int, result string) {
	if left == n && right == n {
		p.listR = append(p.listR, result)
	}

	if left < n {
		p.gen(left+1, right, n, result+"(")
	}
	if left > right && right < n {
		p.gen(left, right+1, n, result+")")
	}
}
