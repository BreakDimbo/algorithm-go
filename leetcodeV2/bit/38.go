package bit

/*
	solution 1: 递归
	solution 2: 递推
*/

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	s := countAndSay(n - 1)
	res := ""
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && s[i] == '1' && s[i+1] == '1' {
			res += "21"
			i++ // 注意这里以此消了两个
		} else if s[i] == '1' {
			res += "11"
		} else if s[i] == '2' {
			res += "12"
		}
	}
	return res
}
